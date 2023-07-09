package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
)

func TestServerInformation(t *testing.T) {
	t.Run("CreateServer", testCreateServer)
	t.Run("FetchServer", testFetchServer)
	t.Run("EditServer", testEditServer)
	t.Run("CreateChannel", testCreateChannel)
	t.Run("MarkServerAsRead", testMarkServerAsRead)
	t.Run("DeleteOrLeaveServer", testDeleteOrLeaveServer)
}

func testCreateServer(t *testing.T) {
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

	os.Setenv("TEST_SERVER", s.Server.Id)
}

func testFetchServer(t *testing.T) {
	if os.Getenv("TEST_SERVER") == "" {
		t.Skip("TEST_SERVER is not set")
		return
	}

	cli := ITestStartup(t)

	s, err := cli.Rest.FetchServer(os.Getenv("TEST_SERVER"))

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

func testEditServer(t *testing.T) {
	if os.Getenv("TEST_SERVER") == "" {
		t.Skip("TEST_SERVER is not set")
		return
	}

	cli := ITestStartup(t)

	s, err := cli.Rest.EditServer(os.Getenv("TEST_SERVER"), &types.DataEditServer{
		Name:        "Test Server Editted Name",
		Description: "Test",
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

func testCreateChannel(t *testing.T) {
	if os.Getenv("TEST_SERVER") == "" {
		t.Skip("TEST_SERVER is not set")
		return
	}

	cli := ITestStartup(t)

	c, err := cli.Rest.CreateChannel(os.Getenv("TEST_SERVER"), &types.DataCreateChannel{
		Name: "testchannel1",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if c == nil {
		t.Error("c is nil but should not be", c)
		return
	}

	t.Log("c:", c)
}

func testMarkServerAsRead(t *testing.T) {
	if os.Getenv("TEST_SERVER") == "" {
		t.Skip("TEST_SERVER is not set")
		return
	}

	cli := ITestStartup(t)

	err := cli.Rest.MarkServerAsRead(os.Getenv("TEST_SERVER"))

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("successfully marked server as read")
}

func testDeleteOrLeaveServer(t *testing.T) {
	if os.Getenv("TEST_SERVER") == "" {
		t.Skip("TEST_SERVER is not set")
		return
	}

	cli := ITestStartup(t)

	err := cli.Rest.DeleteOrLeaveServer(os.Getenv("TEST_SERVER"), false)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("successfully deleted server")
}
