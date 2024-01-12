package events

import "github.com/infinitybotlist/grevolt/types"

type UserRelationship struct {
	Event

	// Your user Id
	//
	// <the your above is very important, its the currently logged in user id, not the
	// new relationship's user id>
	Id string `json:"id"`

	// <not well documented, its likely the new relationship's user object>
	User *types.User `json:"user"`

	// <relationship status, same as API>
	//
	// <in source code, this is mentioned as deprecated, so avoid using this field>
	Relationship types.RelationshipStatus `json:"status"`
}
