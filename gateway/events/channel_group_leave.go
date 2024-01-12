package events

type ChannelGroupLeave struct {
	Event

	// Channel Id
	Id string `json:"id"`

	// User Id
	UserId string `json:"user"`
}
