package events

type ChannelGroupJoin struct {
	Event

	// Channel Id
	Id string `json:"id"`

	// User Id
	UserId string `json:"user"`
}
