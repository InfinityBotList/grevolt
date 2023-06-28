package tests

import "testing"

const VC = "01H3XRYTRR8AN2MD0E29EG1J8N"

func TestJoinCall(t *testing.T) {
	cli := ITestStartup(t)

	tk, err := cli.Rest.JoinCall(VC)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("joined call:", tk.Token)
}
