// Package auth contains the struct for auth
package auth

// CfClearance data
type CfClearance struct {
	// User agent on which the cloudflare challenge was made
	UserAgent string

	// Cookie value returned from the cloudflare challenge
	CookieValue string
}

// Token is a session token used for authorization
type Token struct {
	// Whether the token is a bot token
	//
	// Note: this is ignored on gateway, but required for making REST requests
	//
	// Also, note that preparing a websocket connection does require REST
	// and so this field must be properly set anyways
	Bot bool

	// The session token
	Token string

	// CfClearance token
	//
	// In the event of under attack mode, you can complete this challenge yourself on your browser
	// and manually provide a cfclearance token to get your bot running provided the bot and you
	// are on the same IP
	CfClearance *CfClearance
}
