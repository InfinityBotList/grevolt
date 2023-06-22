package client

import (
	"errors"
	"time"

	"github.com/infinitybotlist/grevolt/client/geneva"
	"github.com/infinitybotlist/grevolt/client/restcli"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/rest/restconfig"
)

type Client struct {
	Rest      restcli.RestClient
	Websocket gateway.GatewayClient
}

// New returns a new client with default options
func New() *Client {
	return &Client{
		Rest: restcli.RestClient{
			Config: restconfig.DefaultRestConfig,
		},
		Websocket: gateway.GatewayClient{
			APIVersion: "1",
			Timeout:    10 * time.Second,
		},
	}
}

// Authorizes to both rest and websocket (websocket not implemented yet)
func (c *Client) Authorize(token *geneva.Token) {
	// Rest client
	c.Rest.Config.SessionToken = token
	c.Websocket.SessionToken = token
}

// Opens a websocket connection. Not mandatory for REST requests
func (c *Client) Open() error {
	// Fetch the websocket URL
	cfg, apiError, err := c.Rest.QueryNode()

	if err != nil {
		return err
	}

	if apiError != nil {
		return errors.New(apiError.Type())
	}

	// Set the websocket URL
	c.Websocket.WSUrl = cfg.Ws

	if c.Websocket.Logger == nil {
		c.Websocket.Logger = c.Rest.Config.Logger
	}

	return c.Websocket.Open()
}
