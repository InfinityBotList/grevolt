package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Creates a webhook which 3rd party platforms can use to send messages
//
// <target is the channel id>
func (c *RestClient) CreateWebhook(target string, webhook *types.DataCreateWebhook) (*types.Webhook, error) {
	return rest.Request[types.Webhook]{Method: rest.POST, Path: "channels/" + target + "/webhooks", Json: webhook}.With(&c.Config)
}

// Gets all webhooks inside the channel
//
// <target is the channel id>
func (c *RestClient) GetAllWebhooks(target string) (*types.WebhookList, error) {
	return rest.Request[types.WebhookList]{Path: "channels/" + target + "/webhooks"}.With(&c.Config)
}
