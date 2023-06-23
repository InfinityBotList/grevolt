package gateway

import "fmt"

func (w *GatewayClient) HandleEvent(event []byte) {
	fmt.Println(string(event))

}
