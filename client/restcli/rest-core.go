package restcli

import (
	"github.com/infinitybotlist/grevolt/rest/clientapi"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch the server configuration for this Revolt instance.
func (c *RestClient) QueryNode() (*types.RevoltConfig, *types.APIError, error) {
	var cfg *types.RevoltConfig
	apiErr, err := clientapi.NewReq(&c.Config).Get("/").DoAndMarshal(&cfg)
	return cfg, apiErr, err
}
