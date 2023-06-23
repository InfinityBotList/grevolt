package events

import "github.com/infinitybotlist/grevolt/types"

type ChannelUpdate struct {
	Event

	// Channel ID
	Id string `json:"id"`

	// Partial channel object, not all data is available
	//
	// Exactly which fields are available is subject to change and thus not documented.
	Data *types.Channel `json:"data"`

	// Clear is a field to remove, one of Icon/Description
	//
	// Grevolt plays it safe here, and uses the ChannelField type already used
	// throughout the API just in case the docs are out-of-date.
	//
	// This does not affect users, as it expands and not reduces the possible
	// values.
	Clear []types.FieldsChannel `json:"clear"`
}
