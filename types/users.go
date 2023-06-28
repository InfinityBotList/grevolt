package types

// UserList : List of users
type UserList []*User

// UserPermission : User permission definitions
type UserPermission string

// List of UserPermission
const (
	ACCESS_UserPermission       UserPermission = "Access"
	VIEW_PROFILE_UserPermission UserPermission = "ViewProfile"
	SEND_MESSAGE_UserPermission UserPermission = "SendMessage"
	INVITE_UserPermission       UserPermission = "Invite"
)

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

// FieldsUser : Optional fields on user object
type FieldsUser string

// List of FieldsUser
const (
	AVATAR_FieldsUser             FieldsUser = "Avatar"
	STATUS_TEXT_FieldsUser        FieldsUser = "StatusText"
	STATUS_PRESENCE_FieldsUser    FieldsUser = "StatusPresence"
	PROFILE_CONTENT_FieldsUser    FieldsUser = "ProfileContent"
	PROFILE_BACKGROUND_FieldsUser FieldsUser = "ProfileBackground"
	DISPLAY_NAME_FieldsUser       FieldsUser = "DisplayName"
)

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

type UserFlagResponse struct {
	// Flags
	Flags uint64 `json:"flags"`
}

type DataChangeUsername struct {
	// New username
	Username string `json:"username"`
	// Current account password
	Password string `json:"password"`
}
