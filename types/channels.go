package types

type ChannelList []Channel

// Channel struct.
type Channel struct {
	Id                 string                    `json:"_id,omitempty"`
	ChannelType        ChannelType               `json:"channel_type,omitempty"`
	UserId             string                    `json:"user,omitempty"`
	Nonce              string                    `json:"nonce,omitempty"`
	Active             bool                      `json:"active,omitempty"`
	Recipients         []string                  `json:"recipients,omitempty"`
	LastMessageID      string                    `json:"last_message_id,omitempty"`
	Name               string                    `json:"name,omitempty"`
	OwnerId            string                    `json:"owner,omitempty"`
	Description        string                    `json:"description,omitempty"`
	Icon               *File                     `json:"icon,omitempty"`
	DefaultPermissions *OverrideField            `json:"default_permissions,omitempty"`
	RolePermissions    map[string]*OverrideField `json:"role_permissions,omitempty"`
	Permissions        uint                      `json:"permissions,omitempty"`
	NSFW               bool                      `json:"nsfw,omitempty"`
}

type ChannelType string

// List of ChannelType
const (
	TEXT_ChannelType  ChannelType = "Text"
	VOICE_ChannelType ChannelType = "Voice"
)
