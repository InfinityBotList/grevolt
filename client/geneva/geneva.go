// Package geneva contains common helpers and types for use in both rest and websocket configs
package geneva

// Token is a session token used for authorization
type Token struct {
	// Whether the token is a bot token
	Bot bool
	// The session token
	Token string
}
