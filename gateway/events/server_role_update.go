package events

import "github.com/infinitybotlist/grevolt/types"

type ServerRoleUpdate struct {
	Event

	// Server ID
	Id string `json:"id"`

	// Role ID
	RoleId string `json:"role_id"`

	// Partial channel object, not all data is available
	//
	// Exactly which fields are available is subject to change and thus not documented.
	Data *types.Role `json:"data"`

	// Clear is a field to remove, one of Colour
	//
	// Grevolt plays it safe here, and uses the FieldsRole type already used
	// throughout the API just in case the docs are out-of-date.
	//
	// This does not affect users, as it expands and not reduces the possible
	// values.
	Clear []types.FieldsRole `json:"clear"`
}
