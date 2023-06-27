package types

import "github.com/infinitybotlist/grevolt/types/timestamp"

// FieldsMember : Optional fields on server member object
type FieldsMember string

// List of FieldsMember
const (
	NICKNAME_FieldsMember FieldsMember = "Nickname"
	AVATAR_FieldsMember   FieldsMember = "Avatar"
	ROLES_FieldsMember    FieldsMember = "Roles"
	TIMEOUT_FieldsMember  FieldsMember = "Timeout"
)

// Both lists are sorted by ID.
type AllMemberResponse struct {
	// List of members
	Members []Member `json:"members"`
	// List of users
	Users []User `json:"users"`
}

// Member object
type DataMemberEdit struct {
	// Member nickname
	Nickname string `json:"nickname,omitempty"`
	// Attachment Id to set for avatar
	Avatar string `json:"avatar,omitempty"`
	// Array of role ids
	Roles []string `json:"roles,omitempty"`
	// Timestamp this member is timed out until
	Timeout timestamp.Timestamp `json:"timeout,omitempty"`
	// Fields to remove from channel object
	Remove []FieldsMember `json:"remove,omitempty"`
}

// Unique member id
//
// <this is a composite primary key consisting of server and user id>
type MemberId struct {
	// Server Id
	Server string `json:"server"`
	// User Id
	User string `json:"user"`
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

// Member Query Response
//
// TODO: Provide better docs on what this is
type MemberQueryResponse struct {
	// List of members
	Members []Member `json:"members"`
	// List of users
	Users []User `json:"users"`
}

// Data needed to ban a user, note that all fields are optional
type DataBanCreate struct {
	// Ban reason
	Reason string `json:"reason,omitempty"`
}

// Result from querying list of bans
type BanListResult struct {
	// Users objects
	Users []BannedUser `json:"users"`
	// Ban objects
	Bans []ServerBan `json:"bans"`
}

// Representation of a server ban on Revolt
type ServerBan struct {
	// Unique member id
	Id *ServerBanId `json:"_id"`
	// Reason for ban creation
	Reason string `json:"reason,omitempty"`
}

// Unique member id
type ServerBanId struct {
	// Server Id
	Server string `json:"server"`
	// User Id
	User string `json:"user"`
}

// Just enoguh user information to list bans.
type BannedUser struct {
	// Id of the banned user
	Id string `json:"_id"`
	// Username of the banned user
	Username string `json:"username"`
	// Avatar of the banned user
	Avatar *File `json:"avatar,omitempty"`
}
