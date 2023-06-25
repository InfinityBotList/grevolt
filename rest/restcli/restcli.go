package restcli

import (
	"github.com/infinitybotlist/grevolt/cache/state"
	"github.com/infinitybotlist/grevolt/rest"
)

type RestClient struct {
	Config      rest.RestConfig
	SharedState *state.State
}

// Helper methood for ternary
func ternary(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func runIf(condition bool, f func()) {
	if condition {
		f()
	}
}
