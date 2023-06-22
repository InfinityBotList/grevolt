package main

import (
	"fmt"
	"os"

	"github.com/infinitybotlist/grevolt/client"
	"github.com/infinitybotlist/grevolt/rest/restconfig"
	"gopkg.in/yaml.v3"
)

var config struct {
	SessionTokenBot string `yaml:"session_token_bot"`
}

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
	c := client.New()

	// Authorize the client
	c.Authorize(&restconfig.Token{
		Bot:   true,
		Token: config.SessionTokenBot,
	})

	// Fetch user info
	u, apiErr, err := c.Rest.FetchSelf()

	if err != nil {
		fmt.Println("API Error:", apiErr, "\nError:", err)
	}

	fmt.Println("User:", u)
}
