package events

import "github.com/infinitybotlist/grevolt/types"

type ServerUpdate struct {
	Event

	// Server ID
	Id string `json:"id"`

	// Partial channel object, not all data is available
	//
	// Exactly which fields are available is subject to change and thus not documented.
	Data *types.Server `json:"data"`

	// Clear is a field to remove, one of Icon/Banner/Description
	//
	// Grevolt plays it safe here, and uses the FieldsServer type already used
	// throughout the API just in case the docs are out-of-date.
	//
	// This does not affect users, as it expands and not reduces the possible
	// values.
	Clear []types.FieldsServer `json:"clear"`
}
