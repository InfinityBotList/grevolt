package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Creates an invite to this channel.
//
// Channel must be a TextChannel.
func (c *RestClient) CreateInvite(target string) (*types.CreateInviteResponseInvite, error) {
	return rest.Request[types.CreateInviteResponseInvite]{Method: rest.POST, Path: "channels/" + target + "/invites"}.With(&c.Config)
}
