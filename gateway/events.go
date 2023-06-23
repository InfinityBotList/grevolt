package gateway

import (
	"fmt"

	"github.com/infinitybotlist/grevolt/types/events"
)

func createEvent[T events.EventInterface](
	w *GatewayClient,
	data []byte,
	fn func(w *GatewayClient, evt *T),
) error {
	var evtMarshalled *T

	err := w.Recieve(data, &evtMarshalled)

	if err != nil {
		return err
	}

	if fn == nil {
		return nil
	}

	fn(w, evtMarshalled)

	return nil
}

func (w *GatewayClient) handleEvent(event []byte, typ string) {
	fmt.Println(string(event))

	var err error
	switch typ {
	case "Authenticated":
		err = createEvent[events.Authenticated](w, event, w.EventHandlers.Authenticated)
	case "Ready":
		err = createEvent[events.Ready](w, event, w.EventHandlers.Ready)
	}

	if err != nil {
		w.Logger.Error(err)
	}
}

// Event handler for the websocket
type EventHandlers struct {
	// The server has authenticated your connection and you will shortly start receiving data.
	Authenticated func(w *GatewayClient, e *events.Authenticated)

	// Data for use by client, data structures match the API specification
	Ready func(w *GatewayClient, e *events.Ready)
}
