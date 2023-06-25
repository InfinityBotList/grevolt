package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/rest/ratelimits"
	"github.com/infinitybotlist/grevolt/types"
	"github.com/sethgrid/pester"
	"go.uber.org/zap"
)

type RestConfig struct {
	// The URL of the API
	APIUrl string

	// Timeout for requests
	Timeout time.Duration

	// Logger to use, will be autofilled if not provided
	Logger *zap.SugaredLogger

	// Session token for requests
	SessionToken *auth.Token

	// Ratelimiter
	Ratelimiter *ratelimits.RateLimiter

	// Max tries for requests
	MaxRestRetries int

	// On ratelimit function
	OnRatelimit func(*types.RateLimit)

	// Whether or not to retry on ratelimit
	RetryOnRatelimit bool

	// Pester client
	Pester *pester.Client
}

// DefaultRestConfig is the default configuration for the client
func DefaultRestConfig() RestConfig {
	return RestConfig{
		APIUrl:      "https://api.revolt.chat/",
		Timeout:     10 * time.Second,
		Ratelimiter: ratelimits.NewRatelimiter(),
		OnRatelimit: func(r *types.RateLimit) {
			fmt.Println("Ratelimited:", r)
		},
		RetryOnRatelimit: true,
		Pester:           pester.New(),
	}
}

type Method string

const (
	HEAD    Method = "HEAD"
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	OPTIONS Method = "OPTIONS"
)

// A request to the API
type Request[T any] struct {
	Method  Method
	Path    string
	Json    any
	Headers map[string]string

	// Sequence number for this request, internal
	sequence int

	// Ratelimit bucket, internal
	bucket *ratelimits.Bucket
}

// A response from the API
type Response[T any] struct {
	Request  *Request[T]
	Response *http.Response
}

// Represents a raw byte array (avatars for example)
type Bytes struct {
	Raw []byte
}

// Note that a RestError satisfies the error interface
type RestError struct {
	types.APIError
}

func (r RestError) Error() string {
	return fmt.Sprintln("API Error:", r.APIError.Type(), "", r.APIError)
}
