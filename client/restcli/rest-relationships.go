package restcli

import (
	"github.com/infinitybotlist/grevolt/rest/clientapi"
	"github.com/infinitybotlist/grevolt/types"
)

// Retrieve a list of mutual friends and servers with another user.
func (c *RestClient) FetchMutualFriendsAndServers(target string) (*types.MutualResponse, *types.APIError, error) {
	var mr *types.MutualResponse
	apiErr, err := clientapi.NewReq(&c.Config).Get("users/" + target + "/mutual").DoAndMarshal(&mr)
	return mr, apiErr, err
}

// Accept another user's friend request
func (c *RestClient) AcceptFriendRequest(target string) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Put("users/" + target + "/friend").DoAndMarshal(&u)
	return u, apiErr, err
}

// Denies another user's friend request or removes an existing friend.
func (c *RestClient) DenyFriendRequestOrRemoveFriend(target string) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Delete("users/" + target + "/friend").DoAndMarshal(&u)
	return u, apiErr, err
}

// Block another user by their id.
func (c *RestClient) BlockUser(target string) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Put("users/" + target + "/block").DoAndMarshal(&u)
	return u, apiErr, err
}

// Block another user by their id.
func (c *RestClient) UnblockUser(target string) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Delete("users/" + target + "/block").DoAndMarshal(&u)
	return u, apiErr, err
}

// Send a friend request to another user.
func (c *RestClient) SendFriendRequest(d *types.DataSendFriendRequest) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Post("users/friend").Json(d).DoAndMarshal(&u)
	return u, apiErr, err
}
