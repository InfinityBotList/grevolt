package main

import (
	"strings"

	"github.com/infinitybotlist/grevolt/cache/store/basicstore"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/types"
	"github.com/infinitybotlist/grevolt/types/events"
)

const prefix = "!"

func messageHandler(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *events.Message) {
	// TODO: Add bot check, but state isn't advanced enough yet

	if !strings.HasPrefix(evt.Content, prefix) {
		return // Wrong prefix
	}

	msg := strings.TrimPrefix(evt.Content, prefix)

	args := strings.Split(msg, " ")

	go w.BeginTyping(evt.Channel)

	switch args[0] {
	case "args":
		// Send a message to the channel
		w.RestClient.SendMessage(evt.Channel, &types.DataMessageSend{
			Content: "Args: " + strings.Join(args[1:], " "),
		})
	case "ping":
		// Send a message to the channel
		w.RestClient.SendMessage(evt.Channel, &types.DataMessageSend{
			Content: "Pong: " + w.Heartbeat.Latency().String(),
		})
	case "cacheduserlist":
		store := w.SharedState.Users.(*basicstore.BasicStore[types.User])

		kvs := store.KeyValuePairs()

		// Send a message to the channel
		w.RestClient.SendMessage(evt.Channel, &types.DataMessageSend{
			Content: "**Cached users**\n" + strings.Join(kvs, "\n"),
		})
	}

	go w.EndTyping(evt.Channel)
}
