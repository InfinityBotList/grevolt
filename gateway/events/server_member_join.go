package events

type ServerMemberJoin struct {
	Event

	// Server Id
	Id string `json:"id"`

	// User Id
	UserId string `json:"user"`
}
