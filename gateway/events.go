package gateway

import (
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
	if w.EventHandlers.RawSinkFunc != nil {
		w.EventHandlers.RawSinkFunc(w, event, typ)
	}

	var err error
	switch typ {
	case "Authenticated":
		err = createEvent[events.Authenticated](w, event, w.EventHandlers.Authenticated)
	case "Ready":
		err = createEvent[events.Ready](w, event, w.EventHandlers.Ready)
	case "Error":
		err = createEvent[events.Error](w, event, w.EventHandlers.Error)
	}

	if err != nil {
		w.Logger.Error(err)
	}
}

// Event handler for the websocket
type EventHandlers struct {
	// Not an actual revolt event, this is a sink that allows you to provide a function for raw event handling
	RawSinkFunc func(w *GatewayClient, data []byte, typ string)

	// The server has authenticated your connection and you will shortly start receiving data.
	Authenticated func(w *GatewayClient, e *events.Authenticated)

	// Data for use by client, data structures match the API specification
	Ready func(w *GatewayClient, e *events.Ready)

	// An error occurred which meant you couldn't authenticate.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Error func(w *GatewayClient, e *events.Error)
}
