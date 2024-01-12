package events

type EventInterface interface {
	EventType() string
}

type Event struct {
	Type string `json:"type"`
}

func (e Event) EventType() string {
	return e.Type
}
