package tests

import "testing"

func TestCreateInvite(t *testing.T) {
	// Create invite
	cli := ITestStartup(t)

	invite, err := cli.Rest.CreateInvite(TestChannel)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(invite.Id)
}
