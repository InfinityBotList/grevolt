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
//
// <official docs provide this as one endpoint, but they return two different types>
//
// include_users=false
func (c *RestClient) FetchMessagesNoIncludeUsers(target string, q *types.MessageQuery) (*types.MessageList, error) {
	if q == nil {
		q = &types.MessageQuery{}
	}

	params := []string{
		"include_users=false",
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

	return rest.Request[types.MessageList]{Path: "channels/" + target + "/messages?" + strings.Join(params, "&")}.With(&c.Config)
}

// Fetch multiple messages.
//
// <official docs provide this as one endpoint, but they return two different types>
//
// include_users=true
func (c *RestClient) FetchMessagesIncludeUsers(target string, q *types.MessageQuery) (*types.MessageFetchExtendedResponse, error) {
	if q == nil {
		q = &types.MessageQuery{}
	}

	params := []string{
		"include_users=true",
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

	return rest.Request[types.MessageFetchExtendedResponse]{Path: "channels/" + target + "/messages?" + strings.Join(params, "&")}.With(&c.Config)
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
// include_users=false
func (c *RestClient) SearchForMessagesNoIncludeUsers(target string, q *types.MessageSearchQuery) (*types.MessageList, error) {
	if q == nil {
		return nil, errors.New("query cannot be nil")
	}

	q.IncludeUsers = false // Override, must be false when using NoIncludeUsers
	return rest.Request[types.MessageList]{Method: rest.POST, Path: "channels/" + target + "/search", Json: q}.With(&c.Config)
}

// This route searches for messages within the given parameters.
//
// <in actual tests, this endpoint is very slow and takes a long time to respond>
//
// <official docs provide this as one endpoint, but they return two different types>
//
// include_users=true
func (c *RestClient) SearchForMessagesIncludeUsers(target string, q *types.MessageSearchQuery) (*types.MessageFetchExtendedResponse, error) {
	if q == nil {
		return nil, errors.New("query cannot be nil")
	}

	q.IncludeUsers = true // Override, must be true when using IncludeUsers
	return rest.Request[types.MessageFetchExtendedResponse]{Method: rest.POST, Path: "channels/" + target + "/search", Json: q}.With(&c.Config)
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
