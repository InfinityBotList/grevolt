package tests

import (
	"testing"
)

func TestQueryNode(t *testing.T) {
	cli := ITestStartup(t)

	qn, apiErr, err := cli.Rest.QueryNode()

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

	if qn.Ws == "" {
		t.Error("qn.Ws is empty but should not be")
	}

	t.Log("qn:", qn)
}
