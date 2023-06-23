package restcli

import (
	"github.com/infinitybotlist/grevolt/rest/clientapi"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch channel by its id.
func (c *RestClient) FetchChannel(channelID string) (*types.Channel, *types.APIError, error) {
	var d *types.Channel
	apiErr, err := clientapi.NewReq(&c.Config).Get("channels/" + channelID).DoAndMarshal(&d)
	return d, apiErr, err
}

// Deletes a server channel, leaves a group or closes a group.
func (c *RestClient) CloseChannel(channelID string, leaveSilently bool) (*types.APIError, error) {
	var ls string

	if leaveSilently {
		ls = "true"
	} else {
		ls = "false"
	}

	apiErr, err := clientapi.NewReq(&c.Config).Delete("channels/" + channelID + "?leave_silently=" + ls).NoContentErr()
	return apiErr, err
}

// Edit a channel object by its id.
func (c *RestClient) EditChannel(channelID string, d *types.DataEditChannel) (*types.Channel, *types.APIError, error) {
	var dc *types.Channel
	apiErr, err := clientapi.NewReq(&c.Config).Patch("channels/" + channelID).Json(d).DoAndMarshal(&dc)
	return dc, apiErr, err
}
