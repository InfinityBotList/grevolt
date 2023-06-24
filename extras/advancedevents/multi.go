package advancedevents

import (
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/types/events"
)

type Multi[T events.EventInterface] struct {
	Funcs []func(w *gateway.GatewayClient, e *T)
}

// Add a event handler function to a multi
func (m *Multi[T]) Add(f func(w *gateway.GatewayClient, e *T)) {
	m.Funcs = append(m.Funcs, f)
}

// Creates a new multi event handler for running multiple event handlers in a single event
//
// "We truly live in a world of wonders" - Tim Cook
func NewMulti[T events.EventInterface](evts ...func(w *gateway.GatewayClient, e *T)) *Multi[T] {
	return &Multi[T]{
		Funcs: evts,
	}
}

func (m *Multi[T]) Build() func(w *gateway.GatewayClient, e *T) {
	return func(w *gateway.GatewayClient, e *T) {
		for _, f := range m.Funcs {
			f(w, e)
		}
	}
}
