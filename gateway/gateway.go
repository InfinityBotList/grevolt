package gateway

import (
	"errors"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

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

type WsState int

const (
	WsStateClosed     WsState = iota
	WsStateOpening    WsState = iota
	WsStateOpen       WsState = iota
	WsStateClosing    WsState = iota
	WsStateRestarting WsState = iota
)

type NotifyEvent struct {
	Data []byte // Event data
	Type string // Event type
}

type NotifyPayload struct {
	OpCode IOpCode        // Internal library OpCode
	Data   map[string]any // Internal Opcode data
	Event  NotifyEvent    // Event data, only applicable to EVENT_IOpCode
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

	// Websocket state
	State WsState

	// This channel is fired when status updates are received
	StatusChannel chan bool

	// Event handlers, set these to handle events
	EventHandlers EventHandlers

	// unique id describing the heartbeat
	heartbeatId string

	// websocket waitgroups
	wg sync.WaitGroup

	// already initted ws
	wsInitOnce bool
}

func (w *GatewayClient) GatewayURL() string {
	gwUrl := w.WSUrl + "?v=" + w.APIVersion + "&encoding=" + w.Encoding

	w.Logger.Debug("gateway url: " + gwUrl)

	return gwUrl
}

// Opens a websocket connection to the gateway
func (w *GatewayClient) Open() error {
	// Ensure there is only one Open() call at a time
	w.Lock()
	defer w.Unlock()

	if w.State == WsStateOpen || w.State == WsStateOpening {
		return ErrWSAlreadyOpen
	}

	w.Logger.Debug("waiting for old connections (if any) to close")
	w.wg.Wait()

	w.State = WsStateOpening

	w.Logger.Debug("opening connection to gateway")

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
	w.NotifyChannel = make(chan *NotifyPayload)

	if !w.wsInitOnce {
		w.StatusChannel = make(chan bool)
	}

	w.heartbeatId = ""
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

	w.WsConn.SetCloseHandler(func(code int, text string) error {
		w.State = WsStateClosed
		w.Logger.Debug("websocket closed: ", code, text)
		w.NotifyChannel <- &NotifyPayload{
			OpCode: ERROR_IOpCode,
		}

		return nil
	})

	w.State = WsStateOpen
	w.wsInitOnce = true

	go w.handleNotify()
	time.Sleep(1 * time.Second)
	go w.readMessages()

	w.Logger.Debug("opened connection to gateway")

	return nil
}

func (w *GatewayClient) Close() {
	w.State = WsStateClosing
	w.NotifyChannel <- &NotifyPayload{
		OpCode: KILL_IOpCode,
	}
}

// Wait for the gateway to close
func (w *GatewayClient) Wait() {
	for payload := range w.StatusChannel {
		w.Logger.Debug("received exit payload: ", payload)
		if payload {
			// Close the websocket
			return
		}
	}
}

// Waits for the gateway to close, then sends a notification to the notify channel
func (w *GatewayClient) OnDone(notify chan bool) {
	<-w.StatusChannel
	notify <- true
}

type internalMessage struct {
	Type  string `json:"type"`
	Error string `json:"error,omitempty"`
}

func (w *GatewayClient) readMessages() {
	w.wg.Add(1)

	defer func() {
		w.Logger.Debug("wg decr (readMessages)")
		w.wg.Done()
	}()

	w.Logger.Debug("readMessages task started")

	w.WsConn.SetReadDeadline(time.Now().Add(w.Deadline))

	// Before doing anything else, send AUTHENTICATE_IOpCode
	w.NotifyChannel <- &NotifyPayload{
		OpCode: AUTHENTICATE_IOpCode,
	}

	for {
		_, message, err := w.WsConn.ReadMessage()

		var data internalMessage

		// If we have a message, try and decode it first, before checking for a close code
		if len(message) > 0 {
			err = w.Decode(message, &data)

			if err != nil {
				w.Logger.Error("failed to unmarshal message: " + err.Error())
				data = internalMessage{
					Type: "InternalError",
				}
			}

			// Before doing anything else, create a event so it can be handled
			w.NotifyChannel <- &NotifyPayload{
				OpCode: EVENT_IOpCode,
				Event: NotifyEvent{
					Type: data.Type,
					Data: message,
				},
			}

			if data.Type == "" {
				w.Logger.Warn("recieved message with empty type")
				continue
			}

			// Error handling here, before checking frames, allows for detection of invalid auth credential errors
			var err bool = true
			switch data.Type {
			case "Pong":
				w.Logger.Debug("recieved pong from gateway")
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
				w.Logger.Debug("received LabelMe")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: ERROR_IOpCode,
					Data: map[string]any{
						"error": "recieved unknown error: label me",
					},
				}
			case "InternalError":
				w.Logger.Debug("received InternalError")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: ERROR_IOpCode,
					Data: map[string]any{
						"error": "recieved unknown error: internal error",
					},
				}
			case "InvalidSession":
				w.Logger.Debug("received InvalidSession")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: FATAL_IOpCode,
					Data: map[string]any{
						"error": "invalid session",
					},
				}
			case "OnboardingNotFinished":
				w.Logger.Debug("received OnboardingNotFinished")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: FATAL_IOpCode,
					Data: map[string]any{
						"error": "onboarding not finished [OnboardingNotFinished]",
					},
				}
			// TODO: rethink this: ERROR vs NOTIFY
			case "AlreadyAuthenticated":
				w.Logger.Debug("received AlreadyAuthenticated")
				w.NotifyChannel <- &NotifyPayload{
					OpCode: ERROR_IOpCode,
					Data: map[string]any{
						"error": "already authenticated [AlreadyAuthenticated]",
					},
				}
			case "Authenticated":
				w.Logger.Debug("received Authenticated flag")
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
			// Check if its a "use of closed network connection" error
			//
			// These are not fatal and should definitely not be spawning a notify payload
			if strings.Contains(err.Error(), "use of closed network connection") {
				w.Logger.Debug("websocket closed, exiting readMessages()")
				return
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				w.Logger.Debug("error: %v", err)
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

			return
		}

		if len(message) == 0 {
			w.Logger.Warn("recieved empty message")
			continue
		}

		// Recieved message successfully, extend deadline
		w.WsConn.SetReadDeadline(time.Now().Add(w.Deadline))
	}
}

func (w *GatewayClient) handleNotify() {
	restarter := func() {
		// If closed, don't restart
		if w.State == WsStateClosed {
			w.Logger.Debug("not restarting connection to gateway because it was killed")
			return
		}

		w.Logger.Debug("restarting connection to gateway")

		// Send restart opcode
		w.Logger.Debug("sending restart opcode")
		w.WsConn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "closeWsConn"), time.Now().Add(w.Deadline))
		w.WsConn.Close()
		w.State = WsStateRestarting

		// Avoid leaking channels
		close(w.NotifyChannel)

		w.Logger.Debug("opening new connection to gateway")
		go w.Open()
	}

	w.wg.Add(1)
	defer func() {
		w.Logger.Debug("wg decr (handleNotify)")
		w.wg.Done()
	}()

	killer := func() {
		w.State = WsStateClosed
		w.heartbeatId = ""
		w.WsConn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "closeWsConn"), time.Now().Add(w.Deadline))
		w.WsConn.Close()
		w.StatusChannel <- true
	}

	for payload := range w.NotifyChannel {
		switch payload.OpCode {
		case KILL_IOpCode:
			w.Logger.Debug("killing connection to gateway")
			killer()
			return
		case RESTART_IOpCode:
			restarter()
			return
		case AUTHENTICATE_IOpCode:
			// Send authenticate command frame
			w.Logger.Debug("sending authenticate command frame")
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
			killer()
			return
		case EVENT_IOpCode:
			go w.HandleEvent(payload.Event.Data, payload.Event.Type)
		}
	}

	if w.State == WsStateOpen {
		restarter()
	}
}

func (w *GatewayClient) heartbeat(hbId string) {
	w.Logger.Debug("starting heartbeat, ", hbId)
	// Create new ticker
	ticker := time.NewTicker(w.HeartbeatInterval)

	// Send heartbeat
	for range ticker.C {
		if w.heartbeatId != hbId {
			w.Logger.Debug("heartbeat id mismatch, this is normal especially due to restarts")
			ticker.Stop()
			return // Not the current heartbeat ws
		}

		if w.State != WsStateOpen {
			continue
		}

		w.Logger.Debug("sending heartbeat")
		w.Send(map[string]any{
			"type": "Ping",
			"data": time.Now().Unix(),
		})
	}
}
