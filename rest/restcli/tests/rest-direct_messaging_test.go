package tests

import "testing"

func TestFetchDirectMessageChannels(t *testing.T) {
	cli := ITestStartup(t)

	d, err := cli.Rest.FetchDirectMessageChannels()

	if err != nil {
		t.Error(err)
		return
	}

	if d == nil {
		t.Error("d is nil but should not be")
		return
	}

	t.Log("d:", d)
}

func TestOpenDirectMessage(t *testing.T) {
	cli := ITestStartup(t)

	// Fetch self's ID
	self, err := cli.Rest.FetchSelf()

	if err != nil {
		t.Error(err)
		return
	}

	c, err := cli.Rest.OpenDirectMessage(self.Id)

	if err != nil {
		t.Error(err)
		return
	}

	if c == nil {
		t.Error("c is nil but should not be")
		return
	}

	t.Log("d:", c)
	t.Log("channel id:", c.Id)
	t.Log("channel name:", c.Name)
	t.Log("channel type:", c.ChannelType)
}
