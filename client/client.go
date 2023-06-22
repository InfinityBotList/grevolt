package client

import (
	"github.com/infinitybotlist/grevolt/client/geneva"
	"github.com/infinitybotlist/grevolt/client/restcli"
	"github.com/infinitybotlist/grevolt/rest/restconfig"
)

type Client struct {
	Rest restcli.RestClient
}

// New returns a new client with default options
func New() *Client {
	return &Client{
		Rest: restcli.RestClient{
			Config: restconfig.DefaultRestConfig,
		},
	}
}

// Authorizes to both rest and websocket (websocket not implemented yet)
func (c *Client) Authorize(token *geneva.Token) {
	// Rest client
	c.Rest.Config.SessionToken = token
}
