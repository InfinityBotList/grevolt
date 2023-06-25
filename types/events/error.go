package events

type Error struct {
	Event
	Error string `json:"error"`
}
