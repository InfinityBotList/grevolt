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
	"go.uber.org/zap"
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

	if r.bucket != nil {
		config.Logger.Debug(
			"Acquired bucket",
			zap.String("key", r.bucket.Key),
			zap.Int("limit", r.bucket.Limit),
			zap.Int("remaining", r.bucket.Remaining),
		)
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

	config.Logger.Debug(
		"Make request",
		zap.String("method", string(r.Method)),
		zap.String("url", config.APIUrl+r.Path),
		zap.Int("bodySize", len(body)),
	)
	req, err := http.NewRequest(string(r.Method), config.APIUrl+r.Path, bytes.NewReader(body))

	if err != nil {
		r.bucket.Release(nil)
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	for _, cookie := range r.Cookies {
		req.AddCookie(&cookie)
	}

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
	case http.StatusBadGateway, http.StatusGatewayTimeout, http.StatusServiceUnavailable:
		// Retry sending request if possible
		if r.sequence < config.MaxRestRetries {

			config.Logger.Error("Request failed, retrying...", zap.String("path", r.Path), zap.String("status", resp.Status))

			config.Ratelimiter.LockBucketObject(r.bucket)
			r.sequence++

			return r.Request(config)
		} else {
			return nil, fmt.Errorf("exceeded max retries HTTP %s, %s", resp.Status, r.Path)
		}
	case http.StatusTooManyRequests:
		// Rate limited
		var rl *types.RateLimit
		err = json.NewDecoder(resp.Body).Decode(&rl)
		if err != nil {
			return nil, errors.New("rate limit unmarshal error: " + err.Error())
		}

		if config.RetryOnRatelimit {
			config.Logger.Error("Request failed [ratelimited]", zap.String("path", r.Path), zap.String("status", resp.Status), zap.Int64("retryIn", rl.RetryAfter))
			time.Sleep(time.Duration(rl.RetryAfter)*time.Millisecond + time.Duration(r.sequence)*2*time.Millisecond)

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

		if r.InitialResp != nil {
			v = *r.InitialResp
		}

		err = json.NewDecoder(resp.Body).Decode(&v)

		if err != nil {
			return nil, err
		}

		for _, f := range config.OnMarshal {
			err = f(&RequestData{
				Method:  r.Method,
				Path:    r.Path,
				Json:    r.Json,
				Headers: r.Headers,
				Config:  config,
			}, &v)

			if err != nil {
				return nil, err
			}
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

		if config.SessionToken.CfClearance != nil {
			r.Headers["user-agent"] = config.SessionToken.CfClearance.UserAgent
			r.Cookies = append(r.Cookies, http.Cookie{
				Name:  "cf_clearance",
				Value: config.SessionToken.CfClearance.CookieValue,
			})
		} else {
			r.Headers["user-agent"] = "grevolt/" + version.Version
		}
	}

	resp, err := r.Request(config)

	if err != nil {
		return err
	}

	config.Logger.Debug("Request made", zap.Int("statusCode", resp.StatusCode))

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
