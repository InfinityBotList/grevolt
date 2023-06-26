package types

import (
	"time"

	"github.com/infinitybotlist/grevolt/types/timestamp"
)

type Object map[string]any

// Emoji parent, not properly generated
type EmojiParent struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// FieldsWebhook : Optional fields on webhook object
//
// <undocumented, from https://github.com/revoltchat/backend/blob/master/crates/core/database/src/models/channel_webhooks/ops/mongodb.rs#L71>
type FieldsWebhook string

const (
	AVATAR_FieldsWebhook FieldsWebhook = "Avatar"
)

// InviteType : The type of an invite
type InviteType string

const (
	GROUP_InviteType  InviteType = "Group"
	SERVER_InviteType InviteType = "Server"
)

// Representation of an created invite on Revolt
type CreateInviteResponseInvite struct {
	// The Id of the invite
	Id string `json:"_id,omitempty"`

	// The type of the invite
	Type InviteType `json:"type,omitempty"`

	// The creator of the invite
	Creator string `json:"creator,omitempty"`

	// The channel the invite is for
	Channel string `json:"channel,omitempty"`

	// The server the invite points to if it is a server invite
	Server string `json:"server,omitempty"`
}

type AccountInfo struct {
	Id    string `json:"_id"`
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

// Query exec stats
type CollectionStatsQueryExecStats struct {
	// Stats regarding collection scans
	CollectionScans *Object `json:"collectionScans"`
}

// Parent information
type DataCreateEmojiParent struct {
}

// New report status
type DataEditReportStatus struct {
}

// Timestamp this member is timed out until
type DataMemberEditTimeout struct {
}

// Content being reported
type DataReportContentContent struct {
}

// Positioning and size

// Timestamp at which data keeping begun
type IndexAccessSince struct {
}

// Access information
type IndexAccesses struct {
	// Operations since timestamp
	Ops uint64 `json:"ops"`
	// Timestamp at which data keeping begun
	Since timestamp.Timestamp `json:"since"`
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

// Timestamp this member is timed out until
type MemberTimeout struct {
}

// Stats regarding collection scans
type QueryExecStatsCollectionScans struct {
	// Number of total collection scans
	Total uint64 `json:"total"`
	// Number of total collection scans not using a tailable cursor
	NonTailable uint64 `json:"nonTailable"`
}

// Permissions available to this role
type RolePermissions struct {
	// Allow bit flags
	A uint64 `json:"a"`
	// Disallow bit flags
	D uint64 `json:"d"`
}

// Unique member id
type ServerBanId struct {
	// Server Id
	Server string `json:"server"`
	// User Id
	User string `json:"user"`
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

type CaptchaFeature struct {
	// Whether captcha is enabled
	Enabled bool `json:"enabled"`
	// Client key used for solving captcha
	Key string `json:"key"`
}

// Representation of a channel on Revolt

// Composite primary key consisting of channel and user id

// Query collection scan stats
type CollectionScans struct {
	// Number of total collection scans
	Total uint64 `json:"total"`
	// Number of total collection scans not using a tailable cursor
	NonTailable uint64 `json:"nonTailable"`
}

// Collection stats
type CollectionStats struct {
	// Namespace
	Ns string `json:"ns"`
	// Local time
	LocalTime timestamp.Timestamp `json:"localTime"`
	// Latency stats
	LatencyStats map[string]LatencyStats `json:"latencyStats"`
	// Query exec stats
	QueryExecStats *CollectionStatsQueryExecStats `json:"queryExecStats"`
	// Number of documents in collection
	Count uint64 `json:"count"`
}

type CreateVoiceUserResponse struct {
	// Token for authenticating with the voice server
	Token string `json:"token"`
}

type CreateWebhookBody struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
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
	Rank uint64 `json:"rank,omitempty"`
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

type DataEditRole struct {
	// Role name
	Name string `json:"name,omitempty"`
	// Role colour
	Colour string `json:"colour,omitempty"`
	// Whether this role should be displayed separately
	Hoist bool `json:"hoist,omitempty"`
	// Ranking position  Smaller values take priority.
	Rank uint64 `json:"rank,omitempty"`
	// Fields to remove from role object
	Remove []FieldsRole `json:"remove,omitempty"`
}

type DataEditSession struct {
	// Session friendly name
	FriendlyName string `json:"friendly_name"`
}

type DataHello struct {
	// Whether onboarding is required
	Onboarding bool `json:"onboarding"`
}

type DataMemberEdit struct {
	// Member nickname
	Nickname string `json:"nickname,omitempty"`
	// Attachment Id to set for avatar
	Avatar string `json:"avatar,omitempty"`
	// Array of role ids
	Roles []string `json:"roles,omitempty"`
	// Timestamp this member is timed out until
	Timeout time.Time `json:"timeout,omitempty"`
	// Fields to remove from channel object
	Remove []FieldsMember `json:"remove,omitempty"`
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
	Permissions uint64 `json:"permissions"`
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

type FetchServerResponse struct {
	Server
}

// FieldsMember : Optional fields on server member object
type FieldsMember string

// List of FieldsMember
const (
	NICKNAME_FieldsMember FieldsMember = "Nickname"
	AVATAR_FieldsMember   FieldsMember = "Avatar"
	ROLES_FieldsMember    FieldsMember = "Roles"
	TIMEOUT_FieldsMember  FieldsMember = "Timeout"
)

// FieldsRole : Optional fields on server object
type FieldsRole string

// List of FieldsRole
const (
	COLOUR_FieldsRole FieldsRole = "Colour"
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
	Ops uint64 `json:"ops"`
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

// Histogram entry
type LatencyHistogramEntry struct {
	// Time
	Micros uint64 `json:"micros"`
	// Count
	Count uint64 `json:"count"`
}

// Collection latency stats
type LatencyStats struct {
	// Total operations
	Ops uint64 `json:"ops"`
	// Timestamp at which data keeping begun
	Latency uint64 `json:"latency"`
	// Histogram representation of latency data
	Histogram []LatencyHistogramEntry `json:"histogram"`
}

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
	JoinedAt timestamp.Timestamp `json:"joined_at,omitempty"`
	// Member's nickname
	Nickname string `json:"nickname,omitempty"`
	// Avatar attachment
	Avatar *File `json:"avatar,omitempty"`
	// Member's roles
	Roles []string `json:"roles,omitempty"`
	// Timestamp this member is timed out until
	Timeout timestamp.Timestamp `json:"timeout,omitempty"`
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

// MfaMethod : MFA method
type MfaMethod string

// List of MFAMethod
const (
	PASSWORD_MfaMethod MfaMethod = "Password"
	RECOVERY_MfaMethod MfaMethod = "Recovery"
	TOTP_MfaMethod     MfaMethod = "Totp"
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
	EmailOtp        bool `json:"email_otp"`
	TrustedHandover bool `json:"trusted_handover"`
	EmailMfa        bool `json:"email_mfa"`
	TotpMfa         bool `json:"totp_mfa"`
	SecurityKeyMfa  bool `json:"security_key_mfa"`
	RecoveryActive  bool `json:"recovery_active"`
}

type NewRoleResponse struct {
	// Id of the role
	Id string `json:"id"`
	// New role
	Role *Role `json:"role"`
}

type OptionsFetchSettings struct {
	// Keys to fetch
	Keys []string `json:"keys"`
}

type OptionsQueryStale struct {
	// Array of message IDs
	Ids []string `json:"ids"`
}

// Collection query execution stats
type QueryExecStats struct {
	// Stats regarding collection scans
	CollectionScans *QueryExecStatsCollectionScans `json:"collectionScans"`
}

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
	Rank uint64 `json:"rank,omitempty"`
}

// Representation of a server ban on Revolt
type ServerBan struct {
	// Unique member id
	Id *ServerBanId `json:"_id"`
	// Reason for ban creation
	Reason string `json:"reason,omitempty"`
}

type SessionInfo struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

// Enum to map into different models that can be saved in a snapshot

// Server Stats
type Stats struct {
	// Index usage information
	Indices map[string][]Index `json:"indices"`
	// Collection stats
	CollStats map[string]CollectionStats `json:"coll_stats"`
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
	P256dh   string `json:"p256dh"`
	Auth     string `json:"auth"`
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
