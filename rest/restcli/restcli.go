package restcli

import "github.com/infinitybotlist/grevolt/rest"

type RestClient struct {
	Config rest.RestConfig
}

// Helper methood for ternary
func ternary(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
