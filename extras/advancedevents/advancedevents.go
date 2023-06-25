// Package advancedevents provides a much more advanced event and error handling framework
// for events
package advancedevents

import (
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/types/events"
)

type EventHandlerFunc[T events.EventInterface] func(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *T) error

// Event errors
type EventFunc[T events.EventInterface] struct {
	// The ID of the handler
	ID string

	// The handler itself
	Handler EventHandlerFunc[T]

	// Error handlers for this specific event handler
	ErrorHandlers []ErrorHandler[T]
}

type ErrorHandler[T events.EventInterface] func(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *T, err error, handler EventFunc[T])

// An event function, note that events are loaded *syncronously*
// so you should make your event handlers accordingly
type EventHandler[T events.EventInterface] struct {
	// It is recommended to use Add() instead as it is more ergonomic,
	// but this allows you to set handlers
	//
	// If a handler errors, the error will be passed to the error handlers
	//
	// No other handlers will be called on error making it useful for db calls and
	// ``on_command_error`` implementations
	Handlers []EventFunc[T]

	// Global error handler for the event
	//
	// This is per event, not per handler
	GlobalErrorHandlers []ErrorHandler[T]
}

// Creates a new event handler
func NewEventHandler[T events.EventInterface]() *EventHandler[T] {
	return &EventHandler[T]{
		Handlers:            []EventFunc[T]{},
		GlobalErrorHandlers: []ErrorHandler[T]{},
	}
}

func (ef *EventHandler[T]) Build() func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
	return func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
		for _, handler := range ef.Handlers {
			err := handler.Handler(w, ctx, e)

			if err != nil {
				if len(handler.ErrorHandlers) > 0 {
					for _, h := range handler.ErrorHandlers {
						h(w, ctx, e, err, handler)
					}
				}

				// Global handlers
				if len(ef.GlobalErrorHandlers) > 0 {
					for _, h := range ef.GlobalErrorHandlers {
						h(w, ctx, e, err, handler)
					}
				}
			}
		}
	}
}

// Adds a handler for errors
//
// Type is "func(w *GatewayClient, evt *T, err error, handler EventFunc[T])"
func (ef *EventHandler[T]) AddGlobalErrorHandler(fn ErrorHandler[T]) *EventHandler[T] {
	h := append(ef.GlobalErrorHandlers, fn)
	ef.GlobalErrorHandlers = h

	return ef
}

// Adds an simple event handler
func (ef *EventHandler[T]) Add(fns ...EventHandlerFunc[T]) *EventHandler[T] {
	for _, fn := range fns {
		h := append(ef.Handlers, EventFunc[T]{
			ID:      "",
			Handler: fn,
		})
		ef.Handlers = h
	}

	return ef
}

// Adds an event handler
func (ef *EventHandler[T]) AddRaw(fns ...EventFunc[T]) *EventHandler[T] {
	for _, fn := range fns {
		h := append(ef.Handlers, fn)
		ef.Handlers = h
	}

	return ef
}

// Wrap a event on top of potential other non-advancedevent events and friends
func Wrap[T events.EventInterface](evt func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T), funcs ...func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T)) func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
	if evt == nil {
		return func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
			for _, fn := range funcs {
				fn(w, ctx, e)
			}
		}
	}

	return func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
		for _, fn := range funcs {
			fn(w, ctx, e)
		}

		evt(w, ctx, e)
	}
}

// Wrap a event below potential other non-advancedevent events and friends
func WrapEnd[T events.EventInterface](evt func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T), funcs ...func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T)) func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
	if evt == nil {
		return func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
			for _, fn := range funcs {
				fn(w, ctx, e)
			}
		}
	}

	return func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
		evt(w, ctx, e)
		for _, fn := range funcs {
			fn(w, ctx, e)
		}
	}
}
