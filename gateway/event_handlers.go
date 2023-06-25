package gateway

import "github.com/infinitybotlist/grevolt/types/events"

// Event handler for the websocket
//
// See https://github.com/revoltchat/backend/tree/master/crates/quark/src/events for event list
type EventHandlers struct {
	// An error occurred which meant you couldn't authenticate.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Error Event[events.Error]

	// The server has authenticated your connection and you will shortly start receiving data.
	Authenticated Event[events.Authenticated]

	// Several events have been sent, process each item of v as its own event.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Bulk Event[events.Bulk]

	// Ping response from the server.
	//
	// <Note that grevolt handles these for you in general, but you can provide additional logic here>
	Pong Event[events.Pong]

	// Data for use by client, data structures match the API specification
	Ready Event[events.Ready]

	// Message received, the event object has the same schema as the Message object in the API with the addition of an event type.
	Message Event[events.Message]

	// Message edited or otherwise updated.
	MessageUpdate Event[events.MessageUpdate]

	// Message has data being appended to it.
	MessageAppend Event[events.MessageAppend]

	// Message has been deleted.
	MessageDelete Event[events.MessageDelete]

	// A reaction has been added to a message.
	MessageReact Event[events.MessageReact]

	// A reaction has been removed from a message.
	MessageUnreact Event[events.MessageUnreact]

	// A certain reaction has been removed from the message.
	//
	// <the difference between this and MessageUnreact is that
	// this event is sent when a user with manage messages removes
	// a reaction while MessageUnreact is sent when a user removes
	// their own reaction>
	MessageRemoveReaction Event[events.MessageRemoveReaction]

	// Channel created, the event object has the same schema as the Channel object in the API with the addition of an event type.
	ChannelCreate Event[events.ChannelCreate]

	// Channel details updated.
	ChannelUpdate Event[events.ChannelUpdate]

	// Channel has been deleted.
	ChannelDelete Event[events.ChannelDelete]

	// A user has joined the group.
	ChannelGroupJoin Event[events.ChannelGroupJoin]

	// A user has left the group.
	ChannelGroupLeave Event[events.ChannelGroupLeave]

	// A user has started typing in this channel.
	ChannelStartTyping Event[events.ChannelStartTyping]

	// A user has stopped typing in this channel.
	ChannelStopTyping Event[events.ChannelStopTyping]

	// You have acknowledged new messages in this channel up to this message ID.
	//
	// <official docs say the above, but it should be 'A user' instead of 'you'?>
	ChannelAck Event[events.ChannelAck]

	// Server created, the event object has the same schema as the SERVER object in the API with the addition of an event type.
	ServerCreate Event[events.ServerCreate]

	// Server details updated.
	ServerUpdate Event[events.ServerUpdate]

	// Server has been deleted.
	ServerDelete Event[events.ServerDelete]

	// Server member details updated.
	ServerMemberUpdate Event[events.ServerMemberUpdate]

	// A user has joined the group.
	//
	// <this should be server, not group>
	ServerMemberJoin Event[events.ServerMemberJoin]

	// A user has left the group.
	//
	// <this should be server, not group>
	ServerMemberLeave Event[events.ServerMemberLeave]

	// Server role has been updated or created.
	ServerRoleUpdate Event[events.ServerRoleUpdate]

	// Server role has been deleted.
	ServerRoleDelete Event[events.ServerRoleDelete]

	// User has been updated.
	UserUpdate Event[events.UserUpdate]

	// Your relationship with another user has changed.
	UserRelationship Event[events.UserRelationship]

	// Settings updated remotely
	//
	// <undocumented, will likely be available in a future release>
	UserSettingsUpdate Event[events.UserSettingsUpdate]

	// User has been platform banned or deleted their account
	//
	// Clients should remove the following associated data:
	//   - Messages
	//   - DM Channels
	//   - Relationships
	//   - Server Memberships
	//
	// User flags are specified to explain why a wipe is occurring though not all reasons will necessarily ever appear.
	UserPlatformWipe Event[events.UserPlatformWipe]

	// Emoji created, the event object has the same schema as the Emoji object in the API with the addition of an event type.
	EmojiCreate Event[events.EmojiCreate]

	// Emoji has been deleted.
	EmojiDelete Event[events.EmojiDelete]

	// New webhook
	//
	// <undocumented, will likely be available in a future release>
	WebhookCreate Event[events.WebhookCreate]

	// Update existing webhook
	//
	// <undocumented, will likely be available in a future release>
	WebhookUpdate Event[events.WebhookUpdate]

	// Delete existing webhook
	//
	// <undocumented, will likely be available in a future release>
	WebhookDelete Event[events.WebhookDelete]

	// New report
	//
	// <undocumented, will likely be available in a future release>
	ReportCreate Event[events.ReportCreate]

	// Forwarded events from rAuth, currently only session deletion events are forwarded.
	//
	// <this event is special, you likely want AuthDeleteSession and AuthDeleteAllSessions instead>
	Auth Event[events.Auth]

	// A session has been deleted.
	//
	// Eq: Auth->DeleteSession
	Auth_DeleteSession Event[events.Auth_DeleteSession]

	// All sessions for this account have been deleted, optionally excluding a given ID.
	//
	// Eq: Auth->DeleteAllSessions
	Auth_DeleteAllSessions Event[events.Auth_DeleteAllSessions]
}

// Handle handles an event. It acts as a dispatch table for the events.
func (e *EventHandlers) handle(w *GatewayClient, event []byte, typ string) (events.EventInterface, error) {
	// Now, just broadcast the event to the correct handler
	switch typ {
	case "Error":
		return CreateEvent[events.Error](w, event, w.EventHandlers.Error)
	case "Authenticated":
		return CreateEvent[events.Authenticated](w, event, w.EventHandlers.Authenticated)
	case "Bulk":
		return CreateEvent[events.Bulk](w, event, w.EventHandlers.Bulk)
	case "Pong":
		return CreateEvent[events.Pong](w, event, w.EventHandlers.Pong)
	case "Ready":
		return CreateEvent[events.Ready](w, event, w.EventHandlers.Ready)
	case "Message":
		return CreateEvent[events.Message](w, event, w.EventHandlers.Message)
	case "MessageUpdate":
		return CreateEvent[events.MessageUpdate](w, event, w.EventHandlers.MessageUpdate)
	case "MessageAppend":
		return CreateEvent[events.MessageAppend](w, event, w.EventHandlers.MessageAppend)
	case "MessageDelete":
		return CreateEvent[events.MessageDelete](w, event, w.EventHandlers.MessageDelete)
	case "MessageReact":
		return CreateEvent[events.MessageReact](w, event, w.EventHandlers.MessageReact)
	case "MessageUnreact":
		return CreateEvent[events.MessageUnreact](w, event, w.EventHandlers.MessageUnreact)
	case "MessageRemoveReaction":
		return CreateEvent[events.MessageRemoveReaction](w, event, w.EventHandlers.MessageRemoveReaction)
	case "ChannelCreate":
		return CreateEvent[events.ChannelCreate](w, event, w.EventHandlers.ChannelCreate)
	case "ChannelUpdate":
		return CreateEvent[events.ChannelUpdate](w, event, w.EventHandlers.ChannelUpdate)
	case "ChannelDelete":
		return CreateEvent[events.ChannelDelete](w, event, w.EventHandlers.ChannelDelete)
	case "ChannelGroupJoin":
		return CreateEvent[events.ChannelGroupJoin](w, event, w.EventHandlers.ChannelGroupJoin)
	case "ChannelGroupLeave":
		return CreateEvent[events.ChannelGroupLeave](w, event, w.EventHandlers.ChannelGroupLeave)
	case "ChannelStartTyping":
		return CreateEvent[events.ChannelStartTyping](w, event, w.EventHandlers.ChannelStartTyping)
	case "ChannelStopTyping":
		return CreateEvent[events.ChannelStopTyping](w, event, w.EventHandlers.ChannelStopTyping)
	case "ChannelAck":
		return CreateEvent[events.ChannelAck](w, event, w.EventHandlers.ChannelAck)
	case "ServerCreate":
		return CreateEvent[events.ServerCreate](w, event, w.EventHandlers.ServerCreate)
	case "ServerUpdate":
		return CreateEvent[events.ServerUpdate](w, event, w.EventHandlers.ServerUpdate)
	case "ServerDelete":
		return CreateEvent[events.ServerDelete](w, event, w.EventHandlers.ServerDelete)
	case "ServerMemberUpdate":
		return CreateEvent[events.ServerMemberUpdate](w, event, w.EventHandlers.ServerMemberUpdate)
	case "ServerMemberJoin":
		return CreateEvent[events.ServerMemberJoin](w, event, w.EventHandlers.ServerMemberJoin)
	case "ServerMemberLeave":
		return CreateEvent[events.ServerMemberLeave](w, event, w.EventHandlers.ServerMemberLeave)
	case "ServerRoleUpdate":
		return CreateEvent[events.ServerRoleUpdate](w, event, w.EventHandlers.ServerRoleUpdate)
	case "ServerRoleDelete":
		return CreateEvent[events.ServerRoleDelete](w, event, w.EventHandlers.ServerRoleDelete)
	case "UserUpdate":
		return CreateEvent[events.UserUpdate](w, event, w.EventHandlers.UserUpdate)
	case "UserRelationship":
		return CreateEvent[events.UserRelationship](w, event, w.EventHandlers.UserRelationship)
	case "UserSettingsUpdate":
		return CreateEvent[events.UserSettingsUpdate](w, event, w.EventHandlers.UserSettingsUpdate)
	case "UserPlatformWipe":
		return CreateEvent[events.UserPlatformWipe](w, event, w.EventHandlers.UserPlatformWipe)
	case "EmojiCreate":
		return CreateEvent[events.EmojiCreate](w, event, w.EventHandlers.EmojiCreate)
	case "EmojiDelete":
		return CreateEvent[events.EmojiDelete](w, event, w.EventHandlers.EmojiDelete)
	case "WebhookCreate":
		return CreateEvent[events.WebhookCreate](w, event, w.EventHandlers.WebhookCreate)
	case "WebhookUpdate":
		return CreateEvent[events.WebhookUpdate](w, event, w.EventHandlers.WebhookUpdate)
	case "WebhookDelete":
		return CreateEvent[events.WebhookDelete](w, event, w.EventHandlers.WebhookDelete)
	case "ReportCreate":
		return CreateEvent[events.ReportCreate](w, event, w.EventHandlers.ReportCreate)
	case "Auth":
		return CreateEvent[events.Auth](w, event, w.EventHandlers.Auth)
	default:
		w.Logger.Warn("Unknown event type: " + typ)
		return nil, nil
	}
}
