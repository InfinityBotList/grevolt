package restcli

import (
	"github.com/infinitybotlist/grevolt/rest/clientapi"
	"github.com/infinitybotlist/grevolt/types"
)

// Creates an invite to this channel.
//
// Channel must be a TextChannel.
func (c *RestClient) CreateInvite(target string) (*types.CreateInviteResponseInvite, *types.APIError, error) {
	var i *types.CreateInviteResponseInvite
	apiErr, err := clientapi.NewReq(&c.Config).Post("/channels/" + target + "/invites").DoAndMarshal(&i)
	return i, apiErr, err
}
