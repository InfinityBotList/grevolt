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

	if typ == "Bulk" {
		// Bulk is a bit unique, special handling is required
		err := w.HandleBulk(event)

		if err != nil {
			w.Logger.Error(err)
		}
	}

	var err error
	switch typ {
	case "Error":
		err = createEvent[events.Error](w, event, w.EventHandlers.Error)
	case "Authenticated":
		err = createEvent[events.Authenticated](w, event, w.EventHandlers.Authenticated)
	case "Bulk":
		err = createEvent[events.Bulk](w, event, w.EventHandlers.Bulk)
	case "Pong":
		err = createEvent[events.Pong](w, event, w.EventHandlers.Pong)
	case "Ready":
		err = createEvent[events.Ready](w, event, w.EventHandlers.Ready)
	case "Message":
		err = createEvent[events.Message](w, event, w.EventHandlers.Message)
	case "MessageUpdate":
		err = createEvent[events.MessageUpdate](w, event, w.EventHandlers.MessageUpdate)
	case "MessageAppend":
		err = createEvent[events.MessageAppend](w, event, w.EventHandlers.MessageAppend)
	case "MessageDelete":
		err = createEvent[events.MessageDelete](w, event, w.EventHandlers.MessageDelete)
	case "MessageReact":
		err = createEvent[events.MessageReact](w, event, w.EventHandlers.MessageReact)
	case "MessageUnreact":
		err = createEvent[events.MessageUnreact](w, event, w.EventHandlers.MessageUnreact)
	case "MessageRemoveReaction":
		err = createEvent[events.MessageRemoveReaction](w, event, w.EventHandlers.MessageRemoveReaction)
	}

	if err != nil {
		w.Logger.Error(err)
	}
}

// Event handler for the websocket
type EventHandlers struct {
	// Not an actual revolt event, this is a sink that allows you to provide a function for raw event handling
	RawSinkFunc func(w *GatewayClient, data []byte, typ string)

	// An error occurred which meant you couldn't authenticate.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Error func(w *GatewayClient, e *events.Error)

	// The server has authenticated your connection and you will shortly start receiving data.
	Authenticated func(w *GatewayClient, e *events.Authenticated)

	// Several events have been sent, process each item of v as its own event.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Bulk func(w *GatewayClient, e *events.Bulk)

	// Ping response from the server.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Pong func(w *GatewayClient, e *events.Pong)

	// Data for use by client, data structures match the API specification
	Ready func(w *GatewayClient, e *events.Ready)

	// Message received, the event object has the same schema as the Message object in the API with the addition of an event type.
	Message func(w *GatewayClient, e *events.Message)

	// Message edited or otherwise updated.
	MessageUpdate func(w *GatewayClient, e *events.MessageUpdate)

	// Message has data being appended to it.
	MessageAppend func(w *GatewayClient, e *events.MessageAppend)

	// Message has been deleted.
	MessageDelete func(w *GatewayClient, e *events.MessageDelete)

	// A reaction has been added to a message.
	MessageReact func(w *GatewayClient, e *events.MessageReact)

	// A reaction has been removed from a message.
	MessageUnreact func(w *GatewayClient, e *events.MessageUnreact)

	// A certain reaction has been removed from the message.
	//
	// <the difference between this and MessageUnreact is that
	// this event is sent when a user with manage messages removes
	// a reaction while MessageUnreact is sent when a user removes
	// their own reaction>
	MessageRemoveReaction func(w *GatewayClient, e *events.MessageRemoveReaction)
}
