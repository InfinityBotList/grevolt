package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

func TestBots(t *testing.T) {
	cli := ITestStartup(t)

	if cli.Rest.Config.SessionToken.Bot {
		t.Skip("Thes tests (mostly) only work for user accounts")
		return
	}

	t.Run("CreateBot", testCreateBot)
	t.Run("FetchPublicBot", testFetchPublicBot)
	t.Run("InviteBot", testInviteBot)
	t.Run("FetchBot", testFetchBot)
	t.Run("FetchOwnedBots", testFetchOwnedBots)
	t.Run("EditBot", testEditBot)
	t.Run("DeleteBot", testDeleteBot) // Needs to be last because it deletes the bot
}

func testCreateBot(t *testing.T) {
	cli := ITestStartup(t)

	qn, apiErr, err := cli.Rest.CreateBot(&types.DataCreateBot{
		Name: "PogListingBot1",
	})

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if qn == nil {
		t.Error("qn is nil but should not be")
		return
	}

	t.Log("Created bot:", qn.Id, "with token:", qn.Token)

	os.Setenv("BOT_TOKEN", qn.Token)
	os.Setenv("BOT_ID", qn.Id)
}

func testFetchPublicBot(t *testing.T) {
	if os.Getenv("BOT_ID") == "" {
		panic("BOT_ID is not set by testCreateBot")
	}

	cli := ITestStartup(t)

	qn, apiErr, err := cli.Rest.FetchPublicBot(os.Getenv("BOT_ID"))

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if qn == nil {
		t.Error("qn is nil but should not be")
		return
	}

	t.Log("Created bot:", qn.Id, "with name:", qn.Username, "and desc", qn.Description)
}

func testInviteBot(t *testing.T) {
	if os.Getenv("BOT_ID") == "" {
		panic("BOT_ID is not set by testCreateBot")
	}

	cli := ITestStartup(t)

	apiErr, err := cli.Rest.InviteBot(os.Getenv("BOT_ID"), &types.DataInviteBot{
		Server: os.Getenv("BOT_INVITE_SERVER_ID"),
	})

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Invited bot to server successfully!")
}

func testFetchBot(t *testing.T) {
	if os.Getenv("BOT_ID") == "" {
		panic("BOT_ID is not set by testCreateBot")
	}

	cli := ITestStartup(t)

	qn, apiErr, err := cli.Rest.FetchBot(os.Getenv("BOT_ID"))

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if qn == nil {
		t.Error("qn is nil but should not be")
		return
	}

	t.Log("Fetched bot:", qn)
}

func testFetchOwnedBots(t *testing.T) {
	cli := ITestStartup(t)

	qn, apiErr, err := cli.Rest.FetchOwnedBots()

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if qn == nil {
		t.Error("qn is nil but should not be")
		return
	}

	if len(qn.Bots) == 0 {
		t.Error("qn.Bots is empty but should not be")
		return
	}

	if len(qn.Users) == 0 {
		t.Error("qn.Users is empty but should not be")
		return
	}

	t.Log("Fetched owned bots response:", qn)
}

func testDeleteBot(t *testing.T) {
	if os.Getenv("BOT_ID") == "" {
		panic("BOT_ID is not set by testCreateBot")
	}

	cli := ITestStartup(t)

	apiErr, err := cli.Rest.DeleteBot(os.Getenv("BOT_ID"))

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Deleted bot successfully!")
}

func testEditBot(t *testing.T) {
	if os.Getenv("BOT_ID") == "" {
		panic("BOT_ID is not set by testCreateBot")
	}

	cli := ITestStartup(t)

	qn, apiErr, err := cli.Rest.EditBot(os.Getenv("BOT_ID"), &types.DataEditBot{
		Name: "Miguel",
	})

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if qn == nil {
		t.Error("qn is nil but should not be")
		return
	}

	t.Log("Editted bot:", qn)
}
