package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Sets permissions for the specified role in this channel.
//
// Channel must be a TextChannel or VoiceChannel.
//
// <api uses SetRolePermission but this also exists on server information>
func (c *RestClient) ChannelSetRolePermission(target, roleId string, override *types.PermissionsPatchOverrideField) (*types.Channel, error) {
	return rest.Request[types.Channel]{Method: rest.PUT, Path: "channels/" + target + "/permissions/" + roleId, Json: override}.With(&c.Config)
}

// Sets permissions for the default role in this channel.
//
// Channel must be a Group, TextChannel or VoiceChannel.
//
// <api uses SetDefaultPermission but this also exists on server information>
func (c *RestClient) ChannelSetDefaultPermission(target string, override *types.PermissionsPatchOverrideField) (*types.Channel, error) {
	return rest.Request[types.Channel]{Method: rest.PUT, Path: "channels/" + target + "/permissions/default", Json: override}.With(&c.Config)
}
