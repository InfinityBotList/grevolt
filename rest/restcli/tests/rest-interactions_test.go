package tests

import (
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

const (
	ReactionChannel = "01H3W6M9Y83SVNPPSA9M4XFN9R"
	ReactionMessage = "01H3W7SBZ03585AWYJPA2ZZ2DJ"
)

func TestAddReactionToMessage(t *testing.T) {
	cli := ITestStartup(t)

	err := cli.Rest.AddReactionToMessage(ReactionChannel, ReactionMessage, "%F0%9F%98%80")

	if err != nil {
		t.Error(err)
	}

	t.Log("successfully added reaction")
}

func TestRemoveReactionsToMessage(t *testing.T) {
	cli := ITestStartup(t)

	err := cli.Rest.RemoveReactionsToMessage(ReactionChannel, ReactionMessage, "%F0%9F%98%80", &types.DataReactionsRemove{
		RemoveAll: true,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log("successfully added reaction")
}

func TestRemoveAllReactionsFromMessage(t *testing.T) {
	cli := ITestStartup(t)

	err := cli.Rest.RemoveAllReactionsFromMessage(ReactionChannel, ReactionMessage)

	if err != nil {
		t.Error(err)
	}

	t.Log("successfully added reaction")
}
