package gatewaymiddleware

import (
	"fmt"
	"os"

	"github.com/infinitybotlist/grevolt/gateway"
)

// This creates a debug middleware to allow logging of events
// with the “DEBUG“ environment variable set to “true“
func EventDebug(w *gateway.GatewayClient) {
	w.RawSinkFunc = append(
		w.RawSinkFunc,
		func(w *gateway.GatewayClient, data []byte, typ string) {
			if os.Getenv("DEBUG") == "true" {
				fmt.Println(string(data))
			}
		},
	)
}
