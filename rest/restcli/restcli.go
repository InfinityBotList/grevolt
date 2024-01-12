package restcli

import (
	"github.com/infinitybotlist/grevolt/cache/state"
	"github.com/infinitybotlist/grevolt/rest"
)

type RestClient struct {
	Config rest.RestConfig
}

// DefaultRestClient returns a new rest client with the default configuration
func DefaultRestClient(state *state.State) *RestClient {
	return &RestClient{
		Config: rest.DefaultRestConfig(state),
	}
}

// Helper methood for ternary
func ternary(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

// Helper method to shorten the common case of boolean to string conversion
func boolean(condition bool) string {
	return ternary(condition, "true", "false")
}

func runIf(condition bool, f func()) {
	if condition {
		f()
	}
}
