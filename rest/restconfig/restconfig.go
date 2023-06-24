// Package restconfig defines a configuration for REST requests
package restconfig

import (
	"fmt"
	"time"

	"github.com/infinitybotlist/grevolt/client/auth"
	"github.com/infinitybotlist/grevolt/rest/clientapi/ratelimits"
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
