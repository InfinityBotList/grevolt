package events

// An auth event, its special because of the event type stuff
type Auth struct {
	Event

	Type string `json:"event_type"`
}

type Auth_DeleteSession struct {
	Auth

	// User Id
	Id string `json:"user_id"`

	// Session Id
	SessionId string `json:"session_id"`
}

type Auth_DeleteAllSessions struct {
	Auth

	// User Id
	Id string `json:"user_id"`

	// <Session Id to exclude from deletion>
	ExcludeSessionId string `json:"exclude_session_id"`
}
