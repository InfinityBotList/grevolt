package types

// FieldsBot : Optional fields on bot object
type FieldsBot string

// List of FieldsBot
const (
	TOKEN_FieldsBot            FieldsBot = "Token"
	INTERACTIONS_URL_FieldsBot FieldsBot = "InteractionsURL"
)

// Bot information for if the user is a bot
type BotInformation struct {
	// Id of the owner of this bot
	Owner string `json:"owner"`
}

// Representation of a bot on Revolt
type Bot struct {
	// Bot Id  This equals the associated bot user's id.
	Id string `json:"_id"`
	// User Id of the bot owner
	Owner string `json:"owner"`
	// Token used to authenticate requests for this bot
	Token string `json:"token"`
	// Whether the bot is public (may be invited by anyone)
	Public bool `json:"public"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this bot should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
	// Reserved; URL for handling interactions
	InteractionsUrl string `json:"interactions_url,omitempty"`
	// URL for terms of service
	TermsOfServiceUrl string `json:"terms_of_service_url,omitempty"`
	// URL for privacy policy
	PrivacyPolicyUrl string `json:"privacy_policy_url,omitempty"`
	// Enum of bot flags
	Flags uint64 `json:"flags,omitempty"`
}

// Public Bot
type PublicBot struct {
	// Bot Id
	Id string `json:"_id"`
	// Bot Username
	Username string `json:"username"`
	// Profile Avatar
	Avatar string `json:"avatar"`
	// Profile Description
	Description string `json:"description"`
}

// Both lists are sorted by their IDs.
type OwnedBotsResponse struct {
	// Bot objects
	Bots []*Bot `json:"bots"`
	// User objects
	Users []*User `json:"users"`
}

// Bot Response
type FetchBotResponse struct {
	// Bot object
	Bot *Bot `json:"bot"`
	// User object
	User *User `json:"user"`
}

// Data needed to invite a bot to a server or group
//
// <official docs seem to miss server but its there>
type DataInviteBot struct {
	// Server Id
	Server string `json:"server,omitempty"`
	// Group id
	Group string `json:"group,omitempty"`
}

// Data needed to create a bot
//
// <unless specified, all fields are required>
type DataCreateBot struct {
	// Bot username
	Name string `json:"name"`
}

// Data needed to edit a bot
type DataEditBot struct {
	// Bot username
	Name string `json:"name,omitempty"`
	// Whether the bot can be added by anyone
	Public bool `json:"public,omitempty"`
	// Whether analytics should be gathered for this bot  Must be enabled in order to show up on [Revolt Discover](https://rvlt.gg).
	Analytics bool `json:"analytics,omitempty"`
	// Interactions URL
	InteractionsUrl string `json:"interactions_url,omitempty"`
	// Fields to remove from bot object
	Remove []*FieldsBot `json:"remove,omitempty"`
}
