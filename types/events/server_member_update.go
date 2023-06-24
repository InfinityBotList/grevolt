package events

import "github.com/infinitybotlist/grevolt/types"

type ServerMemberUpdateData struct {
	// Server Id
	ServerId string `json:"server"`

	// User Id
	//
	// <in reality, the member id>
	UserId string `json:"user"`
}

type ServerMemberUpdate struct {
	Event

	// Ids
	Id *ServerMemberUpdateData `json:"id"`

	// Partial server member object, not all data is available
	//
	// Exactly which fields are available is subject to change and thus not documented.
	Data *types.Member `json:"data"`

	// Clear is a field to remove, one of Nickname/Avatar
	//
	// Grevolt plays it safe here, and uses the FieldsServer type already used
	// throughout the API just in case the docs are out-of-date.
	//
	// This does not affect users, as it expands and not reduces the possible
	// values.
	Clear []types.FieldsMember `json:"clear"`
}
