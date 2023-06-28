package restcli

import (
	"errors"

	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Retrieves all users who are part of this group.
//
// <target is the group channel id>
func (c *RestClient) FetchGroupMembers(target string) (*types.UserList, error) {
	return rest.Request[types.UserList]{Path: "/channels/" + target + "/members"}.With(&c.Config)
}

// Create a new group channel.
func (c *RestClient) CreateGroup(d *types.DataCreateGroup) (*types.Channel, error) {
	// Basic validation since revolt just outright rejects invalid requests.
	if len(d.Users) == 0 || d.Name == "" {
		return nil, errors.New("name and users are required")
	}

	return rest.Request[types.Channel]{Method: rest.POST, Path: "/channels/create", Json: d}.With(&c.Config)
}

// Adds another user to the group.
//
// <target is the group channel id, member is the member to add to the group>
func (c *RestClient) AddMemberToGroup(target, member string) error {
	return rest.Request[types.APIError]{Method: rest.PUT, Path: "/channels/" + target + "/recipients/" + member}.NoContent(&c.Config)
}

// Removes a user from the group.
func (c *RestClient) RemoveMemberFromGroup(target, member string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "/channels/" + target + "/recipients/" + member}.NoContent(&c.Config)
}
