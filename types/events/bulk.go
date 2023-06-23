package events

type Bulk struct {
	Event
	V []map[string]any `json:"v"`
}
