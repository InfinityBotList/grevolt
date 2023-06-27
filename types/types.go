package types

type DataLogin struct {
	// Either

	// Email
	Email string `json:"email,omitempty"`

	// Password
	Password string `json:"password,omitempty"`

	// Or

	// Unvalidated or authorised MFA ticket
	//
	// Used to resolve the correct account
	MfaTicket string `json:"mfa_ticket,omitempty"`

	// MFA response
	MfaResponse *DataLoginMfaResponse `json:"mfa_response,omitempty"`
}

// Note, you must specify one (and only one) of the following fields if specifiying this
type DataLoginMfaResponse struct {
	// Password
	Password string `json:"password,omitempty"`
	// Recovery Code
	RecoveryCode string `json:"recovery_code,omitempty"`
	// TOTP code
	TotpCode string `json:"totp_code,omitempty"`
}

type DataEditSession struct {
	// Session friendly name
	FriendlyName string `json:"friendly_name"`
}

// Representation of a session on Revolt
type SessionInfo struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}
