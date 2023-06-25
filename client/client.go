package client

import (
	"errors"
	"os"
	"time"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/rest/restcli"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Client struct {
	Rest      *restcli.RestClient
	Websocket *gateway.GatewayClient
}

// New returns a new client with default options
func New() Client {
	w := zapcore.AddSync(os.Stdout)

	var level = zap.InfoLevel
	if os.Getenv("DEBUG") == "true" {
		level = zap.DebugLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		level,
	)

	logger := zap.New(core).Sugar()

	c := Client{
		Rest: &restcli.RestClient{
			Config: rest.DefaultRestConfig(),
		},
		Websocket: &gateway.GatewayClient{
			APIVersion: "1",
			Timeout:    10 * time.Second,
		},
	}

	c.Rest.Config.Logger = logger.Named("rest")
	c.Rest.Config.Ratelimiter.Logger = logger.Named("ratelimiter")
	c.Websocket.Logger = logger.Named("websocket").Desugar()

	return c
}

// Authorizes to both rest and websocket (websocket not implemented yet)
func (c *Client) Authorize(token *auth.Token) {
	// Rest client
	c.Rest.Config.SessionToken = token
	c.Websocket.SessionToken = token
}

// Prepares a websocket client
//
// # This does not open the websocket
//
// Use the Open() method on the websocket to open the websocket
func (c *Client) PrepareWS() error {
	if c.Rest.Config.SessionToken == nil {
		return errors.New("no session token provided")
	}

	// Fetch the websocket URL
	cfg, err := c.Rest.QueryNode()

	if err != nil {
		return err
	}

	// Set the websocket URL
	c.Websocket.WSUrl = cfg.Ws

	if c.Websocket.Logger == nil {
		c.Websocket.Logger = c.Rest.Config.Logger.Desugar()
	}

	return nil
}
