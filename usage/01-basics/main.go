package main

import (
	"fmt"
	"os"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/types/events"
	"gopkg.in/yaml.v3"
)

var config struct {
	SessionTokenBot string `yaml:"session_token_bot"`
}

var Client client.Client

func main() {
	// Open config.yaml
	f, err := os.Open("config.yaml")

	if err != nil {
		panic(err)
	}

	// Decode config.yaml
	err = yaml.NewDecoder(f).Decode(&config)

	if err != nil {
		panic(err)
	}

	// Create a new client
	Client = client.New()

	// Authorize the client
	Client.Authorize(&auth.Token{
		Bot:   true,
		Token: config.SessionTokenBot,
	})

	Client.Websocket.EventHandlers.Ready = func(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *events.Ready) {
		fmt.Println("Websocket up")

		// Get self
		u, err := Client.Rest.FetchSelf()

		if err != nil {
			panic("Failed to fetch self:" + err.Error())
		}

		fmt.Println("Logged in as:", u.Username)
	}

	Client.Websocket.EventHandlers.Message = messageHandler

	// Connect to the websocket
	err = Client.Websocket.Open()

	if err != nil {
		panic(err)
	}

	// Wait for the websocket to close
	Client.Websocket.Wait()
}
