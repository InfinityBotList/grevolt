// Package auth contains the struct for auth
package auth

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
}
