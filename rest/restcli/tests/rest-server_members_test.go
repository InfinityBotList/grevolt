package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

func TestServerMembers(t *testing.T) {
	cli := ITestStartup(t)

	s, err := cli.Rest.CreateServer(&types.DataCreateServer{
		Name:        "Test Server",
		Description: "Test",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if s == nil || s.Server == nil {
		t.Error("s/s.Server is nil but should not be", s)
		return
	}

	t.Log("s:", s)

	qn, err := cli.Rest.CreateBot(&types.DataCreateBot{
		Name: "PogListingBot1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if qn == nil {
		t.Error("qn is nil but should not be")
		return
	}

	// Invite bot to server
	err = cli.Rest.InviteBot(qn.Id, &types.DataInviteBot{
		Server: s.Server.Id,
	})

	if err != nil {
		t.Error(err)
		return
	}

	// Create invite on server for testing purposes
	invite, err := cli.Rest.CreateInvite(s.Channels[0].Id)

	if err != nil {
		t.Error(err)
		return
	}

	os.Setenv("TEST_INVITE_SMTESTS", invite.Id)
	os.Setenv("TEST_BOT_ID_SMTESTS", qn.Id)
	os.Setenv("TEST_SERVER_SMTESTS", s.Server.Id)

	t.Run("FetchMembers", testFetchMembers)
	t.Run("FetchMember", testFetchMember)
	t.Run("EditMember", testEditMember)
	t.Run("QueryMembersByName", testQueryMembersByName)
	t.Run("FetchInvites", testFetchInvites)
	t.Run("BanUser", testBanUser)
	t.Run("UnbanUser", testUnbanUser)
	t.Run("FetchBans", testFetchBans)

	// Delete server
	err = cli.Rest.DeleteOrLeaveServer(s.Server.Id, true)

	if err != nil {
		t.Error(err)
		return
	}

	// Delete bot
	err = cli.Rest.DeleteBot(qn.Id)

	if err != nil {
		t.Error(err)
		return
	}
}

func testFetchMembers(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	s, err := cli.Rest.FetchMembers(os.Getenv("TEST_SERVER_SMTESTS"))

	if err != nil {
		t.Error(err)
		return
	}

	if s == nil || len(s.Members) == 0 {
		t.Error("s is nil but should not be", s)
		return
	}

	t.Log("s:", s)
}

func testFetchMember(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	if os.Getenv("TEST_BOT_ID_SMTESTS") == "" {
		t.Skip("TEST_BOT_ID_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	s, err := cli.Rest.FetchMember(os.Getenv("TEST_SERVER_SMTESTS"), os.Getenv("TEST_BOT_ID_SMTESTS"))

	if err != nil {
		t.Error(err)
		return
	}

	if s == nil {
		t.Error("s is nil but should not be", s)
		return
	}

	t.Log("s:", s)
}

func testEditMember(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	if os.Getenv("TEST_BOT_ID_SMTESTS") == "" {
		t.Skip("TEST_BOT_ID_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	s, err := cli.Rest.EditMember(os.Getenv("TEST_SERVER_SMTESTS"), os.Getenv("TEST_BOT_ID_SMTESTS"), &types.DataMemberEdit{
		Nickname: "EdittedPogBot",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if s == nil {
		t.Error("s is nil but should not be", s)
		return
	}

	t.Log("s:", s)
}

func testQueryMembersByName(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	if os.Getenv("TEST_BOT_ID_SMTESTS") == "" {
		t.Skip("TEST_BOT_ID_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	err := cli.Rest.QueryMembersByName(os.Getenv("TEST_SERVER_SMTESTS"))

	if err == nil {
		t.Error("PollMessageChanges must always return an error")
	}
}

func testFetchInvites(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	s, err := cli.Rest.FetchInvites(os.Getenv("TEST_SERVER_SMTESTS"))

	if err != nil {
		t.Error(err)
		return
	}

	if s == nil {
		t.Error("s is nil but should not be", s)
		return
	}

	t.Log("s:", s)
}

func testBanUser(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	if os.Getenv("TEST_BOT_ID_SMTESTS") == "" {
		t.Skip("TEST_BOT_ID_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	sb, err := cli.Rest.BanUser(os.Getenv("TEST_SERVER_SMTESTS"), os.Getenv("TEST_BOT_ID_SMTESTS"), &types.DataBanCreate{
		Reason: "Test",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if sb == nil {
		t.Error("sb is nil but should not be", sb)
		return
	}

	t.Log("sb:", sb)
}

func testUnbanUser(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	if os.Getenv("TEST_BOT_ID_SMTESTS") == "" {
		t.Skip("TEST_BOT_ID_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	err := cli.Rest.UnbanUser(os.Getenv("TEST_SERVER_SMTESTS"), os.Getenv("TEST_BOT_ID_SMTESTS"))

	if err != nil {
		t.Error(err)
		return
	}
}

func testFetchBans(t *testing.T) {
	if os.Getenv("TEST_SERVER_SMTESTS") == "" {
		t.Skip("TEST_SERVER_SMTESTS is not set")
		return
	}

	cli := ITestStartup(t)

	sb, err := cli.Rest.FetchBans(os.Getenv("TEST_SERVER_SMTESTS"))

	if err != nil {
		t.Error(err)
		return
	}

	if sb == nil {
		t.Error("sb is nil but should not be", sb)
		return
	}

	t.Log("sb:", len(sb.Bans))
}

/*
func testKickMember
func testFetchBans
*/
