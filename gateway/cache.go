package gateway

import (
	"github.com/infinitybotlist/grevolt/cache/diff"
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
		if err != nil {
			// Cache the channel
			err := w.SharedState.AddChannel(evt.Data)

			if err != nil {
				return err
			}
		}

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
	}

	return nil
}
