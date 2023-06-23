package events

// Several events have been sent, process each item of v as its own event.
//
// < the library handles this for you >
type Bulk struct {
	Event
	V []map[string]any `json:"v"`
}
