package types

type Object map[string]any

// Emoji parent, not properly generated
type EmojiParent struct {
	Type string `json:"type"`
	ID   string `json:"id"`
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

// Channel struct.
type Channel struct {
	Id                 string         `json:"_id,omitempty"`
	ChannelType        string         `json:"channel_type,omitempty"`
	UserId             string         `json:"user,omitempty"`
	Nonce              string         `json:"nonce,omitempty"`
	Active             bool           `json:"active,omitempty"`
	Recipients         []string       `json:"recipients,omitempty"`
	LastMessageID      string         `json:"last_message_id,omitempty"`
	Name               string         `json:"name,omitempty"`
	OwnerId            string         `json:"owner,omitempty"`
	Description        string         `json:"description,omitempty"`
	Icon               *File          `json:"icon,omitempty"`
	DefaultPermissions *OverrideField `json:"default_permissions,omitempty"`
	RolePermissions    interface{}    `json:"role_permissions,omitempty"`
	Permissions        uint           `json:"permissions,omitempty"`
	NSFW               bool           `json:"nsfw,omitempty"`
}

type ChannelList []Channel

type ChannelType string

// List of ChannelType
const (
	TEXT_ChannelType ChannelType = "Text"
	VOICE_ChannelType ChannelType = "Voice"
)

type DataInviteBot struct {
	// Server Id
	Server string `json:"server,omitempty"`
	// Group id
	Group string `json:"group,omitempty"`
}

// Begin types

type AccountInfo struct {
	Id string `json:"_id"`
	Email string `json:"email"`
}

// Account Strike
type AccountStrike struct {
	// Strike Id
	Id string `json:"_id"`
	// Id of reported user
	UserId string `json:"user_id"`
	// Attached reason
	Reason string `json:"reason"`
}

// Both lists are sorted by ID.
type AllMemberResponse struct {
	// List of members
	Members []Member `json:"members"`
	// List of users
	Users []User `json:"users"`
}

// Avatar of the banned user

// Composite key pointing to a user's view of a channel
type ChannelUnreadId struct {
	// Channel Id
	Channel string `json:"channel"`
	// User Id
	User string `json:"user"`
}

// Local time
type CollectionStatsLocalTime struct {
}

// Query exec stats
type CollectionStatsQueryExecStats struct {
	// Stats regarding collection scans
	CollectionScans *Object `json:"collectionScans"`
}

// Server object
type CreateServerResponseServer struct {
	// Unique Id
	Id string `json:"_id"`
	// User id of the owner
	Owner string `json:"owner"`
	// Name of the server
	Name string `json:"name"`
	// Description for the server
	Description string `json:"description,omitempty"`
	// Channels within this server
	Channels []string `json:"channels"`
	// Categories for this server
	Categories []Category `json:"categories,omitempty"`
	// Configuration for sending system event messages
	SystemMessages *ServerSystemMessages `json:"system_messages,omitempty"`
	// Roles for this server
	Roles map[string]Role `json:"roles,omitempty"`
	// Default set of server and channel permissions
	DefaultPermissions int64 `json:"default_permissions"`
	// Icon attachment
	Icon *Object `json:"icon,omitempty"`
	// Banner attachment
	Banner *Object `json:"banner,omitempty"`
	// Bitfield of server flags
	Flags int64 `json:"flags,omitempty"`
	// Whether this server is flagged as not safe for work
	Nsfw bool `json:"nsfw,omitempty"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this server should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
}

// Channel type
type DataCreateChannelType_ struct {
}

// Parent information
type DataCreateEmojiParent struct {
}

// New report status
type DataEditReportStatus struct {
}

// System message configuration

// New user profile data  This is applied as a partial.
type DataEditUserProfile struct {
	// Text to set as user profile description
	Content string `json:"content,omitempty"`
	// Attachment Id for background
	Background string `json:"background,omitempty"`
}

// New user status

// Timestamp this member is timed out until
type DataMemberEditTimeout struct {
}

// Information about how this message should be interacted with
type DataMessageSendInteractions struct {
	// Reactions which should always appear and be distinct
	Reactions []string `json:"reactions,omitempty"`
	// Whether reactions should be restricted to the given list  Can only be set to true if reactions list is of at least length 1
	RestrictReactions bool `json:"restrict_reactions,omitempty"`
}

// Masquerade to apply to this message
type DataMessageSendMasquerade struct {
	// Replace the display name shown on this message
	Name string `json:"name,omitempty"`
	// Replace the avatar shown on this message (URL to image file)
	Avatar string `json:"avatar,omitempty"`
	// Replace the display role colour shown on this message  Must have `ManageRole` permission to use
	Colour string `json:"colour,omitempty"`
}

// Allow / deny values to set for this role
type DataPermissions struct {
	// Allow bit flags
	Allow int64 `json:"allow"`
	// Disallow bit flags
	Deny int64 `json:"deny"`
}

// Content being reported
type DataReportContentContent struct {
}

// Allow / deny values for the role in this server.
type DataSetServerRolePermissionPermissions struct {
	// Allow bit flags
	Allow int64 `json:"allow"`
	// Disallow bit flags
	Deny int64 `json:"deny"`
}

// What owns this emoji

// Bot object
type FetchBotResponseBot struct {
	// Bot Id  This equals the associated bot user's id.
	Id string `json:"_id"`
	// User Id of the bot owner
	Owner string `json:"owner"`
	// Token used to authenticate requests for this bot
	Token string `json:"token"`
	// Whether the bot is public (may be invited by anyone)
	Public bool `json:"public"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this bot should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
	// Reserved; URL for handling interactions
	InteractionsUrl string `json:"interactions_url,omitempty"`
	// URL for terms of service
	TermsOfServiceUrl string `json:"terms_of_service_url,omitempty"`
	// URL for privacy policy
	PrivacyPolicyUrl string `json:"privacy_policy_url,omitempty"`
	// Enum of bot flags
	Flags int64 `json:"flags,omitempty"`
}

// User object
type FetchBotResponseUser struct {
	// Unique Id
	Id string `json:"_id"`
	// Username
	Username string `json:"username"`
	// Discriminator
	Discriminator string `json:"discriminator"`
	// Display name
	DisplayName string `json:"display_name,omitempty"`
	// Avatar attachment
	Avatar *Object `json:"avatar,omitempty"`
	// Relationships with other users
	Relations []Relationship `json:"relations,omitempty"`
	// Bitfield of user badges
	Badges int64 `json:"badges,omitempty"`
	// User's current status
	Status *Object `json:"status,omitempty"`
	// User's profile page
	Profile *Object `json:"profile,omitempty"`
	// Enum of user flags
	Flags int64 `json:"flags,omitempty"`
	// Whether this user is privileged
	Privileged bool `json:"privileged,omitempty"`
	// Bot information
	Bot *Object `json:"bot,omitempty"`
	// Current session user's relationship with this user
	Relationship *Object `json:"relationship,omitempty"`
	// Whether this user is currently online
	Online bool `json:"online,omitempty"`
}

// Parsed metadata of this file
type FileMetadata struct {
}

// Positioning and size

// Timestamp at which data keeping begun
type IndexAccessSince struct {
}

// Access information
type IndexAccesses struct {
	// Operations since timestamp
	Ops int64 `json:"ops"`
	// Timestamp at which data keeping begun
	Since *Object `json:"since"`
}

// Avatar attachment

// Unique member id
type MemberId struct {
	// Server Id
	Server string `json:"server"`
	// User Id
	User string `json:"user"`
}

// Time at which this user joined the server
type MemberJoinedAt struct {
}

// Timestamp this member is timed out until
type MemberTimeout struct {
}

// Time at which this message was last edited
type MessageEdited struct {
}

// Information about how this message should be interacted with
type MessageInteractions struct {
	// Reactions which should always appear and be distinct
	Reactions []string `json:"reactions,omitempty"`
	// Whether reactions should be restricted to the given list  Can only be set to true if reactions list is of at least length 1
	RestrictReactions bool `json:"restrict_reactions,omitempty"`
}

// Name and / or avatar overrides for this message
type MessageMasquerade struct {
	// Replace the display name shown on this message
	Name string `json:"name,omitempty"`
	// Replace the avatar shown on this message (URL to image file)
	Avatar string `json:"avatar,omitempty"`
	// Replace the display role colour shown on this message  Must have `ManageRole` permission to use
	Colour string `json:"colour,omitempty"`
}

// System message
type MessageSystem struct {
}

// The webhook that sent this message

// New role
type NewRoleResponseRole struct {
	// Role name
	Name string `json:"name"`
	// Permissions available to this role
	Permissions *Object `json:"permissions"`
	// Colour used for this role  This can be any valid CSS colour
	Colour string `json:"colour,omitempty"`
	// Whether this role should be shown separately on the member sidebar
	Hoist bool `json:"hoist,omitempty"`
	// Ranking of this role
	Rank int64 `json:"rank,omitempty"`
}

// Message sort direction  By default, it will be sorted by latest.
type OptionsMessageSearchSort struct {
}

// Stats regarding collection scans
type QueryExecStatsCollectionScans struct {
	// Number of total collection scans
	Total int64 `json:"total"`
	// Number of total collection scans not using a tailable cursor
	NonTailable int64 `json:"nonTailable"`
}

// Reported content
type ReportContent struct {
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
type RevoltConfigFeatures struct {
	// hCaptcha configuration
	Captcha *Object `json:"captcha"`
	// Whether email verification is enabled
	Email bool `json:"email"`
	// Whether this server is invite only
	InviteOnly bool `json:"invite_only"`
	// File server service configuration
	Autumn *Object `json:"autumn"`
	// Proxy service configuration
	January *Object `json:"january"`
	// Voice server configuration
	Voso *Object `json:"voso"`
}

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

// Permissions available to this role
type RolePermissions struct {
	// Allow bit flags
	A int64 `json:"a"`
	// Disallow bit flags
	D int64 `json:"d"`
}

// Unique member id
type ServerBanId struct {
	// Server Id
	Server string `json:"server"`
	// User Id
	User string `json:"user"`
}

// Banner attachment

// Icon attachment

// Configuration for sending system event messages
type ServerSystemMessages struct {
	// ID of channel to send user join messages in
	UserJoined string `json:"user_joined,omitempty"`
	// ID of channel to send user left messages in
	UserLeft string `json:"user_left,omitempty"`
	// ID of channel to send user kicked messages in
	UserKicked string `json:"user_kicked,omitempty"`
	// ID of channel to send user banned messages in
	UserBanned string `json:"user_banned,omitempty"`
}

// Snapshot of content

// Server involved in snapshot
type SnapshotWithContextServer struct {
	// Unique Id
	Id string `json:"_id"`
	// User id of the owner
	Owner string `json:"owner"`
	// Name of the server
	Name string `json:"name"`
	// Description for the server
	Description string `json:"description,omitempty"`
	// Channels within this server
	Channels []string `json:"channels"`
	// Categories for this server
	Categories []Category `json:"categories,omitempty"`
	// Configuration for sending system event messages
	SystemMessages *ServerSystemMessages `json:"system_messages,omitempty"`
	// Roles for this server
	Roles map[string]Role `json:"roles,omitempty"`
	// Default set of server and channel permissions
	DefaultPermissions int64 `json:"default_permissions"`
	// Icon attachment
	Icon *Object `json:"icon,omitempty"`
	// Banner attachment
	Banner *Object `json:"banner,omitempty"`
	// Bitfield of server flags
	Flags int64 `json:"flags,omitempty"`
	// Whether this server is flagged as not safe for work
	Nsfw bool `json:"nsfw,omitempty"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this server should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
}

// Avatar attachment

// Bot information
type UserBot struct {
	// Id of the owner of this bot
	Owner string `json:"owner"`
}

// User's profile page

// Background visible on user's profile

// Current session user's relationship with this user

// User's current status

// Current presence option

// The avatar of the webhook

type AnyOfMessageQuery struct {
}

type AuthifierError struct {
}

type BanListResult struct {
	// Users objects
	Users []BannedUser `json:"users"`
	// Ban objects
	Bans []ServerBan `json:"bans"`
}
// BandcampType : Type of remote Bandcamp content
type BandcampType string

// List of BandcampType
const (
	ALBUM_BandcampType BandcampType = "Album"
	TRACK_BandcampType BandcampType = "Track"
)

// Just enoguh user information to list bans.
type BannedUser struct {
	// Id of the banned user
	Id string `json:"_id"`
	// Username of the banned user
	Username string `json:"username"`
	// Avatar of the banned user
	Avatar *File `json:"avatar,omitempty"`
}

// Representation of a bot on Revolt
type Bot struct {
	// Bot Id  This equals the associated bot user's id.
	Id string `json:"_id"`
	// User Id of the bot owner
	Owner string `json:"owner"`
	// Token used to authenticate requests for this bot
	Token string `json:"token"`
	// Whether the bot is public (may be invited by anyone)
	Public bool `json:"public"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this bot should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
	// Reserved; URL for handling interactions
	InteractionsUrl string `json:"interactions_url,omitempty"`
	// URL for terms of service
	TermsOfServiceUrl string `json:"terms_of_service_url,omitempty"`
	// URL for privacy policy
	PrivacyPolicyUrl string `json:"privacy_policy_url,omitempty"`
	// Enum of bot flags
	Flags int64 `json:"flags,omitempty"`
}

// Bot information for if the user is a bot
type BotInformation struct {
	// Id of the owner of this bot
	Owner string `json:"owner"`
}

type BuildInformation struct {
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

// Response used when multiple messages are fetched
type BulkMessageResponse struct {
}

type CaptchaFeature struct {
	// Whether captcha is enabled
	Enabled bool `json:"enabled"`
	// Client key used for solving captcha
	Key string `json:"key"`
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

// Representation of a channel on Revolt

// Composite primary key consisting of channel and user id


// Query collection scan stats
type CollectionScans struct {
	// Number of total collection scans
	Total int64 `json:"total"`
	// Number of total collection scans not using a tailable cursor
	NonTailable int64 `json:"nonTailable"`
}

// Collection stats
type CollectionStats struct {
	// Namespace
	Ns string `json:"ns"`
	// Local time
	LocalTime *CollectionStatsLocalTime `json:"localTime"`
	// Latency stats
	LatencyStats map[string]LatencyStats `json:"latencyStats"`
	// Query exec stats
	QueryExecStats *CollectionStatsQueryExecStats `json:"queryExecStats"`
	// Number of documents in collection
	Count int64 `json:"count"`
}
// ContentReportReason : Reason for reporting content (message or server)
type ContentReportReason string

// List of ContentReportReason
const (
	NONE_SPECIFIED_ContentReportReason ContentReportReason = "NoneSpecified"
	ILLEGAL_ContentReportReason ContentReportReason = "Illegal"
	ILLEGAL_GOODS_ContentReportReason ContentReportReason = "IllegalGoods"
	ILLEGAL_EXTORTION_ContentReportReason ContentReportReason = "IllegalExtortion"
	ILLEGAL_PORNOGRAPHY_ContentReportReason ContentReportReason = "IllegalPornography"
	ILLEGAL_HACKING_ContentReportReason ContentReportReason = "IllegalHacking"
	EXTREME_VIOLENCE_ContentReportReason ContentReportReason = "ExtremeViolence"
	PROMOTES_HARM_ContentReportReason ContentReportReason = "PromotesHarm"
	UNSOLICITED_SPAM_ContentReportReason ContentReportReason = "UnsolicitedSpam"
	RAID_ContentReportReason ContentReportReason = "Raid"
	SPAM_ABUSE_ContentReportReason ContentReportReason = "SpamAbuse"
	SCAMS_FRAUD_ContentReportReason ContentReportReason = "ScamsFraud"
	MALWARE_ContentReportReason ContentReportReason = "Malware"
	HARASSMENT_ContentReportReason ContentReportReason = "Harassment"
)

type CreateServerResponse struct {
	// Server object
	Server *CreateServerResponseServer `json:"server"`
	// Default channels
	Channels []Channel `json:"channels"`
}

type CreateVoiceUserResponse struct {
	// Token for authenticating with the voice server
	Token string `json:"token"`
}

type CreateWebhookBody struct {
	Name string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}

type Data struct {
	// Allow / deny values to set for this role
	Permissions *DataPermissions `json:"permissions"`
}

type DataAccountDeletion struct {
	// Deletion token
	Token string `json:"token"`
}

type DataBanCreate struct {
	// Ban reason
	Reason string `json:"reason,omitempty"`
}

type DataChangeEmail struct {
	// Valid email address
	Email string `json:"email"`
	// Current password
	CurrentPassword string `json:"current_password"`
}

type DataChangePassword struct {
	// New password
	Password string `json:"password"`
	// Current password
	CurrentPassword string `json:"current_password"`
}

type DataChangeUsername struct {
	// New username
	Username string `json:"username"`
	// Current account password
	Password string `json:"password"`
}

type DataCreateAccount struct {
	// Valid email address
	Email string `json:"email"`
	// Password
	Password string `json:"password"`
	// Invite code
	Invite string `json:"invite,omitempty"`
	// Captcha verification code
	Captcha string `json:"captcha,omitempty"`
}

type DataCreateBot struct {
	// Bot username
	Name string `json:"name"`
}

type DataCreateChannel struct {
	// Channel type
	Type_ *DataCreateChannelType_ `json:"type,omitempty"`
	// Channel name
	Name string `json:"name"`
	// Channel description
	Description string `json:"description,omitempty"`
	// Whether this channel is age restricted
	Nsfw bool `json:"nsfw,omitempty"`
}

type DataCreateEmoji struct {
	// Server name
	Name string `json:"name"`
	// Parent information
	Parent *DataCreateEmojiParent `json:"parent"`
	// Whether the emoji is mature
	Nsfw bool `json:"nsfw,omitempty"`
}

type DataCreateGroup struct {
	// Group name
	Name string `json:"name"`
	// Group description
	Description string `json:"description,omitempty"`
	// Array of user IDs to add to the group  Must be friends with these users.
	Users []string `json:"users"`
	// Whether this group is age-restricted
	Nsfw bool `json:"nsfw,omitempty"`
}

type DataCreateRole struct {
	// Role name
	Name string `json:"name"`
	// Ranking position  Smaller values take priority.
	Rank int64 `json:"rank,omitempty"`
}

type DataCreateServer struct {
	// Server name
	Name string `json:"name"`
	// Server description
	Description string `json:"description,omitempty"`
	// Whether this server is age-restricted
	Nsfw bool `json:"nsfw,omitempty"`
}

// New strike information
type DataCreateStrike struct {
	// Id of reported user
	UserId string `json:"user_id"`
	// Attached reason
	Reason string `json:"reason"`
}

type DataDefaultChannelPermissions struct {
}

// New strike information
type DataEditAccountStrike struct {
	// New attached reason
	Reason string `json:"reason"`
}

type DataEditBot struct {
	// Bot username
	Name string `json:"name,omitempty"`
	// Whether the bot can be added by anyone
	Public bool `json:"public,omitempty"`
	// Whether analytics should be gathered for this bot  Must be enabled in order to show up on [Revolt Discover](https://rvlt.gg).
	Analytics bool `json:"analytics,omitempty"`
	// Interactions URL
	InteractionsUrl string `json:"interactions_url,omitempty"`
	// Fields to remove from bot object
	Remove []FieldsBot `json:"remove,omitempty"`
}

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
	Remove []FieldsChannel `json:"remove,omitempty"`
}

type DataEditMessage struct {
	// New message content
	Content string `json:"content,omitempty"`
	// Embeds to include in the message
	Embeds []SendableEmbed `json:"embeds,omitempty"`
}

type DataEditReport struct {
	// New report status
	Status *DataEditReportStatus `json:"status,omitempty"`
	// Report notes
	Notes string `json:"notes,omitempty"`
}

type DataEditRole struct {
	// Role name
	Name string `json:"name,omitempty"`
	// Role colour
	Colour string `json:"colour,omitempty"`
	// Whether this role should be displayed separately
	Hoist bool `json:"hoist,omitempty"`
	// Ranking position  Smaller values take priority.
	Rank int64 `json:"rank,omitempty"`
	// Fields to remove from role object
	Remove []FieldsRole `json:"remove,omitempty"`
}

type DataEditServer struct {
	// Server name
	Name string `json:"name,omitempty"`
	// Server description
	Description string `json:"description,omitempty"`
	// Attachment Id for icon
	Icon string `json:"icon,omitempty"`
	// Attachment Id for banner
	Banner string `json:"banner,omitempty"`
	// Category structure for server
	Categories []Category `json:"categories,omitempty"`
	// System message configuration
	SystemMessages *ServerSystemMessages `json:"system_messages,omitempty"`
	// Bitfield of server flags
	Flags int64 `json:"flags,omitempty"`
	// Whether this server is public and should show up on [Revolt Discover](https://rvlt.gg)
	Discoverable bool `json:"discoverable,omitempty"`
	// Whether analytics should be collected for this server  Must be enabled in order to show up on [Revolt Discover](https://rvlt.gg).
	Analytics bool `json:"analytics,omitempty"`
	// Fields to remove from server object
	Remove []FieldsServer `json:"remove,omitempty"`
}

type DataEditSession struct {
	// Session friendly name
	FriendlyName string `json:"friendly_name"`
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
	Badges int64 `json:"badges,omitempty"`
	// Enum of user flags
	Flags int64 `json:"flags,omitempty"`
	// Fields to remove from user object
	Remove []FieldsUser `json:"remove,omitempty"`
}

type DataHello struct {
	// Whether onboarding is required
	Onboarding bool `json:"onboarding"`
}

type DataLogin struct {
}

type DataMemberEdit struct {
	// Member nickname
	Nickname string `json:"nickname,omitempty"`
	// Attachment Id to set for avatar
	Avatar string `json:"avatar,omitempty"`
	// Array of role ids
	Roles []string `json:"roles,omitempty"`
	// Timestamp this member is timed out until
	Timeout *DataMemberEditTimeout `json:"timeout,omitempty"`
	// Fields to remove from channel object
	Remove []FieldsMember `json:"remove,omitempty"`
}

type DataMessageSend struct {
	// Unique token to prevent duplicate message sending  **This is deprecated and replaced by `Idempotency-Key`!**
	Nonce string `json:"nonce,omitempty"`
	// Message content to send
	Content string `json:"content,omitempty"`
	// Attachments to include in message
	Attachments []string `json:"attachments,omitempty"`
	// Messages to reply to
	Replies []Reply `json:"replies,omitempty"`
	// Embeds to include in message  Text embed content contributes to the content length cap
	Embeds []SendableEmbed `json:"embeds,omitempty"`
	// Masquerade to apply to this message
	Masquerade *DataMessageSendMasquerade `json:"masquerade,omitempty"`
	// Information about how this message should be interacted with
	Interactions *DataMessageSendInteractions `json:"interactions,omitempty"`
}

type DataOnboard struct {
	// New username which will be used to identify the user on the platform
	Username string `json:"username"`
}

type DataPasswordReset struct {
	// Reset token
	Token string `json:"token"`
	// New password
	Password string `json:"password"`
	// Whether to logout all sessions
	RemoveSessions bool `json:"remove_sessions,omitempty"`
}

// Data permissions Value - contains allow
type DataPermissionsValue struct {
	Permissions int64 `json:"permissions"`
}

type DataReportContent struct {
	// Content being reported
	Content *DataReportContentContent `json:"content"`
	// Additional report description
	AdditionalContext string `json:"additional_context,omitempty"`
}

type DataResendVerification struct {
	// Email associated with the account
	Email string `json:"email"`
	// Captcha verification code
	Captcha string `json:"captcha,omitempty"`
}

type DataSendFriendRequest struct {
	// Username and discriminator combo separated by #
	Username string `json:"username"`
}

type DataSendPasswordReset struct {
	// Email associated with the account
	Email string `json:"email"`
	// Captcha verification code
	Captcha string `json:"captcha,omitempty"`
}

type DataSetServerRolePermission struct {
	// Allow / deny values for the role in this server.
	Permissions *DataSetServerRolePermissionPermissions `json:"permissions"`
}

// Embed
type Embed struct {
}

// Representation of an Emoji on Revolt
type Emoji struct {
	// Unique Id
	Id string `json:"_id"`
	// What owns this emoji
	Parent *EmojiParent `json:"parent"`
	// Uploader user id
	CreatorId string `json:"creator_id"`
	// Emoji name
	Name string `json:"name"`
	// Whether the emoji is animated
	Animated bool `json:"animated,omitempty"`
	// Whether the emoji is marked as nsfw
	Nsfw bool `json:"nsfw,omitempty"`
}

// Information about what owns this emoji

// Possible API Errors
type ModelError struct {
}

type Feature struct {
	// Whether the service is enabled
	Enabled bool `json:"enabled"`
	// URL pointing to the service
	Url string `json:"url"`
}

// Bot Response
type FetchBotResponse struct {
	// Bot object
	Bot *FetchBotResponseBot `json:"bot"`
	// User object
	User *FetchBotResponseUser `json:"user"`
}

type FetchServerResponse struct {
    Server
}
// FieldsBot : Optional fields on bot object
type FieldsBot string

// List of FieldsBot
const (
	TOKEN_FieldsBot FieldsBot = "Token"
	INTERACTIONS_URL_FieldsBot FieldsBot = "InteractionsURL"
)
// FieldsChannel : Optional fields on channel object
type FieldsChannel string

// List of FieldsChannel
const (
	DESCRIPTION_FieldsChannel FieldsChannel = "Description"
	ICON_FieldsChannel FieldsChannel = "Icon"
	DEFAULT_PERMISSIONS_FieldsChannel FieldsChannel = "DefaultPermissions"
)
// FieldsMember : Optional fields on server member object
type FieldsMember string

// List of FieldsMember
const (
	NICKNAME_FieldsMember FieldsMember = "Nickname"
	AVATAR_FieldsMember FieldsMember = "Avatar"
	ROLES_FieldsMember FieldsMember = "Roles"
	TIMEOUT_FieldsMember FieldsMember = "Timeout"
)
// FieldsRole : Optional fields on server object
type FieldsRole string

// List of FieldsRole
const (
	COLOUR_FieldsRole FieldsRole = "Colour"
)
// FieldsServer : Optional fields on server object
type FieldsServer string

// List of FieldsServer
const (
	DESCRIPTION_FieldsServer FieldsServer = "Description"
	CATEGORIES_FieldsServer FieldsServer = "Categories"
	SYSTEM_MESSAGES_FieldsServer FieldsServer = "SystemMessages"
	ICON_FieldsServer FieldsServer = "Icon"
	BANNER_FieldsServer FieldsServer = "Banner"
)
// FieldsUser : Optional fields on user object
type FieldsUser string

// List of FieldsUser
const (
	AVATAR_FieldsUser FieldsUser = "Avatar"
	STATUS_TEXT_FieldsUser FieldsUser = "StatusText"
	STATUS_PRESENCE_FieldsUser FieldsUser = "StatusPresence"
	PROFILE_CONTENT_FieldsUser FieldsUser = "ProfileContent"
	PROFILE_BACKGROUND_FieldsUser FieldsUser = "ProfileBackground"
	DISPLAY_NAME_FieldsUser FieldsUser = "DisplayName"
)

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
	Size int64 `json:"size"`
	// Whether this file was deleted
	Deleted bool `json:"deleted,omitempty"`
	// Whether this file was reported
	Reported bool `json:"reported,omitempty"`
	MessageId string `json:"message_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	ServerId string `json:"server_id,omitempty"`
	// Id of the object this file is associated with
	ObjectId string `json:"object_id,omitempty"`
}

type FlagResponse struct {
	// Flags
	Flags int64 `json:"flags"`
}

// Image
type Image struct {
	// URL to the original image
	Url string `json:"url"`
	// Width of the image
	Width int64 `json:"width"`
	// Height of the image
	Height int64 `json:"height"`
	// Positioning and size
	Size *ImageSize `json:"size"`
}
// ImageSize : Image positioning and size
type ImageSize string

// List of ImageSize
const (
	LARGE_ImageSize ImageSize = "Large"
	PREVIEW_ImageSize ImageSize = "Preview"
)

// Collection index
type Index struct {
	// Index name
	Name string `json:"name"`
	// Access information
	Accesses *IndexAccesses `json:"accesses"`
}

// Index access information
type IndexAccess struct {
	// Operations since timestamp
	Ops int64 `json:"ops"`
	// Timestamp at which data keeping begun
	Since *IndexAccessSince `json:"since"`
}

// Information to guide interactions on this message
type Interactions struct {
	// Reactions which should always appear and be distinct
	Reactions []string `json:"reactions,omitempty"`
	// Whether reactions should be restricted to the given list  Can only be set to true if reactions list is of at least length 1
	RestrictReactions bool `json:"restrict_reactions,omitempty"`
}

// Representation of an invite to a channel on Revolt
type Invite struct {
}

type InviteBotDestination struct {
}

type InviteJoinResponse struct {
}

type InviteResponse struct {
}

// Histogram entry
type LatencyHistogramEntry struct {
	// Time
	Micros int64 `json:"micros"`
	// Count
	Count int64 `json:"count"`
}

// Collection latency stats
type LatencyStats struct {
	// Total operations
	Ops int64 `json:"ops"`
	// Timestamp at which data keeping begun
	Latency int64 `json:"latency"`
	// Histogram representation of latency data
	Histogram []LatencyHistogramEntry `json:"histogram"`
}
// LightspeedType : Type of remote Lightspeed.tv content
type LightspeedType string

// List of LightspeedType
const (
	CHANNEL_LightspeedType LightspeedType = "Channel"
)

// Name and / or avatar override information
type Masquerade struct {
	// Replace the display name shown on this message
	Name string `json:"name,omitempty"`
	// Replace the avatar shown on this message (URL to image file)
	Avatar string `json:"avatar,omitempty"`
	// Replace the display role colour shown on this message  Must have `ManageRole` permission to use
	Colour string `json:"colour,omitempty"`
}

// Representation of a member of a server on Revolt
type Member struct {
	// Unique member id
	Id *MemberId `json:"_id"`
	// Time at which this user joined the server
	JoinedAt *MemberJoinedAt `json:"joined_at"`
	// Member's nickname
	Nickname string `json:"nickname,omitempty"`
	// Avatar attachment
	Avatar *File `json:"avatar,omitempty"`
	// Member's roles
	Roles []string `json:"roles,omitempty"`
	// Timestamp this member is timed out until
	Timeout *MemberTimeout `json:"timeout,omitempty"`
}

// Composite primary key consisting of server and user id
type MemberCompositeKey struct {
	// Server Id
	Server string `json:"server"`
	// User Id
	User string `json:"user"`
}

type MemberQueryResponse struct {
	// List of members
	Members []Member `json:"members"`
	// List of users
	Users []User `json:"users"`
}

// Representation of a Message on Revolt
type Message struct {
	// Unique Id
	Id string `json:"_id"`
	// Unique value generated by client sending this message
	Nonce string `json:"nonce,omitempty"`
	// Id of the channel this message was sent in
	Channel string `json:"channel"`
	// Id of the user or webhook that sent this message
	Author string `json:"author"`
	// The webhook that sent this message
	Webhook *MessageWebhook `json:"webhook,omitempty"`
	// Message content
	Content string `json:"content,omitempty"`
	// System message
	System *MessageSystem `json:"system,omitempty"`
	// Array of attachments
	Attachments []*File `json:"attachments,omitempty"`
	// Time at which this message was last edited
	Edited *MessageEdited `json:"edited,omitempty"`
	// Attached embeds to this message
	Embeds []Embed `json:"embeds,omitempty"`
	// Array of user ids mentioned in this message
	Mentions []string `json:"mentions,omitempty"`
	// Array of message ids this message is replying to
	Replies []string `json:"replies,omitempty"`
	// Hashmap of emoji IDs to array of user IDs
	Reactions map[string][]string `json:"reactions,omitempty"`
	// Information about how this message should be interacted with
	Interactions *MessageInteractions `json:"interactions,omitempty"`
	// Name and / or avatar overrides for this message
	Masquerade *MessageMasquerade `json:"masquerade,omitempty"`
}

// Filter and sort messages by time
type MessageQuery struct {
	// Maximum number of messages to fetch  For fetching nearby messages, this is \\`(limit + 1)\\`.
	Limit int64 `json:"limit,omitempty"`
	// Parent channel ID
	Channel string `json:"channel,omitempty"`
	// Message author ID
	Author string `json:"author,omitempty"`
	// Search query
	Query string `json:"query,omitempty"`
}
// MessageSort : Sort used for retrieving messages
type MessageSort string

// List of MessageSort
const (
	RELEVANCE_MessageSort MessageSort = "Relevance"
	LATEST_MessageSort MessageSort = "Latest"
	OLDEST_MessageSort MessageSort = "Oldest"
)

// Information about the webhook bundled with Message
type MessageWebhook struct {
	Name string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}

// Metadata associated with file
type Metadata struct {
}
// MfaMethod : MFA method
type MfaMethod string

// List of MFAMethod
const (
	PASSWORD_MfaMethod MfaMethod = "Password"
	RECOVERY_MfaMethod MfaMethod = "Recovery"
	TOTP_MfaMethod MfaMethod = "Totp"
)

// MFA response
type MfaResponse struct {
}

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
	EmailOtp bool `json:"email_otp"`
	TrustedHandover bool `json:"trusted_handover"`
	EmailMfa bool `json:"email_mfa"`
	TotpMfa bool `json:"totp_mfa"`
	SecurityKeyMfa bool `json:"security_key_mfa"`
	RecoveryActive bool `json:"recovery_active"`
}

type MutualResponse struct {
	// Array of mutual user IDs that both users are friends with
	Users []string `json:"users"`
	// Array of mutual server IDs that both users are in
	Servers []string `json:"servers"`
}

type NewRoleResponse struct {
	// Id of the role
	Id string `json:"id"`
	// New role
	Role *NewRoleResponseRole `json:"role"`
}

type OneOfReport struct {
}

type OptionsBulkDelete struct {
	// Message IDs
	Ids []string `json:"ids"`
}

type OptionsFetchSettings struct {
	// Keys to fetch
	Keys []string `json:"keys"`
}

type OptionsMessageSearch struct {
	// Full-text search query  See [MongoDB documentation](https://docs.mongodb.com/manual/text-search/#-text-operator) for more information.
	Query string `json:"query"`
	// Maximum number of messages to fetch
	Limit int64 `json:"limit,omitempty"`
	// Message id before which messages should be fetched
	Before string `json:"before,omitempty"`
	// Message id after which messages should be fetched
	After string `json:"after,omitempty"`
	// Message sort direction  By default, it will be sorted by latest.
	Sort *OptionsMessageSearchSort `json:"sort,omitempty"`
	// Whether to include user (and member, if server channel) objects
	IncludeUsers bool `json:"include_users,omitempty"`
}

type OptionsQueryStale struct {
	// Array of message IDs
	Ids []string `json:"ids"`
}

// Representation of a single permission override
type Override struct {
	// Allow bit flags
	Allow int64 `json:"allow"`
	// Disallow bit flags
	Deny int64 `json:"deny"`
}

// Representation of a single permission override as it appears on models and in the database
type OverrideField struct {
	// Allow bit flags
	A int64 `json:"a"`
	// Disallow bit flags
	D int64 `json:"d"`
}

// Both lists are sorted by their IDs.
type OwnedBotsResponse struct {
	// Bot objects
	Bots []Bot `json:"bots"`
	// User objects
	Users []User `json:"users"`
}
// Permission : Permission value on Revolt  This should be restricted to the lower 52 bits to prevent any potential issues with Javascript. Also leave empty spaces for future permission flags to be added.
type PermissionFriendly string

// List of Permission
const (
	MANAGE_CHANNEL_PermissionFriendly PermissionFriendly = "ManageChannel"
	MANAGE_SERVER_PermissionFriendly PermissionFriendly = "ManageServer"
	MANAGE_PERMISSIONS_PermissionFriendly PermissionFriendly = "ManagePermissions"
	MANAGE_ROLE_PermissionFriendly PermissionFriendly = "ManageRole"
	MANAGE_CUSTOMISATION_PermissionFriendly PermissionFriendly = "ManageCustomisation"
	KICK_MEMBERS_PermissionFriendly PermissionFriendly = "KickMembers"
	BAN_MEMBERS_PermissionFriendly PermissionFriendly = "BanMembers"
	TIMEOUT_MEMBERS_PermissionFriendly PermissionFriendly = "TimeoutMembers"
	ASSIGN_ROLES_PermissionFriendly PermissionFriendly = "AssignRoles"
	CHANGE_NICKNAME_PermissionFriendly PermissionFriendly = "ChangeNickname"
	MANAGE_NICKNAMES_PermissionFriendly PermissionFriendly = "ManageNicknames"
	CHANGE_AVATAR_PermissionFriendly PermissionFriendly = "ChangeAvatar"
	REMOVE_AVATARS_PermissionFriendly PermissionFriendly = "RemoveAvatars"
	VIEW_CHANNEL_PermissionFriendly PermissionFriendly = "ViewChannel"
	READ_MESSAGE_HISTORY_PermissionFriendly PermissionFriendly = "ReadMessageHistory"
	SEND_MESSAGE_PermissionFriendly PermissionFriendly = "SendMessage"
	MANAGE_MESSAGES_PermissionFriendly PermissionFriendly = "ManageMessages"
	MANAGE_WEBHOOKS_PermissionFriendly PermissionFriendly = "ManageWebhooks"
	INVITE_OTHERS_PermissionFriendly PermissionFriendly = "InviteOthers"
	SEND_EMBEDS_PermissionFriendly PermissionFriendly = "SendEmbeds"
	UPLOAD_FILES_PermissionFriendly PermissionFriendly = "UploadFiles"
	MASQUERADE_PermissionFriendly PermissionFriendly = "Masquerade"
	REACT_PermissionFriendly PermissionFriendly = "React"
	CONNECT_PermissionFriendly PermissionFriendly = "Connect"
	SPEAK_PermissionFriendly PermissionFriendly = "Speak"
	VIDEO_PermissionFriendly PermissionFriendly = "Video"
	MUTE_MEMBERS_PermissionFriendly PermissionFriendly = "MuteMembers"
	DEAFEN_MEMBERS_PermissionFriendly PermissionFriendly = "DeafenMembers"
	MOVE_MEMBERS_PermissionFriendly PermissionFriendly = "MoveMembers"
	GRANT_ALL_SAFE_PermissionFriendly PermissionFriendly = "GrantAllSafe"
	GRANT_ALL_PermissionFriendly PermissionFriendly = "GrantAll"
)
// Presence : Presence status
type Presence string

// List of Presence
const (
	ONLINE_Presence Presence = "Online"
	IDLE_Presence Presence = "Idle"
	FOCUS_Presence Presence = "Focus"
	BUSY_Presence Presence = "Busy"
	INVISIBLE_Presence Presence = "Invisible"
)

// Public Bot
type PublicBot struct {
	// Bot Id
	Id string `json:"_id"`
	// Bot Username
	Username string `json:"username"`
	// Profile Avatar
	Avatar string `json:"avatar"`
	// Profile Description
	Description string `json:"description"`
}

// Collection query execution stats
type QueryExecStats struct {
	// Stats regarding collection scans
	CollectionScans *QueryExecStatsCollectionScans `json:"collectionScans"`
}

// Relationship entry indicating current status with other user
type Relationship struct {
	Id string `json:"_id"`
	Status *RelationshipStatus `json:"status"`
}
// RelationshipStatus : User's relationship with another user (or themselves)
type RelationshipStatus string

// List of RelationshipStatus
const (
	NONE_RelationshipStatus RelationshipStatus = "None"
	USER_RelationshipStatus RelationshipStatus = "User"
	FRIEND_RelationshipStatus RelationshipStatus = "Friend"
	OUTGOING_RelationshipStatus RelationshipStatus = "Outgoing"
	INCOMING_RelationshipStatus RelationshipStatus = "Incoming"
	BLOCKED_RelationshipStatus RelationshipStatus = "Blocked"
	BLOCKED_OTHER_RelationshipStatus RelationshipStatus = "BlockedOther"
)

// Representation of a message reply before it is sent.
type Reply struct {
	// Message Id
	Id string `json:"id"`
	// Whether this reply should mention the message's author
	Mention bool `json:"mention"`
}

// User-generated platform moderation report.
type Report struct {
	// Unique Id
	Id string `json:"_id"`
	// Id of the user creating this report
	AuthorId string `json:"author_id"`
	// Reported content
	Content *ReportContent `json:"content"`
	// Additional report context
	AdditionalContext string `json:"additional_context"`
	// Additional notes included on the report
	Notes string `json:"notes,omitempty"`
}

// Status of the report
type ReportStatus struct {
}
// ReportStatusString : Just the status of the report
type ReportStatusString string

// List of ReportStatusString
const (
	CREATED_ReportStatusString ReportStatusString = "Created"
	REJECTED_ReportStatusString ReportStatusString = "Rejected"
	RESOLVED_ReportStatusString ReportStatusString = "Resolved"
)

// The content being reported
type ReportedContent struct {
}

type ResponseLogin struct {
}

type ResponseTotpSecret struct {
	Secret string `json:"secret"`
}

type ResponseVerify struct {
}

type RevoltConfig struct {
	// Revolt API Version
	Revolt string `json:"revolt"`
	// Features enabled on this Revolt node
	Features *RevoltConfigFeatures `json:"features"`
	// WebSocket URL
	Ws string `json:"ws"`
	// URL pointing to the client serving this node
	App string `json:"app"`
	// Web Push VAPID public key
	Vapid string `json:"vapid"`
	// Build information
	Build *RevoltConfigBuild `json:"build"`
}

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

// Representation of a server role
type Role struct {
	// Role name
	Name string `json:"name"`
	// Permissions available to this role
	Permissions *RolePermissions `json:"permissions"`
	// Colour used for this role  This can be any valid CSS colour
	Colour string `json:"colour,omitempty"`
	// Whether this role should be shown separately on the member sidebar
	Hoist bool `json:"hoist,omitempty"`
	// Ranking of this role
	Rank int64 `json:"rank,omitempty"`
}

// Representation of a text embed before it is sent.
type SendableEmbed struct {
	IconUrl string `json:"icon_url,omitempty"`
	Url string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Media string `json:"media,omitempty"`
	Colour string `json:"colour,omitempty"`
}

// Representation of a server on Revolt
type Server struct {
	// Unique Id
	Id string `json:"_id"`
	// User id of the owner
	Owner string `json:"owner"`
	// Name of the server
	Name string `json:"name"`
	// Description for the server
	Description string `json:"description,omitempty"`
	// Channels within this server
	Channels []string `json:"channels"`
	// Categories for this server
	Categories []Category `json:"categories,omitempty"`
	// Configuration for sending system event messages
	SystemMessages *ServerSystemMessages `json:"system_messages,omitempty"`
	// Roles for this server
	Roles map[string]Role `json:"roles,omitempty"`
	// Default set of server and channel permissions
	DefaultPermissions int64 `json:"default_permissions"`
	// Icon attachment
	Icon *File `json:"icon,omitempty"`
	// Banner attachment
	Banner *File `json:"banner,omitempty"`
	// Bitfield of server flags
	Flags int64 `json:"flags,omitempty"`
	// Whether this server is flagged as not safe for work
	Nsfw bool `json:"nsfw,omitempty"`
	// Whether to enable analytics
	Analytics bool `json:"analytics,omitempty"`
	// Whether this server should be publicly discoverable
	Discoverable bool `json:"discoverable,omitempty"`
}

// Representation of a server ban on Revolt
type ServerBan struct {
	// Unique member id
	Id *ServerBanId `json:"_id"`
	// Reason for ban creation
	Reason string `json:"reason,omitempty"`
}

type SessionInfo struct {
	Id string `json:"_id"`
	Name string `json:"name"`
}

// Enum to map into different models that can be saved in a snapshot

// Snapshot of some content with required data to render
type SnapshotWithContext struct {
	// Users involved in snapshot
	Users []User `json:"_users"`
	// Channels involved in snapshot
	Channels []Channel `json:"_channels"`
	// Server involved in snapshot
	Server *SnapshotWithContextServer `json:"_server,omitempty"`
	// Unique Id
	Id string `json:"_id"`
	// Report parent Id
	ReportId string `json:"report_id"`
	// Snapshot of content
	Content *Object `json:"content"`
}

// Information about special remote content
type Special struct {
}

// Server Stats
type Stats struct {
	// Index usage information
	Indices map[string][]Index `json:"indices"`
	// Collection stats
	CollStats map[string]CollectionStats `json:"coll_stats"`
}

// Representation of a system event message
type SystemMessage struct {
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
// TwitchType : Type of remote Twitch content
type TwitchType string

// List of TwitchType
const (
	CHANNEL_TwitchType TwitchType = "Channel"
	VIDEO_TwitchType TwitchType = "Video"
	CLIP_TwitchType TwitchType = "Clip"
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
	Badges int64 `json:"badges,omitempty"`
	// User's current status
	Status *UserStatus `json:"status,omitempty"`
	// User's profile page
	Profile *UserProfile `json:"profile,omitempty"`
	// Enum of user flags
	Flags int64 `json:"flags,omitempty"`
	// Whether this user is privileged
	Privileged bool `json:"privileged,omitempty"`
	// Bot information
	Bot *UserBot `json:"bot,omitempty"`
	// Current session user's relationship with this user
	Relationship string `json:"relationship,omitempty"`
	// Whether this user is currently online
	Online bool `json:"online,omitempty"`
}
// UserPermission : User permission definitions
type UserPermission string

// List of UserPermission
const (
	ACCESS_UserPermission UserPermission = "Access"
	VIEW_PROFILE_UserPermission UserPermission = "ViewProfile"
	SEND_MESSAGE_UserPermission UserPermission = "SendMessage"
	INVITE_UserPermission UserPermission = "Invite"
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
// UserReportReason : Reason for reporting a user
type UserReportReason string

// List of UserReportReason
const (
	NONE_SPECIFIED_UserReportReason UserReportReason = "NoneSpecified"
	UNSOLICITED_SPAM_UserReportReason UserReportReason = "UnsolicitedSpam"
	SPAM_ABUSE_UserReportReason UserReportReason = "SpamAbuse"
	INAPPROPRIATE_PROFILE_UserReportReason UserReportReason = "InappropriateProfile"
	IMPERSONATION_UserReportReason UserReportReason = "Impersonation"
	BAN_EVASION_UserReportReason UserReportReason = "BanEvasion"
	UNDERAGE_UserReportReason UserReportReason = "Underage"
)

// User's active status
type UserStatus struct {
	// Custom status text
	Text string `json:"text,omitempty"`
	// Current presence option
	Presence Presence `json:"presence,omitempty"`
}

// Video
type Video struct {
	// URL to the original video
	Url string `json:"url"`
	// Width of the video
	Width int64 `json:"width"`
	// Height of the video
	Height int64 `json:"height"`
}

type VoiceFeature struct {
	// Whether voice is enabled
	Enabled bool `json:"enabled"`
	// URL pointing to the voice API
	Url string `json:"url"`
	// URL pointing to the voice WebSocket server
	Ws string `json:"ws"`
}

// Web Push subscription
type WebPushSubscription struct {
	Endpoint string `json:"endpoint"`
	P256dh string `json:"p256dh"`
	Auth string `json:"auth"`
}

// Webhook
type Webhook struct {
	// Webhook Id
	Id string `json:"id"`
	// The name of the webhook
	Name string `json:"name"`
	// The avatar of the webhook
	Avatar *File `json:"avatar,omitempty"`
	// The channel this webhook belongs to
	ChannelId string `json:"channel_id"`
	// The private token for the webhook
	Token string `json:"token,omitempty"`
}
