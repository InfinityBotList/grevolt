package types

// Representiation of a User on Revolt.
type User struct {
	// Unique Id
	Id string `json:"_id"`
	// Username
	Username string `json:"username"`
	// Discriminator
	Discriminator string `json:"discriminator"`
	// Display name
	DisplayName string `json:"display_name,omitempty"`
	// Avatar attachment
	Avatar *File `json:"avatar,omitempty"`
	// Relationships with other users
	Relations []Relationship `json:"relations,omitempty"`
	// Bitfield of user badges
	Badges uint64 `json:"badges,omitempty"`
	// User's current status
	Status *UserStatus `json:"status,omitempty"`
	// User's profile page
	Profile *UserProfile `json:"profile,omitempty"`
	// Enum of user flags
	Flags uint64 `json:"flags,omitempty"`
	// Whether this user is privileged
	Privileged bool `json:"privileged,omitempty"`
	// Bot information
	Bot *BotInformation `json:"bot,omitempty"`
	// Current session user's relationship with this user
	Relationship RelationshipStatus `json:"relationship,omitempty"`
	// Whether this user is currently online
	Online bool `json:"online,omitempty"`
}

// UserPermission : User permission definitions
type UserPermission string

// List of UserPermission
const (
	ACCESS_UserPermission       UserPermission = "Access"
	VIEW_PROFILE_UserPermission UserPermission = "ViewProfile"
	SEND_MESSAGE_UserPermission UserPermission = "SendMessage"
	INVITE_UserPermission       UserPermission = "Invite"
)

// User's profile
type UserProfile struct {
	// Text content on user's profile
	Content string `json:"content,omitempty"`
	// Background visible on user's profile
	Background *File `json:"background,omitempty"`
}

type UserProfileData struct {
	// Text to set as user profile description
	Content string `json:"content,omitempty"`
	// Attachment Id for background
	Background string `json:"background,omitempty"`
}

// User's active status
type UserStatus struct {
	// Custom status text
	Text string `json:"text,omitempty"`
	// Current presence option
	Presence Presence `json:"presence,omitempty"`
}

// Presence : Presence status
type Presence string

// List of Presence
const (
	ONLINE_Presence    Presence = "Online"
	IDLE_Presence      Presence = "Idle"
	FOCUS_Presence     Presence = "Focus"
	BUSY_Presence      Presence = "Busy"
	INVISIBLE_Presence Presence = "Invisible"
)

// System message configuration

// New user profile data  This is applied as a partial.
type DataEditUserProfile struct {
	// Text to set as user profile description
	Content string `json:"content,omitempty"`
	// Attachment Id for background
	Background string `json:"background,omitempty"`
}

type DataEditUser struct {
	// New display name
	DisplayName string `json:"display_name,omitempty"`
	// Attachment Id for avatar
	Avatar string `json:"avatar,omitempty"`
	// New user status
	Status *UserStatus `json:"status,omitempty"`
	// New user profile data  This is applied as a partial.
	Profile *DataEditUserProfile `json:"profile,omitempty"`
	// Bitfield of user badges
	Badges uint64 `json:"badges,omitempty"`
	// Enum of user flags
	Flags uint64 `json:"flags,omitempty"`
	// Fields to remove from user object
	Remove []FieldsUser `json:"remove,omitempty"`
}
