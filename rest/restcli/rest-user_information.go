package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Retrieve your user information.
func (c *RestClient) FetchSelf() (*types.User, error) {
	return rest.Request[types.User]{Path: "users/@me"}.With(&c.Config)
}

// Retrieve a user's information <given their id, requires a mutual server/group>
func (c *RestClient) FetchUser(target string) (*types.User, error) {
	return rest.Request[types.User]{Path: "users/" + target}.With(&c.Config)
}

// Edit currently authenticated user <given their id and the new user information>
func (c *RestClient) EditUser(target string, user *types.DataEditUser) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.PATCH, Path: "users/" + target, Json: user}.With(&c.Config)
}

// Retrieve a user's flags <given their id, these flags can be checked using flags.HasFlag>.
func (c *RestClient) FetchUserFlags(target string) (*types.FlagResponse, error) {
	return rest.Request[types.FlagResponse]{Path: "users/" + target + "/flags"}.With(&c.Config)
}

// Change your username <untested as it needs a password>
func (c *RestClient) ChangeUsername(d *types.DataChangeUsername) (*types.User, error) {
	return rest.Request[types.User]{Method: rest.PATCH, Path: "users/@me/username", Json: d}.With(&c.Config)
}

// This returns a default avatar based on the given id.
func (c *RestClient) FetchDefaultAvatar(target string) (*rest.Bytes, error) {
	return rest.Request[rest.Bytes]{Path: "users/" + target + "/default_avatar"}.With(&c.Config)
}

// Retrieve a user's information <given their id, requires a mutual server/group>
func (c *RestClient) FetchUserProfile(target string) (*types.UserProfile, error) {
	return rest.Request[types.UserProfile]{Path: "users/" + target + "/profile"}.With(&c.Config)
}
