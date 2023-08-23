package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"
)

func (w *GatewayClient) Decode(data []byte, dst any) error {
	switch w.Encoding {
	case "json":
		return json.Unmarshal(data, &dst)
	case "msgpack":
		// Create buffer
		buf := bytes.NewBuffer(data)

		// Decode msgpack
		decoded := msgpack.NewDecoder(buf)

		decoded.SetCustomStructTag("json")

		err := decoded.Decode(&dst)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (w *GatewayClient) Encode(data any) ([]byte, error) {
	switch w.Encoding {
	case "json":
		// Marshal json
		jsonData, err := json.Marshal(data)

		if err != nil {
			return nil, errors.New("failed to marshal json: " + err.Error())
		}

		return jsonData, nil
	case "msgpack":
		// Marshal msgpack
		var buf bytes.Buffer
		enc := msgpack.NewEncoder(&buf)
		enc.SetCustomStructTag("json")
		err := enc.Encode(data)

		if err != nil {
			return nil, errors.New("failed to marshal msgpack: " + err.Error())
		}

		msgpackData := buf.Bytes()
		return msgpackData, nil
	}

	return nil, errors.New("invalid encoding")
}

func (w *GatewayClient) Send(data map[string]any) error {
	if w.State != WsStateOpen {
		return errors.New("websocket not open")
	}

	err := w.WsConn.SetWriteDeadline(time.Now().Add(w.Deadline))

	if err != nil {
		return errors.New("failed to set write deadline: " + err.Error())
	}

	sendBytes, err := w.Encode(data)

	if err != nil {
		return errors.New("failed to encode data: " + err.Error())
	}

	w.Logger.Debug("Sending data", zap.Any("data", data), zap.Binary("sendBytes", sendBytes))

	switch w.Encoding {
	case "json":
		// Send json
		err = w.WsConn.WriteMessage(websocket.TextMessage, sendBytes)

		if err != nil {
			return errors.New("failed to send json: " + err.Error())
		}
	case "msgpack":
		// Send msgpack
		err = w.WsConn.WriteMessage(websocket.BinaryMessage, sendBytes)

		if err != nil {
			return errors.New("failed to send msgpack: " + err.Error())
		}
	}

	return nil
}
