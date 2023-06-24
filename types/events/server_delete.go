package events

type ServerDelete struct {
	Event

	// Server Id
	Id string `json:"id"`
}
