package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Create a new server.
func (c *RestClient) CreateServer(d *types.DataCreateServer) (*types.CreateServerResponse, error) {
	return rest.Request[types.CreateServerResponse]{Method: rest.POST, Path: "servers/create", Json: d}.With(&c.Config)
}

// Fetch a server by its id.
//
// <target is the server id>
func (c *RestClient) FetchServer(target string) (*types.Server, error) {
	return rest.Request[types.Server]{Path: "servers/" + target}.With(&c.Config)
}

// Deletes a server if owner otherwise leaves.
//
// <target is the server id>
func (c *RestClient) DeleteOrLeaveServer(target string, leaveSilently bool) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "servers/" + target + "?leave_silently=" + boolean(leaveSilently)}.NoContent(&c.Config)
}

// Edit a server by its id.
//
// <target is the server id>
func (c *RestClient) EditServer(target string, d *types.DataEditServer) (*types.Server, error) {
	return rest.Request[types.Server]{Method: rest.PATCH, Path: "servers/" + target, Json: d}.With(&c.Config)
}

// Mark all channels in a server as read.
//
// <target is the server id>
func (c *RestClient) MarkServerAsRead(target string) error {
	return rest.Request[types.APIError]{Method: rest.PUT, Path: "servers/" + target + "/ack"}.NoContent(&c.Config)
}

// Create a new Text or Voice channel.
//
// <target is the server id>
func (c *RestClient) CreateChannel(target string, d *types.DataCreateChannel) (*types.Channel, error) {
	return rest.Request[types.Channel]{Method: rest.POST, Path: "servers/" + target + "/channels", Json: d}.With(&c.Config)
}
