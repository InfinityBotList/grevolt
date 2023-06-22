package clientapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/infinitybotlist/grevolt/rest/restconfig"
	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/version"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// A response from the API
type ClientResponse struct {
	Request  *http.Request
	Response *http.Response
}

// Returns true if the response is a maintenance response (502, 503, 408)
func (c ClientResponse) IsMaint() bool {
	return c.Response.StatusCode == 502 || c.Response.StatusCode == 503 || c.Response.StatusCode == 408
}

// Unmarshals the response body into the given struct
func (c ClientResponse) Json(v any) error {
	return json.NewDecoder(c.Response.Body).Decode(v)
}

// Returns whether the response is OK
func (c ClientResponse) Ok() bool {
	return c.Response.StatusCode == 200
}

// Unmarshals response body if response is OK otherwise returns error
func (c ClientResponse) JsonOk(v any) error {
	if c.Ok() {
		return fmt.Errorf("error status code %d", c.Response.StatusCode)
	}

	return c.Json(v)
}

// Returns the response body
func (c ClientResponse) Body() ([]byte, error) {
	return io.ReadAll(c.Response.Body)
}

// Returns the response body if the response is OK otherwise returns error
func (c ClientResponse) BodyOk() ([]byte, error) {
	if c.Ok() {
		return nil, fmt.Errorf("error status code %d", c.Response.StatusCode)
	}

	return c.Body()
}

// Returns the retry after header. Is a string
func (c ClientResponse) RetryAfter() string {
	return c.Response.Header.Get("Retry-After")
}

// A request to the API
type ClientRequest struct {
	method  string
	path    string
	json    any
	headers map[string]string
	config  *restconfig.RestConfig
}

// Makes a request to the API
func (r ClientRequest) request() (*ClientResponse, error) {
	if r.method == "" {
		r.method = "GET"
	}

	if !strings.HasPrefix(r.path, "/") {
		r.path = "/" + r.path
	}

	var body []byte
	var err error
	if r.json != nil {
		body, err = json.Marshal(r.json)

		if err != nil {
			return nil, err
		}
	}

	r.config.Logger.Debug(r.method, r.config.APIUrl+r.path, " (reqBody:", len(body), "bytes)")

	req, err := http.NewRequest(r.method, r.config.APIUrl+r.path, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	for k, v := range r.headers {
		req.Header.Add(k, v)
	}

	req.Header.Add("User-Agent", "grevolt/"+version.Version)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: r.config.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return &ClientResponse{
		Request:  req,
		Response: resp,
	}, nil
}

// Creates a new request, must be followed by a method call otherwise the request will be invalid
func NewReq(config *restconfig.RestConfig) ClientRequest {
	return ClientRequest{
		config:  config,
		headers: make(map[string]string),
	}
}

// Sets the method of the request
func (r ClientRequest) Method(method string) ClientRequest {
	r.method = method
	return r
}

// Sets the method to HEAD
func (r ClientRequest) Head(path string) ClientRequest {
	r.method = "HEAD"
	r.path = path
	return r
}

// Sets the method to GET
func (r ClientRequest) Get(path string) ClientRequest {
	r.method = "GET"
	r.path = path
	return r
}

// Sets the method to POST
func (r ClientRequest) Post(path string) ClientRequest {
	r.method = "POST"
	r.path = path
	return r
}

// Sets the method to PUT
func (r ClientRequest) Put(path string) ClientRequest {
	r.method = "PUT"
	r.path = path
	return r
}

// Sets the method to PATCH
func (r ClientRequest) Patch(path string) ClientRequest {
	r.method = "PATCH"
	r.path = path
	return r
}

// Sets the method to DELETE
func (r ClientRequest) Delete(path string) ClientRequest {
	r.method = "DELETE"
	r.path = path
	return r
}

// Sets the path of the request
func (r ClientRequest) Path(path string) ClientRequest {
	r.path = path
	return r
}

// Sets the body of the request
func (r ClientRequest) Json(json any) ClientRequest {
	r.json = json
	return r
}

// Sets a header
func (r ClientRequest) Header(key string, value string) ClientRequest {
	r.headers[key] = value
	return r
}

func (r ClientRequest) AutoLogger() ClientRequest {
	w := zapcore.AddSync(os.Stdout)

	var level = zap.InfoLevel
	if os.Getenv("DEBUG") == "true" {
		level = zap.DebugLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		level,
	)

	r.config.Logger = zap.New(core).Sugar()

	return r
}

// Executes the request
func (r ClientRequest) Do() (*ClientResponse, error) {
	if r.config.Logger == nil {
		r.AutoLogger()
	}

	if r.config.SessionToken != nil {
		if r.config.SessionToken.Bot {
			r.headers["x-bot-token"] = r.config.SessionToken.Token
		} else {
			r.headers["x-session-token"] = r.config.SessionToken.Token
		}
	}

	r.config.Logger.Debug(r.headers)

	return r.request()
}

// Executes the request and unmarshals the response body if the response is OK otherwise returns error
func (r ClientRequest) DoAndMarshal(v any) (*types.APIError, error) {
	resp, err := r.Do()

	if err != nil {
		return nil, err
	}

	if resp.Ok() {
		err = resp.Json(v)

		if err != nil {
			return nil, err
		}

		return nil, nil
	} else {
		if resp.Response.StatusCode == 401 {
			// Try and read body
			body := resp.Response.Body

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
		err = resp.Json(&vErr)

		if err != nil {
			return nil, err
		}

		return &vErr, nil
	}
}
