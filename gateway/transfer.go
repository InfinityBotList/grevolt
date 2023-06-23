package gateway

import (
	"encoding/json"

	"github.com/vmihailenco/msgpack/v5"
)

func (w *GatewayClient) Recieve(data []byte, dst any) error {
	switch w.Encoding {
	case "json":
		return json.Unmarshal(data, &dst)
	case "msgpack":
		return msgpack.Unmarshal(data, &dst)
	}

	return nil
}
