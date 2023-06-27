package types

// Web Push subscription, all fields are mandatory/required
type WebPushSubscription struct {
	Endpoint string `json:"endpoint"`
	P256dh   string `json:"p256dh"`
	Auth     string `json:"auth"`
}
