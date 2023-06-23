package events

// Not yet usable, but its getting there
type Error struct {
	Event
	Error string `json:"error"`
}
