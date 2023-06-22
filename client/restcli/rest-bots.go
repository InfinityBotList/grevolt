package restcli

import (
	"github.com/infinitybotlist/grevolt/rest/clientapi"
	"github.com/infinitybotlist/grevolt/types"
)

// Create a new Revolt bot.
func (c *RestClient) CreateBot(d *types.DataCreateBot) (*types.Bot, *types.APIError, error) {
	var bot *types.Bot
	apiErr, err := clientapi.NewReq(&c.Config).Post("bots/create").Json(d).DoAndMarshal(&bot)
	return bot, apiErr, err
}

// Fetch details of a public (or owned) bot by its id.
func (c *RestClient) FetchPublicBot(target string) (*types.PublicBot, *types.APIError, error) {
	var bot *types.PublicBot
	apiErr, err := clientapi.NewReq(&c.Config).Get("bots/" + target + "/invite").DoAndMarshal(&bot)
	return bot, apiErr, err
}

// Invite a bot to a server or group by its id.`
func (c *RestClient) InviteBot(target string, d *types.DataInviteBot) (*types.APIError, error) {
	apiErr, err := clientapi.NewReq(&c.Config).Post("bots/" + target + "/invite").Json(d).NoContentErr()
	return apiErr, err
}

// Fetch details of a bot you own by its id.
func (c *RestClient) FetchBot(target string) (*types.Bot, *types.APIError, error) {
	var bot *types.Bot
	apiErr, err := clientapi.NewReq(&c.Config).Get("bots/" + target).DoAndMarshal(&bot)
	return bot, apiErr, err
}

// Fetch all of the bots that you have control over.
func (c *RestClient) FetchOwnedBots() (*types.OwnedBotsResponse, *types.APIError, error) {
	var bot *types.OwnedBotsResponse
	apiErr, err := clientapi.NewReq(&c.Config).Get("bots/@me").DoAndMarshal(&bot)
	return bot, apiErr, err
}

// Delete a bot by its id.
func (c *RestClient) DeleteBot(target string) (*types.APIError, error) {
	apiErr, err := clientapi.NewReq(&c.Config).Delete("bots/" + target).NoContentErr()
	return apiErr, err
}

// Edit bot details by its id.
func (c *RestClient) EditBot(target string, d *types.DataEditBot) (*types.Bot, *types.APIError, error) {
	var bot *types.Bot
	apiErr, err := clientapi.NewReq(&c.Config).Patch("bots/" + target).Json(d).DoAndMarshal(&bot)
	return bot, apiErr, err
}
