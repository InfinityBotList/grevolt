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
func (c *RestClient) ChannelSetRolePermission(target, roleId string, override *types.PatchOverrideField) (*types.Channel, *types.APIError, error) {
	var ch *types.Channel
	apiErr, err := rest.NewReq(&c.Config).Put("/channels/" + target + "/permissions/" + roleId).Json(override).DoAndMarshal(&ch)
	return ch, apiErr, err
}

// Sets permissions for the default role in this channel.
//
// Channel must be a Group, TextChannel or VoiceChannel.
//
// <api uses SetDefaultPermission but this also exists on server information>
func (c *RestClient) ChannelSetDefaultPermission(target string, override *types.PatchOverrideField) (*types.Channel, *types.APIError, error) {
	var ch *types.Channel
	apiErr, err := rest.NewReq(&c.Config).Put("/channels/" + target + "/permissions/default").Json(override).DoAndMarshal(&ch)
	return ch, apiErr, err
}
