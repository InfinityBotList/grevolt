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
func (c *RestClient) FetchServer(serverID string) (*types.Server, error) {
	return rest.Request[types.Server]{Path: "servers/" + serverID}.With(&c.Config)
}

// Deletes a server if owner otherwise leaves.
func (c *RestClient) DeleteOrLeaveServer(serverID string, leaveSilently bool) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "servers/" + serverID + "?leave_silently=" + boolean(leaveSilently)}.NoContent(&c.Config)
}

// Edit a server by its id.
func (c *RestClient) EditServer(serverID string, d *types.DataEditServer) (*types.Server, error) {
	return rest.Request[types.Server]{Method: rest.PATCH, Path: "servers/" + serverID, Json: d}.With(&c.Config)
}

// Mark all channels in a server as read.
func (c *RestClient) MarkServerAsRead(serverID string) error {
	return rest.Request[types.APIError]{Method: rest.PUT, Path: "servers/" + serverID + "/ack"}.NoContent(&c.Config)
}

// Create a new Text or Voice channel.
func (c *RestClient) CreateChannel(serverID string, d *types.DataCreateChannel) (*types.Channel, error) {
	return rest.Request[types.Channel]{Method: rest.POST, Path: "servers/" + serverID + "/channels", Json: d}.With(&c.Config)
}
