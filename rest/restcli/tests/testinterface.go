package tests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/infinitybotlist/grevolt/auth"
	"github.com/infinitybotlist/grevolt/client"
	"gopkg.in/yaml.v3"
)

type testConfig struct {
	SessionTokenBot  string `yaml:"session_token_bot"`
	SessionTokenUser string `yaml:"session_token_user"`
}

const (
	UserZomatree = "01FD58YK5W7QRV5H3D64KTQYX3"
	TestChannel  = "01G11DTVYAJQCJJ9VZMA6GRND3"
	TestMessage  = "01H3SPT5VV7J5XQ5615WJXHJC2"
	DMableUser   = "01FEZ09YRQ02C5XVBW6DG4QFQC"
	TestRole     = "01H3PQP49C2J75ZC7KGCAP317C"
	EditChannel  = "01GDT82E0JPN8K40TDGM33QPXS"
	EditMessage  = "01H3SVZX37X0HFQP7DJJ35G096"
)

// Defines a set of common functions for testing
func ITestStartup(t *testing.T) *client.Client {
	// Use git rev-parse --show-toplevel to get the root directory
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	// Run the command
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	// Remove the newline from the end of the output
	rootDir := strings.TrimSuffix(string(out), "\n")

	f, err := os.Open(rootDir + "/test.yaml")

	if err != nil {
		panic(err)
	}

	// Decode config.yaml
	var config testConfig
	err = yaml.NewDecoder(f).Decode(&config)

	if err != nil {
		panic(err)
	}

	// Create a new client
	c := client.New()

	// Authorize the client, rn, we use the User API as the bot API is pretty crap
	c.Authorize(&auth.Token{
		Bot:   false,
		Token: config.SessionTokenUser,
	})

	c.Rest.Config.APIUrl = "https://app.revolt.chat/api/"

	return c
}
