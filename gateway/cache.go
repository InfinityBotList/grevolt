package gateway

import (
	"errors"

	"github.com/infinitybotlist/grevolt/cache/diff"
	"github.com/infinitybotlist/grevolt/cache/store"
	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/types/events"
	"go.uber.org/zap"
)

func (w *GatewayClient) cacheEvent(typ string, d events.EventInterface) error {
	switch typ {
	case "Ready":
		evt := d.(*events.Ready)

		// Cache all users
		for _, user := range evt.Users {
			err := w.SharedState.AddUser(user)

			if err != nil {
				return err
			}
		}

		// Cache all servers
		for _, server := range evt.Servers {
			err := w.SharedState.AddServer(server)

			if err != nil {
				return err
			}
		}

		// Cache all channels
		for _, channel := range evt.Channels {
			err := w.SharedState.AddChannel(channel)

			if err != nil {
				return err
			}
		}

		// Cache all members
		for _, member := range evt.Members {
			err := w.SharedState.AddMember(member)

			if err != nil {
				return err
			}
		}
	case "ChannelCreate":
		evt := d.(*events.ChannelCreate)

		// Cache the channel
		err := w.SharedState.AddChannel(evt.Channel)

		if err != nil {
			return err
		}
	case "ChannelUpdate":
		evt := d.(*events.ChannelUpdate)

		// Look for the channel in cache
		c, err := w.SharedState.GetChannel(evt.Id)

		// Cache channel as partial if not found
		if errors.Is(err, store.ErrNotFound) {
			w.Logger.Debug("Channel not found in cache, caching as partial", zap.String("channel", evt.Id))
			// Cache the channel
			err := w.SharedState.AddChannel(evt.Data)

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// Update the channel
		newChan := diff.PartialUpdate[types.Channel](c, evt.Data)

		w.Logger.Debug("Updated channel", zap.Any("now", newChan), zap.Any("patch", evt.Data))

		err = w.SharedState.AddChannel(newChan)

		if err != nil {
			return err
		}
	case "ChannelDelete":
		evt := d.(*events.ChannelDelete)

		// Delete the channel from cache
		err := w.SharedState.DeleteChannel(evt.Id)

		if err != nil {
			return err
		}
	case "ServerCreate":
		evt := d.(*events.ServerCreate)

		// Cache the server
		err := w.SharedState.AddServer(evt.Server)

		if err != nil {
			w.Logger.Error(
				"Failed to cache server",
				zap.Error(err),
				zap.String("type", typ),
				zap.String("server", evt.Server.Id),
			)
		}

		if evt.Channels != nil {
			// Cache all channels
			for _, channel := range evt.Channels {
				err := w.SharedState.AddChannel(channel)

				if err != nil {
					w.Logger.Error(
						"Failed to cache channel",
						zap.Error(err),
						zap.String("type", typ),
						zap.String("channel", channel.Id),
					)
				}
			}
		}
	case "ServerUpdate":
		evt := d.(*events.ServerUpdate)

		// Look for the server in cache
		s, err := w.SharedState.GetServer(evt.Id)

		// Cache server as partial if not found
		if errors.Is(err, store.ErrNotFound) {
			w.Logger.Debug("Server not found in cache, caching as partial", zap.String("server", evt.Id))
			// Cache the server
			err := w.SharedState.AddServer(evt.Data)

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// Update the server
		newServ := diff.PartialUpdate[types.Server](s, evt.Data)

		w.Logger.Debug("Updated channel", zap.Any("now", newServ), zap.Any("patch", evt.Data))

		err = w.SharedState.AddServer(newServ)

		if err != nil {
			return err
		}
	case "ServerDelete":
		evt := d.(*events.ServerDelete)

		// Delete the server from cache
		err := w.SharedState.DeleteServer(evt.Id)

		if err != nil {
			return err
		}
	case "ServerMemberUpdate":
		evt := d.(*events.ServerMemberUpdate)

		// Look for the member in cache
		m, err := w.SharedState.GetMember(evt.Id.Server, evt.Id.User)

		// Cache member as partial if not found
		if errors.Is(err, store.ErrNotFound) {
			w.Logger.Debug("Member not found in cache, caching as partial", zap.String("server", evt.Id.Server), zap.String("member", evt.Id.User))
			// Cache the member
			evt.Data.Id = evt.Id
			err := w.SharedState.AddMember(evt.Data)

			if err != nil {
				return err
			}

			return nil
		} else if err != nil {
			return err
		}

		// Update the member
		newMember := diff.PartialUpdate[types.Member](m, evt.Data)

		w.Logger.Debug("Updated member", zap.Any("now", newMember), zap.Any("patch", evt.Data))

		err = w.SharedState.AddMember(newMember)

		if err != nil {
			return err
		}
	case "UserUpdate":
		evt := d.(*events.UserUpdate)

		// Look for the user in cache
		u, err := w.SharedState.GetUser(evt.Id)

		// Cache user as partial if not found
		if errors.Is(err, store.ErrNotFound) {
			// Fetch from rest here, the info in a UserUpdate is not enough for being a
			// starting point for a requestion
			//
			// Rest automatically performs caching
			if !w.GatewayCache.DisableAutoRestFetching {
				_, err = w.RestClient.FetchUser(evt.Id)
				if err != nil {
					return err
				}
			}

			return nil
		} else if err != nil {
			return err
		}

		// Update the user
		newUser := diff.PartialUpdate[types.User](u, evt.Data)

		w.Logger.Debug("Updated user [UserUpdate]", zap.Any("now", newUser), zap.Any("patch", evt.Data))

		err = w.SharedState.AddUser(newUser)

		if err != nil {
			return err
		}
	case "UserRelationship":
		evt := d.(*events.UserRelationship)

		// Look for the user in cache
		u, err := w.SharedState.GetUser(evt.Id)

		// Cache user as partial if not found
		if errors.Is(err, store.ErrNotFound) {
			w.Logger.Debug("User not found in cache, caching as partial", zap.String("user", evt.Id))
			// Cache the user
			err := w.SharedState.AddUser(evt.User)

			if err != nil {
				return err
			}

			return nil
		} else if err != nil {
			return err
		}

		// Update the user
		newUser := diff.PartialUpdate[types.User](u, evt.User)

		w.Logger.Debug("Updated user", zap.Any("now", newUser), zap.Any("patch", evt.User))

		err = w.SharedState.AddUser(newUser)

		if err != nil {
			return err
		}
	case "EmojiCreate":
		evt := d.(*events.EmojiCreate)

		// Look for the emoji in cache
		e, err := w.SharedState.GetEmoji(evt.Id)

		// Cache emoji as partial if not found
		if errors.Is(err, store.ErrNotFound) {
			w.Logger.Debug("Emoji not found in cache, caching as partial", zap.String("emoji", evt.Id))
			// Cache the emoji
			err := w.SharedState.AddEmoji(evt.Emoji)

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		// Update the emoji
		newEmoji := diff.PartialUpdate[types.Emoji](e, evt.Emoji)

		w.Logger.Debug("Updated emoji", zap.Any("now", newEmoji), zap.Any("patch", evt.Emoji))

		err = w.SharedState.AddEmoji(newEmoji)

		if err != nil {
			return err
		}
	case "EmojiDelete":
		evt := d.(*events.EmojiDelete)

		// Delete the emoji from cache
		err := w.SharedState.DeleteEmoji(evt.Id)

		if err != nil {
			return err
		}
	}

	return nil
}
