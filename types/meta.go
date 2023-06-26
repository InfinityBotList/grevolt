package types

// Rate limit struct
type RateLimit struct {
	RetryAfter int64 `json:"retry_after"`
}

// Revolt API configuration
type RevoltConfig struct {
	// Revolt API Version
	Revolt string `json:"revolt"`
	// Features enabled on this Revolt node
	Features *RevoltFeatures `json:"features"`
	// WebSocket URL
	Ws string `json:"ws"`
	// URL pointing to the client serving this node
	App string `json:"app"`
	// Web Push VAPID public key
	Vapid string `json:"vapid"`
	// Build information
	Build *RevoltConfigBuild `json:"build"`
}

// Features enabled on this Revolt node
type RevoltFeatures struct {
	// hCaptcha configuration
	Captcha *RevoltFeaturesCaptcha `json:"captcha"`
	// Whether email verification is enabled
	Email bool `json:"email"`
	// Whether this server is invite only
	InviteOnly bool `json:"invite_only"`
	// File server service configuration
	Autumn *RevoltFeaturesAutumn `json:"autumn"`
	// Proxy service configuration
	January *RevoltFeaturesJanuary `json:"january"`
	// Voice server configuration
	Voso *RevoltFeaturesVoso `json:"voso"`
}

// Build information
type RevoltConfigBuild struct {
	// Commit Hash
	CommitSha string `json:"commit_sha"`
	// Commit Timestamp
	CommitTimestamp string `json:"commit_timestamp"`
	// Git Semver
	Semver string `json:"semver"`
	// Git Origin URL
	OriginUrl string `json:"origin_url"`
	// Build Timestamp
	Timestamp string `json:"timestamp"`
}

// Features enabled on this Revolt node

// File server service configuration
type RevoltFeaturesAutumn struct {
	// Whether the service is enabled
	Enabled bool `json:"enabled"`
	// URL pointing to the service
	Url string `json:"url"`
}

// hCaptcha configuration
type RevoltFeaturesCaptcha struct {
	// Whether captcha is enabled
	Enabled bool `json:"enabled"`
	// Client key used for solving captcha
	Key string `json:"key"`
}

// Proxy service configuration
type RevoltFeaturesJanuary struct {
	// Whether the service is enabled
	Enabled bool `json:"enabled"`
	// URL pointing to the service
	Url string `json:"url"`
}

// Voice server configuration
type RevoltFeaturesVoso struct {
	// Whether voice is enabled
	Enabled bool `json:"enabled"`
	// URL pointing to the voice API
	Url string `json:"url"`
	// URL pointing to the voice WebSocket server
	Ws string `json:"ws"`
}

// Representation of a File on Revolt Generated by Autumn
type File struct {
	// Unique Id
	Id string `json:"_id"`
	// Tag / bucket this file was uploaded to
	Tag string `json:"tag"`
	// Original filename
	Filename string `json:"filename"`
	// Parsed metadata of this file
	Metadata *FileMetadata `json:"metadata"`
	// Raw content type of this file
	ContentType string `json:"content_type"`
	// Size of this file (in bytes)
	Size uint64 `json:"size"`
	// Whether this file was deleted
	Deleted bool `json:"deleted,omitempty"`
	// Whether this file was reported
	Reported bool `json:"reported,omitempty"`
	// Message Id this file is associated with
	MessageId string `json:"message_id,omitempty"`
	// User Id this file is associated with
	UserId string `json:"user_id,omitempty"`
	// Server Id this file is associated with
	ServerId string `json:"server_id,omitempty"`
	// Id of the object this file is associated with
	ObjectId string `json:"object_id,omitempty"`
}

type FileMetadata struct {
	// Type of this file
	Type string `json:"type"`

	// Width of the video
	Width int `json:"width"`

	// Height of the video
	Height int `json:"height"`
}

// API error, not properly generated
type APIError map[string]any

func (e APIError) Type() string {
	typ, ok := e["type"]

	if !ok {
		return "Unknown"
	}

	// Check if typ is a string
	if _, ok := typ.(string); !ok {
		return "Unknown"
	}

	return typ.(string)
}
