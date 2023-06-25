package tests

import "testing"

func TestCreateInvite(t *testing.T) {
	// Create invite
	cli := ITestStartup(t)

	invite, apiErr, err := cli.Rest.CreateInvite(TestChannel)

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(invite.Id)
}
