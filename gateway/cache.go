package gateway

import (
	"github.com/infinitybotlist/grevolt/types/events"
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
	}

	return nil
}
