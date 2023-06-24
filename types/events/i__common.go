package events

type EventInterface interface {
	EventType() string
	GetEvent() *Event
}

type Event struct {
	Raw  []byte `json:"-"`
	Type string `json:"type"`
}

func (e Event) EventType() string {
	return e.Type
}

func (e Event) GetEvent() *Event {
	return &e
}
