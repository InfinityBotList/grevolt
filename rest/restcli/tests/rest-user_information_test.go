package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

func TestFetchSelf(t *testing.T) {
	cli := ITestStartup(t)

	u, err := cli.Rest.FetchSelf()

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
	u, err := cli.Rest.FetchUser(UserZomatree)

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

func TestEditUser(t *testing.T) {
	cli := ITestStartup(t)

	// First fetch self
	self, err := cli.Rest.FetchSelf()

	if err != nil {
		t.Error(err)
		return
	}

	// We use zomatree here because it's a user that is guaranteed to exist
	u, err := cli.Rest.EditUser(self.Id, &types.DataEditUser{
		DisplayName: "Rootspring",
	})

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

func TestFetchUserFlags(t *testing.T) {
	cli := ITestStartup(t)

	// We use zomatree here because it's a user that is guaranteed to exist
	u, err := cli.Rest.FetchUserFlags(UserZomatree)

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("f:", u.Flags)
}

func TestChangeUsername(t *testing.T) {
	var pass string

	if os.Getenv("PASSWORD") == "" {
		t.Skip("Skipping test because PASSWORD is not set")
	}

	cli := ITestStartup(t)

	// We use zomatree here because it's a user that is guaranteed to exist
	u, err := cli.Rest.ChangeUsername(&types.DataChangeUsername{
		Username: "root",
		Password: pass,
	})

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

func TestFetchDefaultAvatar(t *testing.T) {
	cli := ITestStartup(t)

	// We use zomatree here because it's a user that is guaranteed to exist
	bytes, err := cli.Rest.FetchDefaultAvatar(UserZomatree)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(bytes.Raw), "bytes")
}

func TestFetchUserProfile(t *testing.T) {
	cli := ITestStartup(t)

	// We use zomatree here because it's a user that is guaranteed to exist
	u, err := cli.Rest.FetchUserProfile(UserZomatree)

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
