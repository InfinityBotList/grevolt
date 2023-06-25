package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// This fetches your direct messages, including any DM and group DM conversations.
func (c *RestClient) FetchDirectMessageChannels() (*types.ChannelList, *types.APIError, error) {
	var d *types.ChannelList
	apiErr, err := rest.NewReq(&c.Config).Get("users/dms").DoAndMarshal(&d)
	return d, apiErr, err
}

// Open a DM with another user.
//
// If the target is oneself, a saved messages channel is returned.
func (c *RestClient) OpenDirectMessage(target string) (*types.Channel, *types.APIError, error) {
	var d *types.Channel
	apiErr, err := rest.NewReq(&c.Config).Get("users/" + target + "/dm").DoAndMarshal(&d)
	return d, apiErr, err
}
