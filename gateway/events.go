package gateway

import (
	"errors"

	"github.com/infinitybotlist/grevolt/types/events"
	"go.uber.org/zap"
)

type EventContext struct {
	// Raw event data
	Raw []byte
}

type Event[T events.EventInterface] func(w *GatewayClient, ctx *EventContext, evt *T)

// Emits an event to a function
func CreateEvent[T events.EventInterface](
	w *GatewayClient,
	data []byte,
	fn Event[T],
) error {
	var evtMarshalled *T

	err := w.Decode(data, &evtMarshalled)

	if err != nil {
		return errors.New("decode error: " + err.Error())
	}

	if fn == nil {
		return nil
	}

	fn(w, &EventContext{
		Raw: data,
	}, evtMarshalled)

	return nil
}

func (w *GatewayClient) HandleBulk(event []byte) error {
	var bulkData *events.Bulk

	err := w.Decode(event, &bulkData)

	if err != nil {
		return err
	}

	for _, evt := range bulkData.V {
		bytes, err := w.Encode(evt)

		if err != nil {
			return err
		}

		typ, ok := evt["type"]

		if !ok {
			w.Logger.Error(
				"event has no type",
				zap.Binary("evt", bytes),
			)
		}

		typStr, ok := typ.(string)

		if !ok {
			w.Logger.Error(
				"event type is not a string",
				zap.Binary("evt", bytes),
			)
		}

		w.HandleEvent(bytes, typStr)
	}

	return nil
}

func (w *GatewayClient) HandleAuth(event []byte) error {
	var authData *events.Auth

	err := w.Decode(event, &authData)

	if err != nil {
		return err
	}

	switch authData.Type {
	case "DeleteSession":
		return CreateEvent[events.AuthDeleteSession](w, event, w.EventHandlers.AuthDeleteSession)
	case "DeleteAllSessions":
		return CreateEvent[events.AuthDeleteAllSessions](w, event, w.EventHandlers.AuthDeleteAllSessions)
	default:
		w.Logger.Warn(
			"Unknown auth event type",
			zap.String("eventType", authData.Type),
		)
	}

	return nil
}

func (w *GatewayClient) HandleEvent(event []byte, typ string) {
	if w.RawSinkFunc != nil && len(w.RawSinkFunc) > 0 {
		for _, fn := range w.RawSinkFunc {
			fn(w, event, typ)
		}
	}

	if typ == "Bulk" {
		// Bulk is a bit unique, special handling is required
		err := w.HandleBulk(event)

		if err != nil {
			w.Logger.Error(
				"bulk handler failed",
				zap.Error(err),
			)
		}
	}

	if typ == "Auth" {
		// Auth is a bit unique because of event it, handle it
		err := w.HandleAuth(event)

		if err != nil {
			w.Logger.Error(
				"auth handler failed",
				zap.Error(err),
			)
		}
	}

	var err error
	switch typ {
	case "Error":
		err = CreateEvent[events.Error](w, event, w.EventHandlers.Error)
	case "Authenticated":
		err = CreateEvent[events.Authenticated](w, event, w.EventHandlers.Authenticated)
	case "Bulk":
		err = CreateEvent[events.Bulk](w, event, w.EventHandlers.Bulk)
	case "Pong":
		err = CreateEvent[events.Pong](w, event, w.EventHandlers.Pong)
	case "Ready":
		err = CreateEvent[events.Ready](w, event, w.EventHandlers.Ready)
	case "Message":
		err = CreateEvent[events.Message](w, event, w.EventHandlers.Message)
	case "MessageUpdate":
		err = CreateEvent[events.MessageUpdate](w, event, w.EventHandlers.MessageUpdate)
	case "MessageAppend":
		err = CreateEvent[events.MessageAppend](w, event, w.EventHandlers.MessageAppend)
	case "MessageDelete":
		err = CreateEvent[events.MessageDelete](w, event, w.EventHandlers.MessageDelete)
	case "MessageReact":
		err = CreateEvent[events.MessageReact](w, event, w.EventHandlers.MessageReact)
	case "MessageUnreact":
		err = CreateEvent[events.MessageUnreact](w, event, w.EventHandlers.MessageUnreact)
	case "MessageRemoveReaction":
		err = CreateEvent[events.MessageRemoveReaction](w, event, w.EventHandlers.MessageRemoveReaction)
	case "ChannelCreate":
		err = CreateEvent[events.ChannelCreate](w, event, w.EventHandlers.ChannelCreate)
	case "ChannelUpdate":
		err = CreateEvent[events.ChannelUpdate](w, event, w.EventHandlers.ChannelUpdate)
	case "ChannelDelete":
		err = CreateEvent[events.ChannelDelete](w, event, w.EventHandlers.ChannelDelete)
	case "ChannelGroupJoin":
		err = CreateEvent[events.ChannelGroupJoin](w, event, w.EventHandlers.ChannelGroupJoin)
	case "ChannelGroupLeave":
		err = CreateEvent[events.ChannelGroupLeave](w, event, w.EventHandlers.ChannelGroupLeave)
	case "ChannelStartTyping":
		err = CreateEvent[events.ChannelStartTyping](w, event, w.EventHandlers.ChannelStartTyping)
	case "ChannelStopTyping":
		err = CreateEvent[events.ChannelStopTyping](w, event, w.EventHandlers.ChannelStopTyping)
	case "ChannelAck":
		err = CreateEvent[events.ChannelAck](w, event, w.EventHandlers.ChannelAck)
	case "ServerCreate":
		err = CreateEvent[events.ServerCreate](w, event, w.EventHandlers.ServerCreate)
	case "ServerUpdate":
		err = CreateEvent[events.ServerUpdate](w, event, w.EventHandlers.ServerUpdate)
	case "ServerDelete":
		err = CreateEvent[events.ServerDelete](w, event, w.EventHandlers.ServerDelete)
	case "ServerMemberUpdate":
		err = CreateEvent[events.ServerMemberUpdate](w, event, w.EventHandlers.ServerMemberUpdate)
	case "ServerMemberJoin":
		err = CreateEvent[events.ServerMemberJoin](w, event, w.EventHandlers.ServerMemberJoin)
	case "ServerMemberLeave":
		err = CreateEvent[events.ServerMemberLeave](w, event, w.EventHandlers.ServerMemberLeave)
	case "ServerRoleUpdate":
		err = CreateEvent[events.ServerRoleUpdate](w, event, w.EventHandlers.ServerRoleUpdate)
	case "ServerRoleDelete":
		err = CreateEvent[events.ServerRoleDelete](w, event, w.EventHandlers.ServerRoleDelete)
	case "UserUpdate":
		err = CreateEvent[events.UserUpdate](w, event, w.EventHandlers.UserUpdate)
	case "UserRelationship":
		err = CreateEvent[events.UserRelationship](w, event, w.EventHandlers.UserRelationship)
	case "UserSettingsUpdate":
		err = CreateEvent[events.UserSettingsUpdate](w, event, w.EventHandlers.UserSettingsUpdate)
	case "UserPlatformWipe":
		err = CreateEvent[events.UserPlatformWipe](w, event, w.EventHandlers.UserPlatformWipe)
	case "EmojiCreate":
		err = CreateEvent[events.EmojiCreate](w, event, w.EventHandlers.EmojiCreate)
	case "EmojiDelete":
		err = CreateEvent[events.EmojiDelete](w, event, w.EventHandlers.EmojiDelete)
	case "WebhookCreate":
		err = CreateEvent[events.WebhookCreate](w, event, w.EventHandlers.WebhookCreate)
	case "WebhookUpdate":
		err = CreateEvent[events.WebhookUpdate](w, event, w.EventHandlers.WebhookUpdate)
	case "WebhookDelete":
		err = CreateEvent[events.WebhookDelete](w, event, w.EventHandlers.WebhookDelete)
	case "ReportCreate":
		err = CreateEvent[events.ReportCreate](w, event, w.EventHandlers.ReportCreate)
	case "Auth":
		err = CreateEvent[events.Auth](w, event, w.EventHandlers.Auth)
	default:
		w.Logger.Warn("Unknown event type: " + typ)
	}

	if err != nil {
		w.Logger.Error(
			"Event handling failed",
			zap.Error(err),
			zap.String("type", typ),
		)
	}
}

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
	AuthDeleteSession Event[events.AuthDeleteSession]

	// All sessions for this account have been deleted, optionally excluding a given ID.
	//
	// Eq: Auth->DeleteAllSessions
	AuthDeleteAllSessions Event[events.AuthDeleteAllSessions]

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
}
