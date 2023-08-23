package advancedevents

import (
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/types/events"
)

type Multi[T events.EventInterface] struct {
	Funcs []func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T)
}

// Add a event handler function to a multi
func (m *Multi[T]) Add(f func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T)) {
	m.Funcs = append(m.Funcs, f)
}

// Builds the multi into a list of event handlers
func (m *Multi[T]) Build() func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
	return func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T) {
		for _, f := range m.Funcs {
			f(w, ctx, e)
		}
	}
}

// Creates a new multi event handler for running multiple event handlers in a single event
//
// "We truly live in a world of wonders" - Tim Cook
func NewMulti[T events.EventInterface](evts ...func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *T)) *Multi[T] {
	return &Multi[T]{
		Funcs: evts,
	}
}
