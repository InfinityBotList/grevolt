package restcli

import (
	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Create a new Revolt bot.
func (c *RestClient) CreateBot(d *types.DataCreateBot) (*types.Bot, error) {
	return rest.Request[types.Bot]{Method: rest.POST, Path: "bots/create", Json: d}.With(&c.Config)
}

// Fetch details of a public (or owned) bot by its id.
func (c *RestClient) FetchPublicBot(target string) (*types.PublicBot, error) {
	return rest.Request[types.PublicBot]{Path: "bots/" + target + "/invite"}.With(&c.Config)
}

// Invite a bot to a server or group by its id.`
func (c *RestClient) InviteBot(target string, d *types.DataInviteBot) error {
	return rest.Request[types.APIError]{Method: rest.POST, Path: "bots/" + target + "/invite", Json: d}.NoContent(&c.Config)
}

// Fetch details of a bot you own by its id.
func (c *RestClient) FetchBot(target string) (*types.Bot, error) {
	return rest.Request[types.Bot]{Path: "bots/" + target}.With(&c.Config)
}

// Fetch all of the bots that you have control over.
func (c *RestClient) FetchOwnedBots() (*types.OwnedBotsResponse, error) {
	return rest.Request[types.OwnedBotsResponse]{Path: "bots/@me"}.With(&c.Config)
}

// Delete a bot by its id.
func (c *RestClient) DeleteBot(target string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "bots/" + target}.NoContent(&c.Config)
}

// Edit bot details by its id.
func (c *RestClient) EditBot(target string, d *types.DataEditBot) (*types.Bot, error) {
	return rest.Request[types.Bot]{Method: rest.PATCH, Path: "bots/" + target, Json: d}.With(&c.Config)
}
