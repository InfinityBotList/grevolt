package events

type ServerRoleDelete struct {
	Event

	// Server ID
	Id string `json:"id"`

	// Role ID
	RoleId string `json:"role_id"`
}
