package types

type AccountInfo struct {
	Id    string `json:"_id"`
	Email string `json:"email"`
}

// Data needed to resend verification email or send password reset email.
type DataResendVerificationSendPasswordReset struct {
	// Email associated with the account
	Email string `json:"email"`
	// Captcha verification code
	Captcha string `json:"captcha,omitempty"`
}

// Data needed to create a new account.
type DataCreateAccount struct {
	// Valid email address
	Email string `json:"email"`
	// Password
	Password string `json:"password"`
	// Invite code
	//
	// <optional, not all Revolt nodes need it>
	Invite string `json:"invite,omitempty"`
	// Captcha verification code
	//
	// <optional, not all Revolt nodes need it>
	Captcha string `json:"captcha,omitempty"`
}

// Data needed to confirm password reset and change the password.
type DataConfirmPasswordReset struct {
	// Reset token
	Token string `json:"token"`
	// New password
	Password string `json:"password"`
	// Whether to logout all sessions
	RemoveSessions bool `json:"remove_sessions,omitempty"`
}

// Data needed to change the email address.
type DataChangeEmail struct {
	// Valid email address
	Email string `json:"email"`
	// Current password
	CurrentPassword string `json:"current_password"`
}

// Data needed to change the password.
type DataChangePassword struct {
	// New password
	Password string `json:"password"`
	// Current password
	CurrentPassword string `json:"current_password"`
}

// Data needed to verify an email
type DataVerifyEmail struct {
	// Multi-factor auth ticket
	Ticket *MfaTicket `json:"ticket"`
}

// Data needed to send a password reset email.
type DataSendPasswordReset struct {
	// Email associated with the account
	Email string `json:"email"`
	// Captcha verification code
	Captcha string `json:"captcha,omitempty"`
}

// Response to an account deletion confirmation request.
type ConfirmAccountDeletionResponse struct {
	// Deletion token
	Token string `json:"token"`
}
