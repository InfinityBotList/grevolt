package events

type ChannelAck struct {
	Event

	// Channel Id
	Id string `json:"id"`

	// User Id
	UserId string `json:"user"`

	// Message id
	MessageId string `json:"message_id"`
}
