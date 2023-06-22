package main

import (
	"fmt"
	"os"
	"time"

	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/client/geneva"
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

	fmt.Println("Done")

	for i := 0; i < 2; i++ {
		test1(c, i)
	}

	//test2(c, 1)

	/**/
}

func test1(c *client.Client, i int) {
	err := c.Open()

	if err != nil {
		panic(err)
	}

	go func() {
		// Wait for 10 seconds
		time.Sleep(10 * time.Second)

		// Close the client
		fmt.Println("Closing", i)
		c.Websocket.Close()
	}()

	c.Websocket.Wait()
}

func test2(c *client.Client, i int) {
	err := c.Open()

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
