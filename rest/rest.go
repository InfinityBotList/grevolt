package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/version"
	"github.com/sethgrid/pester"
)

// Makes a request to the API
func (r Request[T]) Request(config *RestConfig) (*http.Response, error) {
	if r.Method == "" {
		r.Method = GET
	}

	if r.bucket == nil {
		r.bucket = config.Ratelimiter.LockBucket(string(r.Method) + ":" + strings.SplitN(r.Path, "?", 2)[0])
	}

	if r.sequence > 0 {
		// Exp backoff, 2^sequence * 100ms
		time.Sleep(time.Duration(1<<r.sequence) * 100 * time.Millisecond)
	}

	config.Logger.Debug("Acquired bucket ", r.bucket)

	if r.bucket != nil {
		config.Logger.Debug("Bucket name ", r.bucket.Key)
	}

	var body []byte
	var err error
	if r.Json != nil {
		body, err = json.Marshal(r.Json)

		if err != nil {
			r.bucket.Release(nil)
			return nil, err
		}
	}

	config.Logger.Debug(r.Method, config.APIUrl+r.Path, " (reqBody:", len(body), "bytes)")

	config.Logger.Debugln("MakeNewRequest", r.Method, config.APIUrl+r.Path, " (reqBody:", len(body), "bytes)")
	req, err := http.NewRequest(string(r.Method), config.APIUrl+r.Path, bytes.NewReader(body))

	if err != nil {
		r.bucket.Release(nil)
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	req.Header.Add("User-Agent", "grevolt/"+version.Version)
	req.Header.Add("Content-Type", "application/json")

	config.Pester.Timeout = config.Timeout
	config.Pester.MaxRetries = config.MaxRestRetries
	config.Pester.Backoff = pester.ExponentialBackoff
	config.Pester.KeepLog = true
	config.Pester.RetryOnHTTP429 = false

	resp, err := pester.Do(req)

	if err != nil {
		r.bucket.Release(nil)
		return nil, err
	}

	err = r.bucket.Release(resp.Header)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusBadGateway:
		// Retry sending request if possible
		if r.sequence < config.MaxRestRetries {

			config.Logger.Errorf("%s Failed (%s), Retrying...", r.Path, resp.Status)

			config.Ratelimiter.LockBucketObject(r.bucket)
			r.sequence++

			return r.Request(config)
		} else {
			return nil, fmt.Errorf("exceeded Max retries HTTP %s, %s", resp.Status, r.Path)
		}
	case http.StatusTooManyRequests:
		// Rate limited
		var rl *types.RateLimit
		err = json.NewDecoder(resp.Body).Decode(&rl)
		if err != nil {
			config.Logger.Errorf("rate limit unmarshal error, %s", err)
			return nil, err
		}

		if config.RetryOnRatelimit {
			config.Logger.Infof("Rate Limiting %s, retry in %v", r.Path, rl.RetryAfter)
			time.Sleep(time.Duration(rl.RetryAfter) * time.Millisecond)

			config.Ratelimiter.LockBucketObject(r.bucket)
			r.sequence++

			return r.Request(config)
		} else {
			return nil, errors.New("ratelimited:" + strconv.Itoa(int(rl.RetryAfter)))
		}
	}

	return resp, nil
}

// Executes the request and unmarshals the response body if the response is OK otherwise returns error
func (r Request[T]) With(config *RestConfig) (*T, error) {
	if len(r.Headers) == 0 {
		r.Headers = make(map[string]string)
	}

	if config.SessionToken != nil {
		if config.SessionToken.Bot {
			r.Headers["x-bot-token"] = config.SessionToken.Token
		} else {
			r.Headers["x-session-token"] = config.SessionToken.Token
		}
	}

	resp, err := r.Request(config)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		var v T

		if _, ok := any(v).(Bytes); ok {
			// We have byte as type, just read response as body
			body, err := io.ReadAll(resp.Body)

			if err != nil {
				return nil, err
			}

			res := any(Bytes{
				Raw: body,
			}).(T)

			return &res, nil
		}

		err = json.NewDecoder(resp.Body).Decode(&v)

		if err != nil {
			return nil, err
		}

		return &v, nil
	} else {
		if resp.StatusCode == 401 {
			// Try and read body
			body := resp.Body

			if body == nil {
				return nil, fmt.Errorf("unauthorized")
			}

			bodyBytes, err := io.ReadAll(body)

			if err != nil {
				return nil, fmt.Errorf("unauthorized, could not read body")
			}

			return nil, fmt.Errorf("unauthorized, body: %s", string(bodyBytes))
		}

		var vErr types.APIError
		err = json.NewDecoder(resp.Body).Decode(&vErr)

		if err != nil {
			return nil, err
		}

		return nil, RestError{APIError: vErr}
	}
}

// Discards the content as long as the response is OK otherwise error is returned
func (r Request[T]) NoContent(config *RestConfig) error {
	if len(r.Headers) == 0 {
		r.Headers = make(map[string]string)
	}

	if config.SessionToken != nil {
		if config.SessionToken.Bot {
			r.Headers["x-bot-token"] = config.SessionToken.Token
		} else {
			r.Headers["x-session-token"] = config.SessionToken.Token
		}
	}

	resp, err := r.Request(config)

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		var vErr types.APIError
		err = json.NewDecoder(resp.Body).Decode(&vErr)

		if err != nil {
			return err
		}

		return RestError{APIError: vErr}
	}

	return nil
}
