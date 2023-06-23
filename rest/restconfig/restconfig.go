// Package restconfig defines a configuration for REST requests
package restconfig

import (
	"time"

	"github.com/infinitybotlist/grevolt/client/auth"
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
}

// DefaultRestConfig is the default configuration for the client
var DefaultRestConfig = RestConfig{
	APIUrl:  "https://api.revolt.chat/",
	Timeout: 10 * time.Second,
}
