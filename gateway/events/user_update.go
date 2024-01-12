package events

import "github.com/infinitybotlist/grevolt/types"

type UserUpdate struct {
	Event

	// User Id
	Id string `json:"id"`

	// Partial channel object, not all data is available
	//
	// Exactly which fields are available is subject to change and thus not documented.
	Data *types.User `json:"data"`

	// Clear is a field to remove, one of ProfileContent/ProfileBackground/StatusText/Avatar
	//
	// Grevolt plays it safe here, and uses the FieldsServer type already used
	// throughout the API just in case the docs are out-of-date.
	//
	// This does not affect users, as it expands and not reduces the possible
	// values.
	Clear []types.FieldsUser `json:"clear"`

	// <undocumented, may exist???>
	EventId string `json:"event_id,omitempty"`
}
