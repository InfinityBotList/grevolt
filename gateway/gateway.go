package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"

	"github.com/gorilla/websocket"
	"github.com/infinitybotlist/grevolt/client/geneva"
	"go.uber.org/zap"
)

// ErrWSAlreadyOpen is thrown when you attempt to open
// a websocket that already is open.
var ErrWSAlreadyOpen = errors.New("web socket already opened")

type IOpCode int

const (
	KILL_IOpCode         IOpCode = iota
	RESTART_IOpCode      IOpCode = iota
	AUTHENTICATE_IOpCode IOpCode = iota
	ERROR_IOpCode        IOpCode = iota
	FATAL_IOpCode        IOpCode = iota
	EVENT_IOpCode        IOpCode = iota
)

type NotifyPayload struct {
	OpCode    IOpCode        // Internal library OpCode
	Data      map[string]any // Internal Opcode data
	EventData []byte         // Event data
}

type GatewayClient struct {
	sync.Mutex

	// The URL of the WS
	WSUrl string
	// API version
	APIVersion string

	// Timeout for requests
	Timeout time.Duration

	// WS Deadline
	Deadline time.Duration

	// Heartbeat interval
	HeartbeatInterval time.Duration

	// Logger to use, will be autofilled if not provided
	Logger *zap.SugaredLogger
	// Session token for requests
	SessionToken *geneva.Token

	// Websocker payload format, either json or msgpack (only these are supported by the client)
	Encoding string

	// The websocket connection
	WsConn *websocket.Conn

	// Notifications from the WS
	//
	// This is very low level and should not be used unless you know what you are doing
	NotifyChannel chan *NotifyPayload

	// Websocket event handlers
	Handlers map[string]func(eventData map[string]any)

	// This channel is fired when Wait() should return
	exitChan chan bool

	// Whether the WS is open or not
	wsOpen bool

	// Whether or not the websocket is currently being openned
	wsOpenning bool

	// Whether the WS has been killed or not
	killed bool

	// Whether the WS is restarting
	restarting bool

	// unique id describing the heartbeat
	heartbeatId string

	// websocket waitgroups
	wg sync.WaitGroup

	// already initted ws
	wsInitOnce bool
}

func (w *GatewayClient) GatewayURL() string {
	gwUrl := w.WSUrl + "?v=" + w.APIVersion + "&encoding=" + w.Encoding

	w.Logger.Info("gateway url: " + gwUrl)

	return gwUrl
}

// Opens a connection to the gateway
func (w *GatewayClient) Open() error {
	// Ensure there is only one Open() call at a time
	w.Lock()
	defer w.Unlock()

	if w.wsOpen || w.wsOpenning {
		return ErrWSAlreadyOpen
	}

	w.Logger.Info("waiting for old connections (if any) to close")
	w.wg.Wait()

	w.wsOpenning = true

	w.Logger.Info("opening connection to gateway")

	if w.Deadline == 0 {
		w.Deadline = 60 * time.Second
	}

	if w.HeartbeatInterval == 0 {
		w.HeartbeatInterval = 10 * time.Second
	}

	if w.Timeout == 0 {
		w.Timeout = 10 * time.Second
	}

	// Reset connection
	w.killed = false
	w.NotifyChannel = make(chan *NotifyPayload)

	if !w.wsInitOnce {
		w.exitChan = make(chan bool)
	}

	w.heartbeatId = ""
	w.wsOpen = false
	w.WsConn = nil

	if w.Encoding == "" {
		w.Encoding = "json"
	}

	if w.Encoding != "json" && w.Encoding != "msgpack" {
		return errors.New("invalid encoding")
	}

	// start ws connection to the gateway
	u, err := url.Parse(w.GatewayURL())

	if err != nil {
		return err
	}

	dialer := websocket.Dialer{
		HandshakeTimeout: w.Timeout,
	}

	w.WsConn, _, err = dialer.Dial(u.String(), nil)

	if err != nil {
		w.Logger.Error("connection error:", err)
		w.Close()
		return errors.New("failed to connect to gateway: " + err.Error())
	}

	go w.handleNotify()
	time.Sleep(1 * time.Second)
	go w.readMessages()

	w.Logger.Info("opened connection to gateway")

	w.wsOpen = true
	w.wsOpenning = false
	w.wsInitOnce = true
	return nil
}

func (w *GatewayClient) Close() {
	close(w.exitChan)
	w.NotifyChannel <- &NotifyPayload{
		OpCode: KILL_IOpCode,
	}
}

// Wait for the gateway to close
func (w *GatewayClient) Wait() {
	for payload := range w.exitChan {
		w.Logger.Info("received exit payload: ", payload)
		if payload {
			// Close the websocket
			return
		}
	}
}

// Waits for the gateway to close, then sends a notification to the notify channel
func (w *GatewayClient) OnDone(notify chan bool) {
	<-w.exitChan
	notify <- true
}

type internalMessage struct {
	Type  string `json:"type"`
	Error string `json:"error,omitempty"`
}

func (w *GatewayClient) readMessages() {
	w.wg.Add(1)

	defer func() {
		w.Logger.Info("wg decr (readMessages)")
		w.wg.Done()
	}()

	fmt.Println("readMessages() started")

	w.WsConn.SetReadDeadline(time.Now().Add(w.Deadline))

	// Before doing anything else, send AUTHENTICATE_IOpCode
	w.NotifyChannel <- &NotifyPayload{
		OpCode: AUTHENTICATE_IOpCode,
	}

	w.WsConn.SetCloseHandler(func(code int, text string) error {
		w.wsOpen = false
		w.Logger.Info("websocket closed: ", code, text)
		w.NotifyChannel <- &NotifyPayload{
			OpCode: ERROR_IOpCode,
		}

		return nil
	})

	for {
		// Check that connection isn't closed
		if w.killed {
			w.Logger.Info("readMessages() killed")
			return
		}

		_, message, err := w.WsConn.ReadMessage()

		var data internalMessage

		// If we have a message, try and decode it first, before checking for a close code
		if len(message) > 0 {
			switch w.Encoding {
			case "json":
				err = json.Unmarshal(message, &data)

				if err != nil {
					w.Logger.Error("failed to unmarshal message: " + err.Error())
					data = internalMessage{
						Type: "InternalError",
					}
				}
			case "msgpack":
				err = msgpack.Unmarshal(message, &data)

				if err != nil {
					w.Logger.Error("failed to unmarshal message: " + err.Error())
					data = internalMessage{
						Type: "InternalError",
					}
				}
			}

			// Error handling here, before checking frames, allows for detection of invalid auth credential errors
			var err bool = true
			switch data.Type {
			case "Pong":
				w.Logger.Info("recieved pong from gateway")
				w.WsConn.SetReadDeadline(time.Now().Add(w.Deadline))
				err = false
			case "NotFound": // Undocumented, but means that auth credentials are invalid
				w.Logger.Error("invalid auth credentials")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: FATAL_IOpCode,
					Data: map[string]any{
						"error": "invalid auth credentials",
					},
				}
			case "LabelMe":
				w.Logger.Info("received LabelMe")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: ERROR_IOpCode,
					Data: map[string]any{
						"error": "recieved unknown error: label me",
					},
				}
			case "InternalError":
				w.Logger.Info("received InternalError")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: ERROR_IOpCode,
					Data: map[string]any{
						"error": "recieved unknown error: internal error",
					},
				}
			case "InvalidSession":
				w.Logger.Info("received InvalidSession")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: FATAL_IOpCode,
					Data: map[string]any{
						"error": "invalid session",
					},
				}
			case "OnboardingNotFinished":
				w.Logger.Info("received OnboardingNotFinished")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: FATAL_IOpCode,
					Data: map[string]any{
						"error": "onboarding not finished [OnboardingNotFinished]",
					},
				}
			// TODO: rethink this: ERROR vs NOTIFY
			case "AlreadyAuthenticated":
				w.Logger.Info("received AlreadyAuthenticated")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: ERROR_IOpCode,
					Data: map[string]any{
						"error": "already authenticated [AlreadyAuthenticated]",
					},
				}
			case "Authenticated":
				w.Logger.Info("received Authenticated flag")
				hbId := uuid.New().String()

				w.heartbeatId = hbId
				go w.heartbeat(hbId)
				err = false
			default: // No error, continue
				err = false
			}

			if err {
				time.Sleep(1 * time.Second)
				return
			}
		}

		// Now we can check error freely as this is a close code
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				w.Logger.Info("error: %v", err)
			}

			// Check if its a "use of closed network connection" error
			//
			// These are not fatal and should definitely not be spawning a notify payload
			if strings.Contains(err.Error(), "use of closed network connection") {
				w.Logger.Info("websocket closed, exiting readMessages()")
				w.wsOpen = false
				return
			}

			// Send whatever we have to the notify channel
			w.Logger.Error("failed to read message: " + err.Error())
			w.NotifyChannel <- &NotifyPayload{
				OpCode: ERROR_IOpCode,
				Data: map[string]any{
					"error": err.Error(),
					"msg":   message,
				},
			}

			w.wsOpen = false
			return
		}

		// Send event over notify channel
		w.NotifyChannel <- &NotifyPayload{
			OpCode:    EVENT_IOpCode,
			EventData: message,
		}

		// Recieved message successfully, extend deadline
		w.WsConn.SetReadDeadline(time.Now().Add(w.Deadline))
	}
}

func (w *GatewayClient) Send(data map[string]any) error {
	err := w.WsConn.SetWriteDeadline(time.Now().Add(w.Deadline))

	if err != nil {
		return errors.New("failed to set write deadline: " + err.Error())
	}

	switch w.Encoding {
	case "json":
		// Marshal json
		jsonData, err := json.Marshal(data)

		if err != nil {
			return errors.New("failed to marshal json: " + err.Error())
		}

		// Send json
		err = w.WsConn.WriteMessage(websocket.TextMessage, jsonData)

		if err != nil {
			return errors.New("failed to send json: " + err.Error())
		}
	case "msgpack":
		// Marshal msgpack
		msgpackData, err := msgpack.Marshal(data)

		if err != nil {
			return errors.New("failed to marshal msgpack: " + err.Error())
		}

		// Send msgpack
		err = w.WsConn.WriteMessage(websocket.BinaryMessage, msgpackData)

		if err != nil {
			return errors.New("failed to send msgpack: " + err.Error())
		}
	}

	return nil
}

func (w *GatewayClient) handleNotify() {
	restarter := func() {
		w.Logger.Info("restarting connection to gateway")

		// If killed, don't restart
		if w.killed {
			w.Logger.Info("not restarting connection to gateway because it was killed")
			return
		}

		// Send restart opcode
		w.Logger.Info("sending restart opcode")
		w.WsConn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "closeWsConn"), time.Now().Add(w.Deadline))
		w.WsConn.Close()
		w.restarting = true
		w.wsOpen = false

		// Avoid leaking channels
		close(w.NotifyChannel)

		w.Logger.Info("opening new connection to gateway")
		go w.Open()
	}

	w.wg.Add(1)
	defer func() {
		w.Logger.Info("wg decr (handleNotify)")
		w.wg.Done()
	}()

	for payload := range w.NotifyChannel {
		switch payload.OpCode {
		case KILL_IOpCode:
			w.Logger.Info("killing connection to gateway")
			w.killed = true
			w.wsOpen = false
			w.heartbeatId = ""
			w.WsConn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "closeWsConn"), time.Now().Add(w.Deadline))
			w.WsConn.Close()
			w.exitChan <- true
			return
		case RESTART_IOpCode:
			restarter()
			return
		case AUTHENTICATE_IOpCode:
			// Send authenticate command frame
			w.Logger.Info("sending authenticate command frame")
			w.Send(map[string]any{
				"type":  "Authenticate",
				"token": w.SessionToken.Token,
			})
		case ERROR_IOpCode:
			w.Logger.Error("error from gateway: ", payload)
			restarter()
			return
		case FATAL_IOpCode:
			w.Logger.Error("fatal error from gateway: ", payload)
			w.killed = true
			w.wsOpen = false
			w.heartbeatId = ""
			return
		case EVENT_IOpCode:
			// TODO: event handling
			fmt.Println(string(payload.EventData))
		}
	}

	if w.wsOpen {
		restarter()
	}
}

func (w *GatewayClient) heartbeat(hbId string) {
	w.Logger.Info("starting heartbeat, ", hbId)
	// Create new ticker
	ticker := time.NewTicker(w.HeartbeatInterval)

	// Send heartbeat
	for range ticker.C {
		if w.heartbeatId != hbId {
			w.Logger.Error("heartbeat id mismatch")
			ticker.Stop()
			return // Not the current heartbeat ws
		}

		if !w.wsOpen || w.killed {
			continue
		}

		w.Logger.Info("sending heartbeat")
		w.Send(map[string]any{
			"type": "Ping",
			"data": time.Now().Unix(),
		})
	}
}
