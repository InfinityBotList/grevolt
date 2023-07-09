package types

type ChannelList []Channel

type ChannelType string

// List of ChannelType
const (
	TEXT_ChannelType  ChannelType = "Text"
	VOICE_ChannelType ChannelType = "Voice"
)

// FieldsChannel : Optional fields on channel object
type FieldsChannel string

// List of FieldsChannel
const (
	DESCRIPTION_FieldsChannel         FieldsChannel = "Description"
	ICON_FieldsChannel                FieldsChannel = "Icon"
	DEFAULT_PERMISSIONS_FieldsChannel FieldsChannel = "DefaultPermissions"
)

// Channel struct.
type Channel struct {
	Id                 string                              `json:"_id,omitempty"`
	ChannelType        ChannelType                         `json:"channel_type,omitempty"`
	Server             string                              `json:"server,omitempty"`
	UserId             string                              `json:"user,omitempty"`
	Nonce              string                              `json:"nonce,omitempty"`
	Active             bool                                `json:"active,omitempty"`
	Recipients         []string                            `json:"recipients,omitempty"`
	LastMessageID      string                              `json:"last_message_id,omitempty"`
	Name               string                              `json:"name,omitempty"`
	OwnerId            string                              `json:"owner,omitempty"`
	Description        string                              `json:"description,omitempty"`
	Icon               *File                               `json:"icon,omitempty"`
	DefaultPermissions *PermissionOverrideField            `json:"default_permissions,omitempty"`
	RolePermissions    map[string]*PermissionOverrideField `json:"role_permissions,omitempty"`
	Permissions        uint                                `json:"permissions,omitempty"`
	NSFW               bool                                `json:"nsfw,omitempty"`
}

// Data for creating a channel
type DataCreateChannel struct {
	// Channel type
	Type ChannelType `json:"type,omitempty"`
	// Channel name
	//
	// <this is required for creating a channel>
	Name string `json:"name,omitempty"`
	// Channel description
	Description string `json:"description,omitempty"`
	// Whether this channel is age restricted
	Nsfw bool `json:"nsfw,omitempty"`
}

// Data for editing a channel
type DataEditChannel struct {
	// Channel name
	Name string `json:"name,omitempty"`
	// Channel description
	Description string `json:"description,omitempty"`
	// Group owner
	Owner string `json:"owner,omitempty"`
	// Icon  Provide an Autumn attachment Id.
	Icon string `json:"icon,omitempty"`
	// Whether this channel is age-restricted
	Nsfw bool `json:"nsfw,omitempty"`
	// Whether this channel is archived
	Archived bool `json:"archived,omitempty"`
	// Fields to remove from the channel
	Remove []FieldsChannel `json:"remove,omitempty"`
}

// System message channel assignments
type SystemMessageChannels struct {
	// ID of channel to send user join messages in
	UserJoined string `json:"user_joined,omitempty"`
	// ID of channel to send user left messages in
	UserLeft string `json:"user_left,omitempty"`
	// ID of channel to send user kicked messages in
	UserKicked string `json:"user_kicked,omitempty"`
	// ID of channel to send user banned messages in
	UserBanned string `json:"user_banned,omitempty"`
}

// Channel category
type Category struct {
	// Unique ID for this category
	Id string `json:"id"`
	// Title for this category
	Title string `json:"title"`
	// Channels in this category
	Channels []string `json:"channels"`
}
