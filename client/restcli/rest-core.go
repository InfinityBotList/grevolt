package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch the server configuration for this Revolt instance.
func (c *RestClient) QueryNode() (*types.RevoltConfig, *types.APIError, error) {
	var cfg *types.RevoltConfig
	apiErr, err := rest.NewReq(&c.Config).Get("/").DoAndMarshal(&cfg)
	return cfg, apiErr, err
}
