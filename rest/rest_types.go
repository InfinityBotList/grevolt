package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/cache/state"
	"github.com/infinitybotlist/grevolt/rest/ratelimits"
	"github.com/infinitybotlist/grevolt/types"
	"github.com/sethgrid/pester"
	"go.uber.org/zap"
)

const (
	// Revolts API
	RevoltAPI = "https://api.revolt.chat/"

	// The staging API, used by the official apps etc, default
	RevoltAPIStaging = "https://app.revolt.chat/api/"
)

type RestConfig struct {
	// The URL of the API
	APIUrl string

	// Timeout for requests
	Timeout time.Duration

	// Logger to use, will be autofilled if not provided
	Logger *zap.Logger

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

	// Functions to run upon successful marshal
	OnMarshal []func(r *RequestData, v any) error

	// Shared state for rest requests
	SharedState *state.State

	// Disable rest caching
	DisableRestCaching bool
}

// DefaultRestConfig return the default configuration for the client with the given state
func DefaultRestConfig(state *state.State) RestConfig {
	return RestConfig{
		APIUrl:      RevoltAPIStaging,
		Timeout:     10 * time.Second,
		Ratelimiter: ratelimits.NewRatelimiter(),
		OnRatelimit: func(r *types.RateLimit) {
			fmt.Println("Ratelimited:", r)
		},
		RetryOnRatelimit: true,
		Pester:           pester.New(),
		OnMarshal: []func(r *RequestData, v any) error{
			Cacher,
		},
		SharedState: state,
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
	Cookies []http.Cookie

	// Initial response, if any
	InitialResp *T

	// Sequence number for this request, internal
	sequence int

	// Ratelimit bucket, internal
	bucket *ratelimits.Bucket
}

// Request data from API, but not generic so can be used in OnMarshal and other functions
//
// In addition to typical Request data, this also contains config.Config
type RequestData struct {
	Method  Method
	Path    string
	Json    any
	Headers map[string]string
	Config  *RestConfig
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
