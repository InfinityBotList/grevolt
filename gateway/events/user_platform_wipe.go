package events

type UserPlatformWipe struct {
	Event

	// User Id
	UserId string `json:"user_id"`

	// <officially, this is a i32, but we use i64 to make room for future flags and for consistency>
	Flags int64
}
