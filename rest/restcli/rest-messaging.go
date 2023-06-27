package restcli

import (
	"errors"
	"strconv"
	"strings"

	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Lets the server and all other clients know that we've seen this message id in this channel.
//
// <here, message is the id of the message to acknowledge and
// channel is the id of the channel the message is in>
func (c *RestClient) AcknowledgeMessage(target, message string) error {
	return rest.Request[types.APIError]{Method: rest.PUT, Path: "channels/" + target + "/ack/" + message}.NoContent(&c.Config)
}

// Fetch multiple messages.
func (c *RestClient) FetchMessages(target string, q *types.MessageQuery) (*types.MessageFetchResponse, error) {
	if q == nil {
		q = &types.MessageQuery{}
	}

	params := []string{}

	var iresp = &types.MessageFetchResponse{}

	if q.IncludeUsers {
		iresp.IncludeUsers = true
		params = append(params, "include_users=true")
	}

	runIf(q.Limit > 0, func() {
		params = append(params, "limit="+strconv.FormatInt(int64(q.Limit), 10))
	})
	runIf(q.Before != "", func() {
		params = append(params, "before="+q.Before)
	})
	runIf(q.After != "", func() {
		params = append(params, "after="+q.After)
	})
	runIf(q.Sort != "", func() {
		params = append(params, "sort="+string(q.Sort))
	})
	runIf(q.Nearby != "", func() {
		params = append(params, "nearby="+q.Nearby)
	})

	return rest.Request[types.MessageFetchResponse]{Path: "channels/" + target + "/messages?" + strings.Join(params, "&"), InitialResp: iresp}.With(&c.Config)
}

// Sends a message to the given channel.
func (c *RestClient) SendMessage(target string, d *types.DataMessageSend) (*types.Message, error) {
	return rest.Request[types.Message]{Method: rest.POST, Path: "channels/" + target + "/messages", Json: d}.With(&c.Config)
}

// This route searches for messages within the given parameters.
//
// <in actual tests, this endpoint is very slow and takes a long time to respond>
//
// <official docs provide this as one endpoint, but they return two different types>
//
// include_users=true
func (c *RestClient) SearchForMessages(target string, q *types.MessageSearchQuery) (*types.MessageFetchResponse, error) {
	if q == nil {
		return nil, errors.New("query cannot be nil")
	}

	var iresp = &types.MessageFetchResponse{IncludeUsers: q.IncludeUsers}

	return rest.Request[types.MessageFetchResponse]{Method: rest.POST, Path: "channels/" + target + "/search", Json: q, InitialResp: iresp}.With(&c.Config)
}

// This route returns any changed message objects and tells you if any have been deleted.
//
// Don't actually poll this route, instead use this to update your local database.
//
// # DEPRECATED
//
// <this will always return an error as it is deprecated *and* undocumented>
func (c *RestClient) PollMessageChanges(_target string, q *types.MessageIds) error {
	return errors.New("this route is deprecated, please use the websocket instead")
}

// Retrieves a message by its id.
func (c *RestClient) FetchMessage(target, msg string) (*types.Message, error) {
	return rest.Request[types.Message]{Path: "channels/" + target + "/messages/" + msg}.With(&c.Config)
}

// Delete a message you've sent or one you have permission to delete.
func (c *RestClient) DeleteMessage(target, msg string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "channels/" + target + "/messages/" + msg}.NoContent(&c.Config)
}

// Edits a message that you've previously sent.
func (c *RestClient) EditMessage(target, msg string, d *types.DataMessageEdit) (*types.Message, error) {
	return rest.Request[types.Message]{Method: rest.PATCH, Path: "channels/" + target + "/messages/" + msg, Json: d}.With(&c.Config)
}

// Delete multiple messages you've sent or one you have permission to delete.
//
// This will always require ManageMessages permission regardless of whether you own the message or not.
//
// Messages must have been sent within the past 1 week.
//
// <here, target is the channel id and d is an array of message ids to delete>
func (c *RestClient) BulkDeleteMessages(target string, d *types.MessageIds) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "channels/" + target + "/messages/bulk", Json: d}.NoContent(&c.Config)
}
