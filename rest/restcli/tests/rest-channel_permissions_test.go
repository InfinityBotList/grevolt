package tests

import (
	"testing"

	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/types/flags/permissions"
)

func TestChannelSetRolePermission(t *testing.T) {
	// Create invite
	cli := ITestStartup(t)

	c, err := cli.Rest.ChannelSetRolePermission(TestChannel, TestRole, &types.PermissionsPatchOverrideField{
		Permissions: &types.PermissionOverride{
			Allow: permissions.NewPermissionManager(
				permissions.ManageChannel,
			).Build(),
			Deny: permissions.NewPermissionManager(
				permissions.ManagePermissions,
			).Build(),
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.RolePermissions)
}

func TestChannelSetDefaultPermission(t *testing.T) {
	// Create invite
	cli := ITestStartup(t)

	c, err := cli.Rest.ChannelSetDefaultPermission(TestChannel, &types.PermissionsPatchOverrideField{
		Permissions: &types.PermissionOverride{
			Allow: permissions.NewPermissionManager(
				permissions.ManageChannel,
			).Build(),
			Deny: permissions.NewPermissionManager(
				permissions.ManagePermissions,
			).Build(),
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(c.RolePermissions)
}
