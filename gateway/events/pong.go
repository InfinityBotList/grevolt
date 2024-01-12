package events

type Pong struct {
	Event
	Data int64 `json:"data"` // unix timestamp
}
