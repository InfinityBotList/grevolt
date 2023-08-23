package tests

import (
	"os"
	"testing"

	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/types/flags/permissions"
)

func TestServerPermissions(t *testing.T) {
	cli := ITestStartup(t)

	s, err := cli.Rest.CreateServer(&types.DataCreateServer{
		Name:        "Test Server Permissions",
		Description: "Testing 1234",
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

	os.Setenv("TEST_SERVER_SPTESTS", s.Server.Id)

	t.Run("CreateRole", testCreateRole)
	t.Run("EditRole", testEditRole)
	t.Run("SetRolePermission", testSetRolePermission)
	t.Run("SetDefaultPermission", testSetDefaultPermission)
	t.Run("DeleteRole", testDeleteRole)

	// Delete server
	err = cli.Rest.DeleteOrLeaveServer(s.Server.Id, true)

	if err != nil {
		t.Error(err)
		return
	}
}

func testCreateRole(t *testing.T) {
	cli := ITestStartup(t)

	s, err := cli.Rest.CreateRole(os.Getenv("TEST_SERVER_SPTESTS"), &types.DataCreateRole{
		Name: "Test Role",
		Rank: 1,
	})

	if err != nil {
		t.Error(err)
		return
	}

	if s == nil {
		t.Error("s is nil but should not be", s)
		return
	}

	os.Setenv("TEST_ROLE_SPTESTS", s.Id)

	t.Log("s:", s)
}

func testEditRole(t *testing.T) {
	cli := ITestStartup(t)

	s, err := cli.Rest.EditRole(os.Getenv("TEST_SERVER_SPTESTS"), os.Getenv("TEST_ROLE_SPTESTS"), &types.DataEditRole{
		Name: "Test Role",
		Rank: 1,
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

func testSetRolePermission(t *testing.T) {
	cli := ITestStartup(t)

	s, err := cli.Rest.SetRolePermission(os.Getenv("TEST_SERVER_SPTESTS"), os.Getenv("TEST_ROLE_SPTESTS"), &types.PermissionsPatchOverrideField{
		Permissions: &types.PermissionOverride{
			Allow: permissions.NewPermissionManager(
				permissions.BanMembers,
			).Build(),
			Deny: permissions.NewPermissionManager(
				permissions.SendMessage,
			).Build(),
		},
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

func testSetDefaultPermission(t *testing.T) {
	cli := ITestStartup(t)

	s, err := cli.Rest.SetDefaultPermission(os.Getenv("TEST_SERVER_SPTESTS"), &types.PermissionUpdate{
		Permissions: permissions.NewPermissionManager(
			permissions.ManageChannel,
		).Build(),
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

func testDeleteRole(t *testing.T) {
	cli := ITestStartup(t)

	err := cli.Rest.DeleteRole(os.Getenv("TEST_SERVER_SPTESTS"), os.Getenv("TEST_ROLE_SPTESTS"))

	if err != nil {
		t.Error(err)
		return
	}
}
