package main

import (
	"fmt"
	"os"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/client"
	"gopkg.in/yaml.v3"
)

var config struct {
	SessionTokenBot string `yaml:"session_token_bot"`
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

	// Authorize the client
	c.Authorize(&auth.Token{
		Bot:   true,
		Token: config.SessionTokenBot,
	})

	// Print the test no and name in bold

	for {
		u, err := c.Rest.FetchSelf()
		fmt.Println("User:", u, "\nError:", err, "UserBot info:", func() any {
			if u != nil {
				return u.Bot
			}
			return "<nil>"
		}())

		u, err = c.Rest.FetchUser("01FEZ09YRQ02C5XVBW6DG4QFQC")

		fmt.Println("User:", u, "\nError:", err, "UserBot info:", func() any {
			if u != nil {
				return u.Bot
			}
			return "<nil>"
		}())
	}
}
