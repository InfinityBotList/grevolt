package types

// MfaMethod : MFA method
type MfaMethod string

// List of MFAMethod
const (
	PASSWORD_MfaMethod MfaMethod = "Password"
	RECOVERY_MfaMethod MfaMethod = "Recovery"
	TOTP_MfaMethod     MfaMethod = "Totp"
)

// Multi-factor auth ticket
type MfaTicket struct {
	// Unique Id
	Id string `json:"_id"`
	// Account Id
	AccountId string `json:"account_id"`
	// Unique Token
	Token string `json:"token"`
	// Whether this ticket has been validated (can be used for account actions)
	Validated bool `json:"validated"`
	// Whether this ticket is authorised (can be used to log a user in)
	Authorised bool `json:"authorised"`
	// TOTP code at time of ticket creation
	LastTotpCode string `json:"last_totp_code,omitempty"`
}

type MultiFactorStatus struct {
	EmailOtp        bool `json:"email_otp"`
	TrustedHandover bool `json:"trusted_handover"`
	EmailMfa        bool `json:"email_mfa"`
	TotpMfa         bool `json:"totp_mfa"`
	SecurityKeyMfa  bool `json:"security_key_mfa"`
	RecoveryActive  bool `json:"recovery_active"`
}

type ResponseTotpSecret struct {
	Secret string `json:"secret"`
}
