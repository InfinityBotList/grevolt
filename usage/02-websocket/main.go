package main

import (
	"fmt"
	"os"
	"time"

	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/client/geneva"
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
	c.Authorize(&geneva.Token{
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

	// Register ready event
	c.Websocket.EventHandlers.Ready = func(w *gateway.GatewayClient, e *events.Ready) {
		fmt.Println("Ready:", e.Users[0], e.Event.Type)

		if e.Event.Type == "" {
			panic("Ready event type is empty")
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
			panic(err)
		}

		c.Websocket.HandleEvent(bulkCmdParsed, "Bulk")
	}

	c.Websocket.EventHandlers.RawSinkFunc = func(w *gateway.GatewayClient, data []byte, typ string) {
		fmt.Println(string(data))
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
		fmt.Println("Closing", i)
		c.Websocket.Close()
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
