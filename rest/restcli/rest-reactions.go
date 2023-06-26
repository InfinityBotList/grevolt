// +interactions <in docs but its actually reactions>
package restcli

import (
	"strings"

	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// React to a given message.
//
// <msg here is the message id, target is the channel id and emoji is the emoji id to react with>
func (c *RestClient) AddReactionToMessage(target string, msg string, emoji string) error {
	return rest.Request[types.APIError]{Method: rest.PUT, Path: "channels/" + target + "/messages/" + msg + "/reactions/" + emoji}.NoContent(&c.Config)
}

// Remove your own, someone else's or all of a given reaction.
//
// Requires ManageMessages if changing others' reactions.
//
// <msg here is the message id, target is the channel id and emoji is the emoji id to react with>
func (c *RestClient) RemoveReactionsToMessage(target string, msg string, emoji string, d *types.DataReactionsRemove) error {
	if d == nil {
		d = &types.DataReactionsRemove{}
	}

	var params []string

	if d.RemoveAll {
		params = append(params, "remove_all=true")
	}

	if d.UserId != "" {
		params = append(params, "user_id="+d.UserId)
	}

	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "channels/" + target + "/messages/" + msg + "/reactions/" + emoji + "?" + strings.Join(params, "&")}.NoContent(&c.Config)
}

// Remove your own, someone else's or all of a given reaction.
//
// Requires ManageMessages permission.
//
// <msg here is the message id, target is the channel id>
func (c *RestClient) RemoveAllReactionsFromMessage(target string, msg string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "channels/" + target + "/messages/" + msg + "/reactions"}.NoContent(&c.Config)
}
