package types

// This will tell you whether the current account requires onboarding or whether you can continue to send requests as usual.
//
// You may skip calling this if you're restoring an existing session.
type DataHello struct {
	// Whether onboarding is required
	Onboarding bool `json:"onboarding"`
}

// This sets a new username, completes onboarding and allows a user to start using Revolt.
type DataOnboard struct {
	// New username which will be used to identify the user on the platform
	Username string `json:"username"`
}
