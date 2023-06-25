package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Retrieve your user information.
func (c *RestClient) FetchSelf() (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := rest.NewReq(&c.Config).Get("users/@me").DoAndMarshal(&u)
	return u, apiErr, err
}

// Retrieve a user's information <given their id, requires a mutual server/group>
func (c *RestClient) FetchUser(target string) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := rest.NewReq(&c.Config).Get("users/" + target).DoAndMarshal(&u)
	return u, apiErr, err
}

// Edit currently authenticated user <given their id and the new user information>
func (c *RestClient) EditUser(target string, user *types.DataEditUser) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := rest.NewReq(&c.Config).Patch("users/" + target).Json(&user).DoAndMarshal(&u)
	return u, apiErr, err
}

// Retrieve a user's flags <given their id, these flags can be checked using flags.HasFlag>.
func (c *RestClient) FetchUserFlags(target string) (*types.FlagResponse, *types.APIError, error) {
	var f *types.FlagResponse
	apiErr, err := rest.NewReq(&c.Config).Get("users/" + target + "/flags").DoAndMarshal(&f)
	return f, apiErr, err
}

// Change your username <untested as it needs a password>
func (c *RestClient) ChangeUsername(d *types.DataChangeUsername) (*types.User, *types.APIError, error) {
	var u *types.User
	apiErr, err := rest.NewReq(&c.Config).Patch("users/@me/username").Json(&d).DoAndMarshal(&u)
	return u, apiErr, err
}

// This returns a default avatar based on the given id.
func (c *RestClient) FetchDefaultAvatar(target string) ([]byte, *types.APIError, error) {
	bytes, apiErr, err := rest.NewReq(&c.Config).Get("users/" + target + "/default_avatar").DoAndMarshalBytes()
	return bytes, apiErr, err
}

// Retrieve a user's information <given their id, requires a mutual server/group>
func (c *RestClient) FetchUserProfile(target string) (*types.UserProfile, *types.APIError, error) {
	var u *types.UserProfile
	apiErr, err := rest.NewReq(&c.Config).Get("users/" + target + "/profile").DoAndMarshal(&u)
	return u, apiErr, err
}
