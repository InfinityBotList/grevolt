package tests

import (
	"testing"
)

func TestFetchSelf(t *testing.T) {
	cli := ITestStartup(t)

	u, apiErr, err := cli.Rest.FetchSelf()

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("u:", u)
}

func TestFetchUser(t *testing.T) {
	cli := ITestStartup(t)

	// We use zomatree here because it's a user that is guaranteed to exist
	u, apiErr, err := cli.Rest.FetchUser(UserZomatree)

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("u:", u)
}
