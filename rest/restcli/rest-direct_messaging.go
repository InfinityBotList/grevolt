package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// This fetches your direct messages, including any DM and group DM conversations.
func (c *RestClient) FetchDirectMessageChannels() (*types.ChannelList, error) {
	return rest.Request[types.ChannelList]{Path: "users/dms"}.With(&c.Config)
}

// Open a DM with another user.
//
// If the target is oneself, a saved messages channel is returned.
func (c *RestClient) OpenDirectMessage(target string) (*types.Channel, error) {
	return rest.Request[types.Channel]{Path: "users/" + target + "/dm"}.With(&c.Config)
}
