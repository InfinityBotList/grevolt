package tests

import (
	"testing"

	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/types/flags/permissions"
)

func TestChannelSetRolePermission(t *testing.T) {
	// Create invite
	cli := ITestStartup(t)

	c, apiErr, err := cli.Rest.ChannelSetRolePermission(TestChannel, TestRole, &types.PatchOverrideField{
		Permissions: &types.Override{
			Allow: permissions.NewPermissionManager(
				permissions.ManageChannel,
			).Build(),
			Deny: permissions.NewPermissionManager(
				permissions.ManagePermissions,
			).Build(),
		},
	})

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.RolePermissions)
}

func TestChannelSetDefaultPermission(t *testing.T) {
	// Create invite
	cli := ITestStartup(t)

	c, apiErr, err := cli.Rest.ChannelSetDefaultPermission(TestChannel, &types.PatchOverrideField{
		Permissions: &types.Override{
			Allow: permissions.NewPermissionManager(
				permissions.ManageChannel,
			).Build(),
			Deny: permissions.NewPermissionManager(
				permissions.ManagePermissions,
			).Build(),
		},
	})

	if apiErr != nil {
		t.Error(apiErr)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.RolePermissions)
}
