// Defines permissions flags, taken from https://github.com/revoltchat/backend/blob/master/crates/quark/src/permissions/defn/permission.rs
package permissions

type Permission uint64

const (
	// * Generic permissions
	/// Manage the channel or channels on the server
	ManageChannel Permission = 1 << 0
	/// Manage the server
	ManageServer Permission = 1 << 1
	/// Manage permissions on servers or channels
	ManagePermissions Permission = 1 << 2
	/// Manage roles on server
	ManageRole Permission = 1 << 3
	/// Manage server customisation (includes emoji)
	ManageCustomisation Permission = 1 << 4

	// % 1 bit reserved

	// * Member permissions
	/// Kick other members below their ranking
	KickMembers Permission = 1 << 6
	/// Ban other members below their ranking
	BanMembers Permission = 1 << 7
	/// Timeout other members below their ranking
	TimeoutMembers Permission = 1 << 8
	/// Assign roles to members below their ranking
	AssignRoles Permission = 1 << 9
	/// Change own nickname
	ChangeNickname Permission = 1 << 10
	/// Change or remove other's nicknames below their ranking
	ManageNicknames Permission = 1 << 11
	/// Change own avatar
	ChangeAvatar Permission = 1 << 12
	/// Remove other's avatars below their ranking
	RemoveAvatars Permission = 1 << 13

	// % 7 bits reserved

	// * Channel permissions
	/// View a channel
	ViewChannel Permission = 1 << 20
	/// Read a channel's past message history
	ReadMessageHistory Permission = 1 << 21
	/// Send a message in a channel
	SendMessage Permission = 1 << 22
	/// Delete messages in a channel
	ManageMessages Permission = 1 << 23
	/// Manage webhook entries on a channel
	ManageWebhooks Permission = 1 << 24
	/// Create invites to this channel
	InviteOthers Permission = 1 << 25
	/// Send embedded content in this channel
	SendEmbeds Permission = 1 << 26
	/// Send attachments and media in this channel
	UploadFiles Permission = 1 << 27
	/// Masquerade messages using custom nickname and avatar
	Masquerade Permission = 1 << 28
	/// React to messages with emojis
	React Permission = 1 << 29

	// * Voice permissions
	/// Connect to a voice channel
	Connect Permission = 1 << 30
	/// Speak in a voice call
	Speak Permission = 1 << 31
	/// Share video in a voice call
	Video Permission = 1 << 32
	/// Mute other members with lower ranking in a voice call
	MuteMembers Permission = 1 << 33
	/// Deafen other members with lower ranking in a voice call
	DeafenMembers Permission = 1 << 34
	/// Move members between voice channels
	MoveMembers Permission = 1 << 35

	// * Misc. permissions
	// % Bits 36 to 52: free area
	// % Bits 53 to 64: do not use

	// * Grant all permissions
	/// Safely grant all permissions
	GrantAllSafe Permission = 0x000F_FFFF_FFFF_FFFF

	/// Grant all permissions
	GrantAll Permission = 18446744073709551615
)
