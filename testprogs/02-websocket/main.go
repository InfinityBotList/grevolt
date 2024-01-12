package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/extras/advancedevents"
	"github.com/infinitybotlist/grevolt/extras/gatewaymiddleware"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/gateway/events"
	"github.com/infinitybotlist/grevolt/types"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var config struct {
	SessionTokenUser string `yaml:"session_token_user"`
}

func main() {
	// Open config.yaml
	f, err := os.Open("../config.yaml")

	if err != nil {
		panic(err)
	}

	// Decode config.yaml
	err = yaml.NewDecoder(f).Decode(&config)

	if err != nil {
		panic(err)
	}

	// Create a new client
	c := client.New()

	fmt.Println(config.SessionTokenUser)

	// Authorize the client
	c.Authorize(&auth.Token{
		Bot:   false,
		Token: config.SessionTokenUser,
	})

	// Look Busy
	self, err := c.Rest.FetchSelf()

	if err != nil {
		panic(err)
	}

	c.Rest.EditUser(self.Id, &types.DataEditUser{
		Status: &types.UserStatus{
			Text:     "Meowing cats!",
			Presence: types.IDLE_Presence,
		},
	})

	// Use messagepack for encoding (better performance)
	c.Websocket.Encoding = "msgpack"

	// Add event debug middleware
	gatewaymiddleware.EventDebug(c.Websocket)

	c.Websocket.EventHandlers.Auth = func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.Auth) {
		fmt.Println("Auth:", e.Event.Type, ctx.Raw)
	}

	// Register ready event
	// Here, we use advancedevents to provide a more advanced error handling system
	c.Websocket.EventHandlers.Ready = advancedevents.NewEventHandler[events.Ready]().AddRaw(
		advancedevents.EventFunc[events.Ready]{
			ID: "readyEvent",
			Handler: func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.Ready) error {
				fmt.Println("Ready:", e.Event.Type, e.Members[0].JoinedAt)

				if e.Event.Type == "" {
					return errors.New("ready event type is empty")
				}

				fmt.Println("Testing bulk commands now")

				// This is how you send custom events so other parts of your code can handle them
				//
				// Custom events must begin with @
				var bulkCmd = &events.Bulk{
					Event: events.Event{
						Type: "Bulk",
					},
					V: []map[string]any{
						{
							"type": "@MyTestBulkEvent",
							"data": 10293,
							"whoa": "!",
						},
						{
							"type": "@MyTestBulkEvent",
							"data": 10294,
							"whoa": "!",
						},
					},
				}

				bulkCmdParsed, err := c.Websocket.Encode(bulkCmd)

				if err != nil {
					return err
				}

				c.Websocket.HandleEvent(bulkCmdParsed, "Bulk")

				return nil
			},
			ErrorHandlers: []advancedevents.ErrorHandler[events.Ready]{
				func(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *events.Ready, err error, handler advancedevents.EventFunc[events.Ready]) {
					w.Logger.Error(
						"Error in ready handler",
						zap.Error(err),
						zap.String("handlerId", handler.ID),
					)
				},
			},
		},
	).Build()

	c.Websocket.EventHandlers.Message = advancedevents.NewEventHandler[events.Message]().Add(
		func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.Message) error {
			fmt.Println("Message:", e.Content, e.Author)
			fmt.Println(c.State.Members.Length())
			return nil
		},
	).AddRaw(
		advancedevents.EventFunc[events.Message]{
			ID: "test",
			Handler: func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.Message) error {
				return errors.New("test error")
			},
			ErrorHandlers: []advancedevents.ErrorHandler[events.Message]{
				func(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *events.Message, err error, handler advancedevents.EventFunc[events.Message]) {
					w.Logger.Error(
						"Error in message handler",
						zap.Error(err),
						zap.String("handlerId", handler.ID),
					)
				},
			},
		},
	).Build()

	// We do a bit of wrapping to demonstrate advancedevents usage
	c.Websocket.EventHandlers.Message = advancedevents.Wrap(
		c.Websocket.EventHandlers.Message,
		advancedevents.NewMulti[events.Message](
			func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.Message) {
				fmt.Println("Multi 1")
			},
			func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.Message) {
				fmt.Println("Multi 2")
			},
		).Build(),
	)

	// Single event style
	c.Websocket.EventHandlers.MessageUpdate = func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.MessageUpdate) {
		fmt.Println("MessageUpdate:", e, e.Data, e.Id, e.ChannelId, e.Data.Content, e.Data.Edited)
	}

	c.Websocket.EventHandlers.MessageDelete = func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.MessageDelete) {
		fmt.Println("MessageDelete:", e, e.Id, e.ChannelId)
	}

	c.Websocket.EventHandlers.MessageReact = func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.MessageReact) {
		fmt.Println("MessageReact:", e, e.ChannelId, e.Id, e.UserId, e.EmojiId)
	}

	c.Websocket.EventHandlers.MessageUnreact = func(w *gateway.GatewayClient, ctx *gateway.EventContext, e *events.MessageUnreact) {
		fmt.Println("MessageUnreact:", e, e.ChannelId, e.Id, e.UserId, e.EmojiId)
	}

	err = c.Websocket.Open()

	if err != nil {
		panic(err)
	}

	go func() {
		// Wait for 10 seconds
		time.Sleep(20 * time.Second)

		// Close the client
		fmt.Println("Restarting", c.Websocket.StatusChannel.ListenersCount())

		// Send restart payload
		c.Websocket.NotifyChannel.Broadcast(&gateway.NotifyPayload{
			OpCode: gateway.RESTART_IOpCode,
		})
	}()

	c.Websocket.Wait()
}
