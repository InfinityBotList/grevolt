package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Retrieve a list of mutual friends and servers with another user.
func (c *RestClient) FetchMutualFriendsAndServers(target string) (*types.MutualResponse, error) {
	return rest.Request[types.MutualResponse]{Path: "users/" + target + "/mutual"}.With(&c.Config)
}

// Accept another user's friend request
func (c *RestClient) AcceptFriendRequest(target string) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.PUT, Path: "users/" + target + "/friend"}.With(&c.Config)
}

// Denies another user's friend request or removes an existing friend.
func (c *RestClient) DenyFriendRequestOrRemoveFriend(target string) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.DELETE, Path: "users/" + target + "/friend"}.With(&c.Config)
}

// Block another user by their id.
func (c *RestClient) BlockUser(target string) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.PUT, Path: "users/" + target + "/block"}.With(&c.Config)
}

// Block another user by their id.
func (c *RestClient) UnblockUser(target string) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.DELETE, Path: "users/" + target + "/block"}.With(&c.Config)
}

// Send a friend request to another user.
func (c *RestClient) SendFriendRequest(d *types.DataSendFriendRequest) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.POST, Path: "users/friend", Json: d}.With(&c.Config)
}
