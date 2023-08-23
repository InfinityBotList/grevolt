// +untested, add tests to tests folder

package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Creates a new server role.
//
// <target is the server id>
func (c *RestClient) CreateRole(target string, d *types.DataCreateRole) (*types.NewRoleResponse, error) {
	return rest.Request[types.NewRoleResponse]{Method: rest.POST, Path: "servers/" + target + "/roles", Json: d}.With(&c.Config)
}

// Delete a server role by its id.
//
// <target is the server id>
// <roleId is the role id>
func (c *RestClient) DeleteRole(target, roleId string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "servers/" + target + "/roles/" + roleId}.NoContent(&c.Config)
}

// Edit a role by its id.
//
// <target is the server id>
// <roleId is the role id>
func (c *RestClient) EditRole(target, roleId string, d *types.DataEditRole) (*types.Role, error) {
	return rest.Request[types.Role]{Method: rest.PATCH, Path: "servers/" + target + "/roles/" + roleId, Json: d}.With(&c.Config)
}

// Sets permissions for the specified role in the server.
//
// <target is the server id>
// <roleId is the role id>
func (c *RestClient) SetRolePermission(target, roleId string, d *types.PermissionsPatchOverrideField) (*types.Server, error) {
	return rest.Request[types.Server]{Method: rest.PUT, Path: "servers/" + target + "/permissions/" + roleId, Json: d}.With(&c.Config)
}

// Sets permissions for the default role in this server.
//
// <target is the server id>
func (c *RestClient) SetDefaultPermission(target string, d *types.PermissionUpdate) (*types.Server, error) {
	return rest.Request[types.Server]{Method: rest.PUT, Path: "servers/" + target + "/permissions/default", Json: d}.With(&c.Config)
}
