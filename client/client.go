package client

import (
	"github.com/infinitybotlist/grevolt/rest/restconfig"
)

type Client struct {
	Rest RestClient
}

type RestClient struct {
	Config restconfig.RestConfig
}

// New returns a new client with default options
func New() *Client {
	return &Client{
		Rest: RestClient{
			Config: restconfig.DefaultRestConfig,
		},
	}
}

// Authorizes to both rest and websocket (websocket not implemented yet)
func (c *Client) Authorize(token *restconfig.Token) {
	// Rest client
	c.Rest.Config.SessionToken = token
}
