package main

import (
	"fmt"
	"os"
	"time"

	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/client/geneva"
	"github.com/infinitybotlist/grevolt/types"
	"gopkg.in/yaml.v3"
)

var config struct {
	SessionTokenUser string `yaml:"session_token_bot"`
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
		Bot:   true,
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
