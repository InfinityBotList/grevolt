package main

import (
	"strings"

	"github.com/infinitybotlist/grevolt/cache/resolvers"
	"github.com/infinitybotlist/grevolt/cache/store/basicstore"
	"github.com/infinitybotlist/grevolt/gateway"
	"github.com/infinitybotlist/grevolt/gateway/events"
	"github.com/infinitybotlist/grevolt/types"
	"go.uber.org/zap"
)

const prefix = "!"

func messageHandler(w *gateway.GatewayClient, ctx *gateway.EventContext, evt *events.Message) {
	// Resolve author, this is needed to check if its a bot or not
	u, err := resolvers.ResolveUser(w.RestClient, evt.Author)

	if err != nil {
		w.Logger.Named("messageHandler").Error("Failed to resolve user", zap.String("user", evt.Author))
	}

	if u.Bot != nil {
		w.Logger.Named("messageHandler").Info("Ignoring bot", zap.Any("bot", u))
		return
	}

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
			Content: "Pong: " + w.LastHeartbeat.Latency().String(),
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
