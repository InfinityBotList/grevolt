package gateway

import (
	"errors"

	"github.com/infinitybotlist/grevolt/types/events"
	"go.uber.org/zap"
)

type EventContext struct {
	// Raw event data
	Raw []byte
}

type Event[T events.EventInterface] func(w *GatewayClient, ctx *EventContext, evt *T)

// Emits an event to a function
//
// +unstable
func CreateEvent[T events.EventInterface](
	w *GatewayClient,
	data []byte,
	fn Event[T],
) (*T, error) {
	var evtMarshalled *T

	err := w.Decode(data, &evtMarshalled)

	if err != nil {
		return nil, errors.New("decode error: " + err.Error())
	}

	if fn == nil {
		return evtMarshalled, nil
	}

	fn(w, &EventContext{
		Raw: data,
	}, evtMarshalled)

	return evtMarshalled, nil
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
			w.Logger.Error(
				"event has no type",
				zap.Binary("evt", bytes),
			)
		}

		typStr, ok := typ.(string)

		if !ok {
			w.Logger.Error(
				"event type is not a string",
				zap.Binary("evt", bytes),
			)
		}

		w.HandleEvent(bytes, typStr)
	}

	return nil
}

func (w *GatewayClient) HandleAuth(event []byte) error {
	var authData *events.Auth

	err := w.Decode(event, &authData)

	if err != nil {
		return err
	}

	switch authData.Type {
	case "DeleteSession":
		_, err = CreateEvent[events.Auth_DeleteSession](w, event, w.EventHandlers.Auth_DeleteSession)
		return err
	case "DeleteAllSessions":
		_, err = CreateEvent[events.Auth_DeleteAllSessions](w, event, w.EventHandlers.Auth_DeleteAllSessions)
		return err
	default:
		w.Logger.Warn(
			"Unknown auth event type",
			zap.String("eventType", authData.Type),
		)
	}

	return nil
}

func (w *GatewayClient) HandleEvent(event []byte, typ string) {
	if w.RawSinkFunc != nil && len(w.RawSinkFunc) > 0 {
		for _, fn := range w.RawSinkFunc {
			fn(w, event, typ)
		}
	}

	// Special event handling
	var err error
	switch typ {
	case "Bulk":
		// For bulk, decode individual events and handle them
		err = w.HandleBulk(event)
	case "Auth":
		// Auth is a bit unique because of event it, handle it
		err = w.HandleAuth(event)
	}

	if err != nil {
		w.Logger.Error(
			"special handler failed",
			zap.String("eventType", typ),
			zap.Error(err),
		)
	}

	evt, err := w.EventHandlers.handle(w, event, typ)

	if err != nil {
		w.Logger.Error(
			"Event handling failed",
			zap.Error(err),
			zap.String("type", typ),
		)
		return
	}

	// Handle caching here
	if evt != nil && !w.DisableCache {
		go func() {
			err := w.cacheEvent(typ, evt)

			if err != nil {
				w.Logger.Error(
					"Failed to cache event",
					zap.Error(err),
					zap.String("type", typ),
				)
			}
		}()
	}
}
