package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

const (
	TestUserInitial = "01FE57SEGM0CBQD6Y7X10VZQ49"
	TestUserSecond  = "01H3MB3T2A7ZHA4EV8J78VVRMH"
)

func TestGroups(t *testing.T) {
	t.Run("CreateGroup", testCreateGroup)
	t.Run("FetchGroupMembers", testFetchGroupMembers)
	t.Run("AddMemberToGroup", testAddMemberToGroup)
	t.Run("RemoveMemberFromGroup", testRemoveMemberFromGroup)
}

func testFetchGroupMembers(t *testing.T) {
	if os.Getenv("GROUP_ID") == "" {
		panic("GROUP_ID not set")
	}

	cli := ITestStartup(t)

	c, err := cli.Rest.FetchGroupMembers(os.Getenv("GROUP_ID"))

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("successfully fetched group members", c)
}

func testCreateGroup(t *testing.T) {
	cli := ITestStartup(t)

	c, err := cli.Rest.CreateGroup(&types.DataCreateGroup{
		Name:        "Test Group",
		Description: "Automated test group",
		Users: []string{
			TestUserInitial,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	os.Setenv("GROUP_ID", c.Id)

	t.Log("successfully created group", c.Name)
}

func testAddMemberToGroup(t *testing.T) {
	if os.Getenv("GROUP_ID") == "" {
		panic("GROUP_ID not set")
	}

	cli := ITestStartup(t)

	err := cli.Rest.AddMemberToGroup(os.Getenv("GROUP_ID"), TestUserSecond)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("successfully added member to group")
}

func testRemoveMemberFromGroup(t *testing.T) {
	if os.Getenv("GROUP_ID") == "" {
		panic("GROUP_ID not set")
	}

	cli := ITestStartup(t)

	// Cleanup here and remove initial user from group
	err := cli.Rest.RemoveMemberFromGroup(os.Getenv("GROUP_ID"), TestUserInitial)

	if err != nil {
		t.Error("Error cleaning up:", err)
		return
	}

	// Close group
	err = cli.Rest.CloseChannel(os.Getenv("GROUP_ID"), false)

	if err != nil {
		t.Error("Error closing group:", err)
		return
	}

	t.Log("successfully removed member from group")
}
