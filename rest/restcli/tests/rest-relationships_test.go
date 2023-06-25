package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

func TestFetchMutualFriendsAndServers(t *testing.T) {
	cli := ITestStartup(t)

	ch, err := cli.Rest.FetchMutualFriendsAndServers(UserZomatree)

	if err != nil {
		t.Error(err)
		return
	}

	if ch == nil {
		t.Error("d is nil but should not be")
		return
	}

	t.Log("ch:", ch)
	t.Log("users:", ch.Users)
	t.Log("servers:", ch.Servers)
}

func TestMutuals(t *testing.T) {
	if os.Getenv("TEST_FRIEND_REQUEST") == "" {
		t.Skip("skipping test; TEST_FRIEND_REQUEST not set")
		return
	}
	t.Run("TestAcceptFriendRequest", testAcceptFriendRequest)
	t.Run("TestDenyFriendRequest", testDenyFriendRequest)
	t.Run("TestBlockUser", testBlockUser)
	t.Run("TestUnblockUser", testUnblockUser)
	t.Run("testSendFriendRequest", testSendFriendRequest)
}

func testAcceptFriendRequest(t *testing.T) {
	cli := ITestStartup(t)

	u, err := cli.Rest.AcceptFriendRequest(os.Getenv("TEST_FRIEND_REQUEST__ACCEPT"))

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("ch:", u)
}

func testDenyFriendRequest(t *testing.T) {
	cli := ITestStartup(t)

	u, err := cli.Rest.DenyFriendRequestOrRemoveFriend(os.Getenv("TEST_FRIEND_REQUEST__DENY"))

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("ch:", u)
}

func testBlockUser(t *testing.T) {
	cli := ITestStartup(t)

	u, err := cli.Rest.BlockUser(os.Getenv("TEST_FRIEND_REQUEST__BLOCKUSER"))

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("ch:", u)
}

func testUnblockUser(t *testing.T) {
	cli := ITestStartup(t)

	u, err := cli.Rest.BlockUser(os.Getenv("TEST_FRIEND_REQUEST__UNBLOCKUSER"))

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("ch:", u)
}

func testSendFriendRequest(t *testing.T) {
	cli := ITestStartup(t)

	u, err := cli.Rest.SendFriendRequest(&types.DataSendFriendRequest{
		Username: os.Getenv("TEST_FRIEND_REQUEST__SEND"),
	})

	if err != nil {
		t.Error(err)
		return
	}

	if u == nil {
		t.Error("u is nil but should not be")
		return
	}

	t.Log("ch:", u)
}
