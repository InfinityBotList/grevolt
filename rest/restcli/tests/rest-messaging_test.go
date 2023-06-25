package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

func TestMessages(t *testing.T) {
	t.Run("AcknowledgeMessage", testAcknowledgeMessage)
	t.Run("FetchMessagesNoIncludeUsers", testFetchMessagesNoIncludeUsers)
	t.Run("FetchMessagesIncludeUsers", testFetchMessagesIncludeUsers)
	t.Run("SendMessage", testSendMessage)
	t.Run("SearchForMessagesNoIncludeUsers", testSearchForMessagesNoIncludeUsers)
	t.Run("SearchForMessagesIncludeUsers", testSearchForMessagesIncludeUsers)
	t.Run("PollMessageChanges", testPollMessageChanges)
	t.Run("FetchMessage", testFetchMessage)
	t.Run("EditMessage", testEditMessage)
	t.Run("DeleteMessage", testDeleteMessage)
	t.Run("BulkDeleteMessages", testBulkDeleteMessages)
}

func testAcknowledgeMessage(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	err := cli.Rest.AcknowledgeMessage(TestChannel, TestMessage)

	if err != nil {
		t.Error(err)
		return
	}
}

func testFetchMessagesNoIncludeUsers(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	list, err := cli.Rest.FetchMessagesNoIncludeUsers(TestChannel, &types.MessageQuery{
		Limit:  10,
		Before: TestMessage,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if list == nil {
		t.Error("list is nil but should not be")
		return
	}

	if len(*list) == 0 {
		t.Error("list is empty but should not be")
		return
	}

	t.Log("Fetched", len(*list), "messages")
}

func testFetchMessagesIncludeUsers(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	list, err := cli.Rest.FetchMessagesIncludeUsers(TestChannel, &types.MessageQuery{
		Limit:  10,
		Before: TestMessage,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if list == nil {
		t.Error("list is nil but should not be")
		return
	}

	if len(list.Messages) == 0 {
		t.Error("list is empty but should not be")
		return
	}

	t.Log("Fetched", len(list.Messages), "messages")
}

func testSendMessage(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	msg, err := cli.Rest.SendMessage(EditChannel, &types.DataMessageSend{
		Content: "Hello, world!",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if msg == nil {
		t.Error("msg is nil but should not be")
		return
	}

	t.Log("Msg Nonce:", msg.Nonce)

	os.Setenv("SEND_MESSAGE", msg.Id)

}

func testSearchForMessagesNoIncludeUsers(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	msgs, err := cli.Rest.SearchForMessagesNoIncludeUsers(TestChannel, &types.MessageSearchQuery{
		Limit: 10,
		Query: "Hello, world!",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if msgs == nil {
		t.Error("msgs is nil but should not be")
		return
	}

	t.Log("Msg count:", len(*msgs))
}

func testSearchForMessagesIncludeUsers(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	msgs, err := cli.Rest.SearchForMessagesIncludeUsers(TestChannel, &types.MessageSearchQuery{
		Limit: 10,
		Query: "Hello, world!",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if msgs == nil {
		t.Error("msgs is nil but should not be")
		return
	}

	t.Log("Msg count:", len(msgs.Messages))
}

func testPollMessageChanges(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	err := cli.Rest.PollMessageChanges(TestChannel, &types.MessageIds{})

	if err == nil {
		t.Error("PollMessageChanges must always return an error")
	}
}

func testFetchMessage(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	msg, err := cli.Rest.FetchMessage(TestChannel, TestMessage)

	if err != nil {
		t.Error(err)
		return
	}

	if msg == nil {
		t.Error("msg is nil but should not be")
		return
	}

	t.Log("Msg Nonce:", msg.Nonce, " |  Content:", msg.Content)
}

func testEditMessage(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	msg, err := cli.Rest.EditMessage(EditChannel, os.Getenv("SEND_MESSAGE"), &types.DataMessageEdit{
		Content: "edit me (edited?)",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if msg == nil {
		t.Error("msg is nil but should not be")
		return
	}

	t.Log("Msg Nonce:", msg.Nonce, " |  Content:", msg.Content)
}

func testDeleteMessage(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	err := cli.Rest.DeleteMessage(EditChannel, os.Getenv("SEND_MESSAGE"))

	if err != nil {
		t.Error(err)
		return
	}
}

func testBulkDeleteMessages(t *testing.T) {
	// Fetch channel
	cli := ITestStartup(t)

	msg, err := cli.Rest.SendMessage(EditChannel, &types.DataMessageSend{
		Content: "Hello, world!",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if msg == nil {
		t.Error("msg is nil but should not be")
		return
	}

	t.Log("Msg Nonce:", msg.Nonce)

	// Now bulk delete from edit channel
	err = cli.Rest.BulkDeleteMessages(EditChannel, &types.MessageIds{
		Ids: []string{
			msg.Id,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}
}
