package restcli

import (
	"github.com/infinitybotlist/grevolt/rest/clientapi"
	"github.com/infinitybotlist/grevolt/types"
)

// Retrieve your user information.
func (c *RestClient) FetchSelf() (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Get("users/@me").DoAndMarshal(&u)
	return u, apiErr, err
}

// Retrieve a user's information <given their id>
func (c *RestClient) FetchUser(id string) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Get("users/" + id).DoAndMarshal(&u)
	return u, apiErr, err
}

// Edit currently authenticated user <given their id and the new user information>
func (c *RestClient) EditUser(id string, user *types.DataEditUser) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := clientapi.NewReq(&c.Config).Patch("users/" + id).Json(&user).DoAndMarshal(&u)
	return u, apiErr, err
}

// Retrieve a user's flags <given their id, these flags can be checked using flags.HasFlag>.
func (c *RestClient) FetchUserFlags(id string) (*types.FlagResponse, *types.APIError, error) {
	var f *types.FlagResponse
	apiErr, err := clientapi.NewReq(&c.Config).Get("users/" + id + "/flags").DoAndMarshal(&f)
	return f, apiErr, err
}
