package client

import (
	"os"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/cache/state"
	"github.com/infinitybotlist/grevolt/cache/store/basicstore"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/rest/restcli"
	"github.com/infinitybotlist/grevolt/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Client struct {
	Rest      *restcli.RestClient
	Websocket *gateway.GatewayClient
	State     *state.State
}

// New returns a new client with default options
func New() *Client {
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

	logger := zap.New(core)

	s := state.State{
		Users:    &basicstore.BasicStore[types.User]{},
		Servers:  &basicstore.BasicStore[types.Server]{},
		Channels: &basicstore.BasicStore[types.Channel]{},
		Members:  &basicstore.BasicStore[types.Member]{},
		Emojis:   &basicstore.BasicStore[types.Emoji]{},
	}

	rest := restcli.DefaultRestClient(&s)

	ws := gateway.DefaultGatewayConfig(rest, &s)

	c := Client{
		Rest:      rest,
		Websocket: ws,
		State:     &s,
	}

	c.Rest.Config.Logger = logger.Named("rest")
	c.Rest.Config.Ratelimiter.Logger = logger.Named("ratelimiter")
	c.Websocket.Logger = logger.Named("websocket")

	return &c
}

// Authorizes to both rest and websocket (websocket not implemented yet)
func (c *Client) Authorize(token *auth.Token) {
	// Rest client
	c.Rest.Config.SessionToken = token
	c.Websocket.SessionToken = token
}
