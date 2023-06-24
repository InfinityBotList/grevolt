package events

// <undocumented, will likely be available in a future release>
type UserSettingsUpdate struct {
	Event

	// User Id
	UserId string `json:"user_id"`

	// User settings
	//
	// <until better documentation is released, you're on your own>
	Update map[string]any `json:"update"`
}
