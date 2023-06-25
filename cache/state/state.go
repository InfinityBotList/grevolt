package state

import (
	"github.com/infinitybotlist/grevolt/cache/store"
	"github.com/infinitybotlist/grevolt/types"
)

// Implements a state ontop of a store implementation
type State struct {
	// Users
	Users store.Store[types.User]

	// Servers
	Servers store.Store[types.Server]

	// Channels
	Channels store.Store[types.Channel]

	// Members
	Members store.Store[types.Member]

	// Emojis
	Emojis store.Store[types.Emoji]
}

// Users

// GetUser returns a user from the state
func (s *State) GetUser(id string) (*types.User, error) {
	return s.Users.Get(id)
}

// AddUser adds a user to the state
func (s *State) AddUser(u *types.User) error {
	return s.Users.Set(u.Id, u)
}

// DeleteUser deletes a user from the state
func (s *State) DeleteUser(id string) error {
	return s.Users.Delete(id)
}

// Servers

// GetServer returns a server from the state
func (s *State) GetServer(id string) (*types.Server, error) {
	return s.Servers.Get(id)
}

// AddServer adds a server to the state
func (s *State) AddServer(se *types.Server) error {
	return s.Servers.Set(se.Id, se)
}

// DeleteServer deletes a server from the state
func (s *State) DeleteServer(id string) error {
	return s.Servers.Delete(id)
}

// Channels

// GetChannel returns a channel from the state
func (s *State) GetChannel(id string) (*types.Channel, error) {
	return s.Channels.Get(id)
}

// AddChannel adds a channel to the state
func (s *State) AddChannel(c *types.Channel) error {
	return s.Channels.Set(c.Id, c)
}

// DeleteChannel deletes a channel from the state
func (s *State) DeleteChannel(id string) error {
	return s.Channels.Delete(id)
}

// Members

// GetMember returns a member from the state
func (s *State) GetMember(serverId, userId string) (*types.Member, error) {
	return s.Members.Get(serverId + "/" + userId)
}

// AddMember adds a member to the state
func (s *State) AddMember(m *types.Member) error {
	return s.Members.Set(m.Id.Server+"/"+m.Id.User, m)
}

// DeleteMember deletes a member from the state
func (s *State) DeleteMember(serverId, userId string) error {
	return s.Members.Delete(serverId + "/" + userId)
}

// Emojis

// GetEmoji returns an emoji from the state
func (s *State) GetEmoji(id string) (*types.Emoji, error) {
	return s.Emojis.Get(id)
}

// AddEmoji adds an emoji to the state
func (s *State) AddEmoji(e *types.Emoji) error {
	return s.Emojis.Set(e.Id, e)
}

// DeleteEmoji deletes an emoji from the state
func (s *State) DeleteEmoji(id string) error {
	return s.Emojis.Delete(id)
}
