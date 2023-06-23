// Package auth contains the struct for auth, this is required due to import cycles
package auth

// Token is a session token used for authorization
type Token struct {
	// Whether the token is a bot token
	Bot bool
	// The session token
	Token string
}
