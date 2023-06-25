package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch the server configuration for this Revolt instance.
func (c *RestClient) QueryNode() (*types.RevoltConfig, error) {
	return rest.Request[types.RevoltConfig]{Path: "/"}.With(&c.Config)
}
