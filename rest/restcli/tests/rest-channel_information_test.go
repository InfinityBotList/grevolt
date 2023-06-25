package tests

import (
	"testing"

	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

func TestChannelInformation(t *testing.T) {
	t.Run("FetchChannel", testFetchChannel)
	t.Run("EditChannel", testEditChannel)
	t.Run("CloseChannel", testCloseChannel)
}

func testFetchChannel(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	c, err := cli.Rest.FetchChannel(TestChannel)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.ChannelType, c.Name, c.Id)
}

func testEditChannel(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	c, err := cli.Rest.EditChannel(TestChannel, &types.DataEditChannel{
		Description: "This is a test channel",
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.ChannelType, c.Name, c.Id)
}

func testCloseChannel(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	c, err := cli.Rest.OpenDirectMessage(DMableUser)

	if err != nil {
		t.Error(err)
		return
	}

	if c == nil {
		t.Error("c is nil but should not be")
		return
	}

	err = cli.Rest.CloseChannel(c.Id, true)

	if err != nil {
		if apiErr, ok := err.(*rest.RestError); ok && apiErr.Type() == "NoEffect" {
			// This is good actually, it means the api call went through properly
			t.Log("NoEffect, this is good")
			return
		}

		t.Error(err)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}
}
