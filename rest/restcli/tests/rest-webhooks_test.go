package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

const TestWebhookChannel = "01H404KKXTG7XNKV4MCERDXM6C"

func TestWebhooks(t *testing.T) {
	t.Run("CreateWebhook", testCreateWebhook)
	t.Run("GetAllWebhooks", testGetAllWebhooks)
}

func testCreateWebhook(t *testing.T) {
	cli := ITestStartup(t)

	wh, err := cli.Rest.CreateWebhook(TestWebhookChannel, &types.DataCreateWebhook{
		Name: "Test Webhook",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if wh == nil {
		t.Error("wh is nil but should not be")
		return
	}

	os.Setenv("TEST_WEBHOOK_ID", wh.Id)

	t.Log("successfully created webhook", wh)
}

func testGetAllWebhooks(t *testing.T) {
	cli := ITestStartup(t)

	wh, err := cli.Rest.GetAllWebhooks(TestWebhookChannel)

	if err != nil {
		t.Error(err)
		return
	}

	if wh == nil {
		t.Error("wh is nil but should not be")
		return
	}

	t.Log("successfully fetched webhooks", wh)
}
