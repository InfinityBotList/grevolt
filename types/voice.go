package types

// Response from the 'Join Call' endpoint
type CreateVoiceUserResponse struct {
	// Token for authenticating with the voice server
	Token string `json:"token"`
}
