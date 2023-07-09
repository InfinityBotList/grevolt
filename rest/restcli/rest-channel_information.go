package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch channel by its id.
func (c *RestClient) FetchChannel(channelID string) (*types.Channel, error) {
	return rest.Request[types.Channel]{Path: "channels/" + channelID}.With(&c.Config)
}

// Deletes a server channel, leaves a group or closes a group.
func (c *RestClient) CloseChannel(channelID string, leaveSilently bool) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "channels/" + channelID + "?leave_silently=" + boolean(leaveSilently)}.NoContent(&c.Config)
}

// Edit a channel object by its id.
func (c *RestClient) EditChannel(channelID string, d *types.DataEditChannel) (*types.Channel, error) {
	return rest.Request[types.Channel]{Method: rest.PATCH, Path: "channels/" + channelID, Json: d}.With(&c.Config)
}
