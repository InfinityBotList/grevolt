package friendly

// Permission : Permission value on Revolt  This should be restricted to the lower 52 bits to prevent any potential issues with Javascript. Also leave empty spaces for future permission flags to be added.
//
// <note that you probably want the bitflags version of this available under flags/permissions>
type PermissionFriendly string

// List of Permission
const (
	MANAGE_CHANNEL_PermissionFriendly       PermissionFriendly = "ManageChannel"
	MANAGE_SERVER_PermissionFriendly        PermissionFriendly = "ManageServer"
	MANAGE_PERMISSIONS_PermissionFriendly   PermissionFriendly = "ManagePermissions"
	MANAGE_ROLE_PermissionFriendly          PermissionFriendly = "ManageRole"
	MANAGE_CUSTOMISATION_PermissionFriendly PermissionFriendly = "ManageCustomisation"
	KICK_MEMBERS_PermissionFriendly         PermissionFriendly = "KickMembers"
	BAN_MEMBERS_PermissionFriendly          PermissionFriendly = "BanMembers"
	TIMEOUT_MEMBERS_PermissionFriendly      PermissionFriendly = "TimeoutMembers"
	ASSIGN_ROLES_PermissionFriendly         PermissionFriendly = "AssignRoles"
	CHANGE_NICKNAME_PermissionFriendly      PermissionFriendly = "ChangeNickname"
	MANAGE_NICKNAMES_PermissionFriendly     PermissionFriendly = "ManageNicknames"
	CHANGE_AVATAR_PermissionFriendly        PermissionFriendly = "ChangeAvatar"
	REMOVE_AVATARS_PermissionFriendly       PermissionFriendly = "RemoveAvatars"
	VIEW_CHANNEL_PermissionFriendly         PermissionFriendly = "ViewChannel"
	READ_MESSAGE_HISTORY_PermissionFriendly PermissionFriendly = "ReadMessageHistory"
	SEND_MESSAGE_PermissionFriendly         PermissionFriendly = "SendMessage"
	MANAGE_MESSAGES_PermissionFriendly      PermissionFriendly = "ManageMessages"
	MANAGE_WEBHOOKS_PermissionFriendly      PermissionFriendly = "ManageWebhooks"
	INVITE_OTHERS_PermissionFriendly        PermissionFriendly = "InviteOthers"
	SEND_EMBEDS_PermissionFriendly          PermissionFriendly = "SendEmbeds"
	UPLOAD_FILES_PermissionFriendly         PermissionFriendly = "UploadFiles"
	MASQUERADE_PermissionFriendly           PermissionFriendly = "Masquerade"
	REACT_PermissionFriendly                PermissionFriendly = "React"
	CONNECT_PermissionFriendly              PermissionFriendly = "Connect"
	SPEAK_PermissionFriendly                PermissionFriendly = "Speak"
	VIDEO_PermissionFriendly                PermissionFriendly = "Video"
	MUTE_MEMBERS_PermissionFriendly         PermissionFriendly = "MuteMembers"
	DEAFEN_MEMBERS_PermissionFriendly       PermissionFriendly = "DeafenMembers"
	MOVE_MEMBERS_PermissionFriendly         PermissionFriendly = "MoveMembers"
	GRANT_ALL_SAFE_PermissionFriendly       PermissionFriendly = "GrantAllSafe"
	GRANT_ALL_PermissionFriendly            PermissionFriendly = "GrantAll"
)
