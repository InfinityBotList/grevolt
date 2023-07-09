package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch channel by its id.
//
// <target is the channel id>
func (c *RestClient) FetchChannel(target string) (*types.Channel, error) {
	return rest.Request[types.Channel]{Path: "channels/" + target}.With(&c.Config)
}

// Deletes a server channel, leaves a group or closes a group.
//
// <target is the channel id>
func (c *RestClient) CloseChannel(target string, leaveSilently bool) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "channels/" + target + "?leave_silently=" + boolean(leaveSilently)}.NoContent(&c.Config)
}

// Edit a channel object by its id.
//
// <target is the channel id>
func (c *RestClient) EditChannel(target string, d *types.DataEditChannel) (*types.Channel, error) {
	return rest.Request[types.Channel]{Method: rest.PATCH, Path: "channels/" + target, Json: d}.With(&c.Config)
}
