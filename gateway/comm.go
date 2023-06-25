package gateway

func (w *GatewayClient) BeginTyping(channelID string) error {
	return w.Send(map[string]any{
		"type":    "BeginTyping",
		"channel": channelID,
	})
}

func (w *GatewayClient) EndTyping(channelID string) error {
	return w.Send(map[string]any{
		"type":    "EndTyping",
		"channel": channelID,
	})
}
