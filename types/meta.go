package types

// Rate limit struct
type RateLimit struct {
	RetryAfter int64 `json:"retry_after"`
}
