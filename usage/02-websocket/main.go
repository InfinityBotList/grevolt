package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/client/auth"
	"github.com/infinitybotlist/grevolt/extras/advancedevents"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/types/events"
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
	self, apiErr, err := c.Rest.FetchSelf()

	if err != nil {
		panic(err)
	}

	if apiErr != nil {
		panic(apiErr.Type())
	}

	c.Rest.EditUser(self.Id, &types.DataEditUser{
		Status: &types.UserStatus{
			Text:     "Meowing cats!",
			Presence: types.IDLE_Presence,
		},
	})

	// Prepare WS client
	err = c.PrepareWS()

	if err != nil {
		panic(err)
	}

	c.Websocket.EventHandlers.RawSinkFunc = func(w *gateway.GatewayClient, data []byte, typ string) {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(string(data))
		}
	}

	c.Websocket.EventHandlers.Auth = func(w *gateway.GatewayClient, e *events.Auth) {
		fmt.Println("Auth:", e.Event.Type)
	}

	// Register ready event
	// Here, we use advancedevents to provide a more advanced error handling system
	c.Websocket.EventHandlers.Ready = advancedevents.NewEventHandler[events.Ready]().AddRaw(
		advancedevents.EventFunc[events.Ready]{
			ID: "readyEvent",
			Handler: func(w *gateway.GatewayClient, e *events.Ready) error {
				fmt.Println("Ready:", e.Users, e.Event.Type)

				if e.Event.Type == "" {
					return errors.New("ready event type is empty")
				}

				fmt.Println("Testing bulk commands now")

				// This is how you send custom events so other parts of your code can handle them
				//
				// Its mostly for debugging purposes tho
				var bulkCmd = &events.Bulk{
					Event: events.Event{
						Type: "Bulk",
					},
					V: []map[string]any{
						{
							"type": "MyTestBulkEvent",
							"data": 10293,
							"whoa": "!",
						},
						{
							"type": "MyTestBulkEvent",
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
				func(w *gateway.GatewayClient, evt *events.Ready, err error, handler advancedevents.EventFunc[events.Ready]) {
					w.Logger.Errorln("Error in ready handler", handler.ID, ":", err)
				},
			},
		},
	).Build()

	c.Websocket.EventHandlers.Message = advancedevents.NewEventHandler[events.Message]().Add(
		func(w *gateway.GatewayClient, e *events.Message) error {
			fmt.Println("Message:", e.Content, e.Author)
			return nil
		},
	).AddRaw(
		advancedevents.EventFunc[events.Message]{
			ID: "test",
			Handler: func(w *gateway.GatewayClient, e *events.Message) error {
				return errors.New("test error")
			},
			ErrorHandlers: []advancedevents.ErrorHandler[events.Message]{
				func(w *gateway.GatewayClient, evt *events.Message, err error, handler advancedevents.EventFunc[events.Message]) {
					w.Logger.Errorln("Error in handler", handler.ID, ":", err)
				},
			},
		},
	).Build()

	// We do a bit of wrapping to demonstrate advancedevents usage
	c.Websocket.EventHandlers.Message = advancedevents.Wrap(
		c.Websocket.EventHandlers.Message,
		advancedevents.NewMulti[events.Message](
			func(w *gateway.GatewayClient, e *events.Message) {
				fmt.Println("Multi 1")
			},
			func(w *gateway.GatewayClient, e *events.Message) {
				fmt.Println("Multi 2")
			},
		).Build(),
	)

	// Single event style
	c.Websocket.EventHandlers.MessageUpdate = func(w *gateway.GatewayClient, e *events.MessageUpdate) {
		fmt.Println("MessageUpdate:", e, e.Data, e.Id, e.ChannelId, e.Data.Content, e.Data.Edited)
	}

	c.Websocket.EventHandlers.MessageDelete = func(w *gateway.GatewayClient, e *events.MessageDelete) {
		fmt.Println("MessageDelete:", e, e.Id, e.ChannelId)
	}

	c.Websocket.EventHandlers.MessageReact = func(w *gateway.GatewayClient, e *events.MessageReact) {
		fmt.Println("MessageReact:", e, e.ChannelId, e.Id, e.UserId, e.EmojiId)
	}

	c.Websocket.EventHandlers.MessageUnreact = func(w *gateway.GatewayClient, e *events.MessageUnreact) {
		fmt.Println("MessageUnreact:", e, e.ChannelId, e.Id, e.UserId, e.EmojiId)
	}

	for i := 0; i < 2; i++ {
		test1(c, i)
	}

	//test2(c, 1)

	/**/
}

func test1(c *client.Client, i int) {
	err := c.Websocket.Open()

	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(10 * time.Second)

		c.Websocket.BeginTyping("01GDT82E0JPN8K40TDGM33QPXS")

		time.Sleep(2 * time.Second)

		c.Websocket.EndTyping("01GDT82E0JPN8K40TDGM33QPXS")

		// Wait for 30 seconds
		time.Sleep(30 * time.Second)

		// Close the client
		//fmt.Println("Closing", i)
		//c.Websocket.Close()
	}()

	c.Websocket.Wait()
}

func test2(c *client.Client, i int) {
	err := c.Websocket.Open()

	if err != nil {
		panic(err)
	}

	/* go func() {
		// Wait for 10 seconds
		time.Sleep(10 * time.Second)

		// Close the client
		fmt.Println("Closing", i)

		// Send restart payload
		for {
			c.Websocket.NotifyChannel <- &gateway.NotifyPayload{
				OpCode: gateway.RESTART_IOpCode,
				Data:   map[string]any{},
			}

			time.Sleep(10 * time.Second)
		}
	}() */
	c.Websocket.Wait()
}
