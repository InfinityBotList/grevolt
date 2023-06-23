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

	err := w.Decode(data, &evtMarshalled)

	if err != nil {
		return err
	}

	if fn == nil {
		return nil
	}

	fn(w, evtMarshalled)

	return nil
}

func (w *GatewayClient) HandleBulk(event []byte) error {
	var bulkData *events.Bulk

	err := w.Decode(event, &bulkData)

	if err != nil {
		return err
	}

	for _, evt := range bulkData.V {
		bytes, err := w.Encode(evt)

		if err != nil {
			return err
		}

		typ, ok := evt["type"]

		if !ok {
			w.Logger.Error("event has no type", string(bytes))
		}

		typStr, ok := typ.(string)

		if !ok {
			w.Logger.Error("event type is not a string", string(bytes))
		}

		w.HandleEvent(bytes, typStr)
	}

	return nil
}

func (w *GatewayClient) HandleEvent(event []byte, typ string) {
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
	case "Bulk":
		// Bulk is a bit unique, special handling is required
		err = w.HandleBulk(event)
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
