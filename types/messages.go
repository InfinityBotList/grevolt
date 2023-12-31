package types

import (
	"encoding/json"

	"github.com/infinitybotlist/grevolt/types/timestamp"
)

// MessageSort : Sort used for retrieving messages
type MessageSort string

// List of MessageSort
const (
	RELEVANCE_MessageSort MessageSort = "Relevance"
	LATEST_MessageSort    MessageSort = "Latest"
	OLDEST_MessageSort    MessageSort = "Oldest"
)

// Filter and sort messages by time
type MessageQuery struct {
	// Maximum number of messages to fetch
	//
	// For fetching nearby messages, this is `(limit + 1)`.
	Limit uint64 `json:"limit,omitempty"`

	// Message id before which messages should be fetched
	Before string `json:"before,omitempty"`

	// Message id after which messages should be fetched
	After string `json:"after,omitempty"`

	// Message sort direction
	Sort MessageSort `json:"sort,omitempty"`

	// Message id to search around
	//
	// Specifying 'nearby' ignores 'before', 'after' and 'sort'.
	// It will also take half of limit rounded as the limits to each side.
	// It also fetches the message ID specified.
	Nearby string `json:"nearby,omitempty"`

	// Whether to include user (and member, if server channel) objects
	IncludeUsers bool `json:"include_users,omitempty"`
}

// Filter and sort messages
type MessageSearchQuery struct {
	// Full-text search query
	//
	// See MongoDB documentation for more information.
	//
	// <https://docs.mongodb.com/manual/text-search/#-text-operator>
	//
	// <Mandatory>
	Query string `json:"query"`

	// Maximum number of messages to fetch
	//
	// For fetching nearby messages, this is `(limit + 1)`.
	Limit uint64 `json:"limit,omitempty"`

	// Message id before which messages should be fetched
	Before string `json:"before,omitempty"`

	// Message id after which messages should be fetched
	After string `json:"after,omitempty"`

	// Message sort direction
	Sort MessageSort `json:"sort,omitempty"`

	// Whether to include user (and member, if server channel) objects
	IncludeUsers bool `json:"include_users"`
}

// Note that this struct is used for both include_users and no_include_users
type MessageFetchResponse struct {
	// Whether to include user (and member, if server channel) objects, used internally by MarshalJSON/UnmarshalJSON
	IncludeUsers bool

	// List of messages
	Messages []*Message `json:"messages,omitempty"`

	// List of users, only set if IncludeUsers is true
	Users []*User `json:"users,omitempty"`

	// List of members, only set if IncludeUsers is true
	Members []*Member `json:"members,omitempty"`
}

func (m *MessageFetchResponse) UnmarshalJSON(data []byte) error {
	if m.IncludeUsers {
		var d struct {
			Messages []*Message `json:"messages,omitempty"`
			Users    []*User    `json:"users,omitempty"`
			Members  []*Member  `json:"members,omitempty"`
		}
		if err := json.Unmarshal(data, &d); err != nil {
			return err
		}
		m.Messages = d.Messages
		m.Users = d.Users
		m.Members = d.Members
		m.IncludeUsers = true
		return nil
	} else {
		var d []*Message

		if err := json.Unmarshal(data, &d); err != nil {
			return err
		}

		m.Messages = d

		return nil
	}
}

func (m *MessageFetchResponse) MarshalJSON() ([]byte, error) {
	if m.IncludeUsers {
		aux := struct {
			Messages []*Message `json:"messages,omitempty"`
			Users    []*User    `json:"users,omitempty"`
			Members  []*Member  `json:"members,omitempty"`
		}{
			Messages: m.Messages,
			Users:    m.Users,
			Members:  m.Members,
		}

		return json.Marshal(&aux)
	} else {
		return json.Marshal(m.Messages)
	}
}

// Message : A message sent in a channel
//
// This struct is data needed to send a message.
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
	Masquerade *MessageMasquerade `json:"masquerade,omitempty"`
	// Information about how this message should be interacted with
	Interactions *MessageInteractions `json:"interactions,omitempty"`
}

// This struct is data needed to edit a message.
type DataMessageEdit struct {
	// New message content
	Content string `json:"content,omitempty"`
	// Embeds to include in the message
	Embeds []SendableEmbed `json:"embeds,omitempty"`
}

// Information to guide interactions on this message
type MessageInteractions struct {
	// Reactions which should always appear and be distinct
	Reactions []string `json:"reactions,omitempty"`
	// Whether reactions should be restricted to the given list  Can only be set to true if reactions list is of at least length 1
	RestrictReactions bool `json:"restrict_reactions,omitempty"`
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
	Edited timestamp.Timestamp `json:"edited,omitempty"`
	// Attached embeds to this message
	Embeds []*MessageEmbed `json:"embeds,omitempty"`
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

// Information about the webhook bundled with Message
type MessageWebhook struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
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

// Representation of a message reply before it is sent.
type Reply struct {
	// Message Id
	Id string `json:"id"`
	// Whether this reply should mention the message's author
	Mention bool `json:"mention"`
}

// System message
type MessageSystem struct {
	// System message type
	Type string `json:"type"`

	// System message content
	Content string `json:"content"`
}

type MessageIds struct {
	// Message IDs
	Ids []string `json:"ids"`
}

type DataReactionsRemove struct {
	// Remove a specific user's reaction
	UserId string `json:"user_id,omitempty"`

	// Remove all reactions
	RemoveAll bool `json:"remove_all,omitempty"`
}
