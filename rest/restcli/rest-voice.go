package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Join Call
//
// <note that grevolt does not support vortex voice at this time, so this endpoint is pretty useless>
//
// <target is the channel id>
func (c *RestClient) JoinCall(target string) (*types.CreateVoiceUserResponse, error) {
	return rest.Request[types.CreateVoiceUserResponse]{Method: rest.POST, Path: "channels/" + target + "/join_call"}.With(&c.Config)
}
