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

	// Print the test no and name in bold
	err = c.Open()

	if err != nil {
		panic(err)
	}

	c.Websocket.Timeout = 2 * time.Second

	go func() {
		time.Sleep(2 * time.Second)
		c.Websocket.Close()
	}()

	c.Websocket.Wait()
}
