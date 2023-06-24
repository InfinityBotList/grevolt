package gateway

import (
	"errors"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/infinitybotlist/grevolt/client/auth"
	"github.com/infinitybotlist/grevolt/gateway/broadcast"
	"go.uber.org/zap"
)

// ErrWSAlreadyOpen is thrown when you attempt to open
// a websocket that already is open.
var ErrWSAlreadyOpen = errors.New("web socket already opened")

// Internal IOpCodes allow for control handling of the WS
//
// This is the primitive for how reading and writing is synchronized
type IOpCode int

const (
	// Kill the websocket
	KILL_IOpCode IOpCode = iota

	// Restart the websocket
	RESTART_IOpCode IOpCode = iota

	// Authenticate the websocket
	AUTHENTICATE_IOpCode IOpCode = iota

	// An error has occured, this also triggers a restart after
	// logging the error
	ERROR_IOpCode IOpCode = iota

	// A fatal error has occured, this kills the websocket
	// after logging the error
	FATAL_IOpCode IOpCode = iota

	// Send an event to HandleEvent to be handled
	EVENT_IOpCode IOpCode = iota
)

// Websocket state, you can use this to monitor the state of the websocket
//
// +unstable
type WsState int

const (
	// The websocket is closed
	WsStateClosed WsState = iota

	// The websocket is opening
	WsStateOpening WsState = iota

	// The websocket is open and receiving events
	WsStateOpen WsState = iota

	// The websocket is closing
	WsStateClosing WsState = iota

	// The websocket is restarting
	WsStateRestarting WsState = iota
)

// NotifyEvents are low level events that are fired when the WS receives a eent
// to be fired. This is highly unstable, you likely want to use the EventHandlers
// instead.
//
// +unstable
type NotifyEvent struct {
	Data []byte // Event data
	Type string // Event type
}

// NotifyPayload is a payload that is sent to the NotifyChannel
//
// This is used to control the WS, and to send events to the EventHandlers. It is
// also highly unstable, you likely want to use the EventHandlers instead.
//
// +unstable
type NotifyPayload struct {
	OpCode IOpCode        // Internal library OpCode
	Data   map[string]any // Internal Opcode data
	Event  NotifyEvent    // Event data, only applicable to EVENT_IOpCode
}

// StatusMessage is a message that is sent to the StatusChannel
// to allow listening to status updates to the websocket
type StatusMessage int

const (
	// DONE is sent when the websocket is done and all listeners
	// should stop listening and end.
	DONE_StatusMessage StatusMessage = iota

	// WSEND is sent when the websocket has closed (restarts etc can cause this)
	//
	// Heartbeaters should listen for this and stop sending heartbeats when
	// this is received.
	WSEND_StatusMessage StatusMessage = iota
)

// Status payload, can be used functions such as Wait() etc
//
// +unstable
type StatusPayload struct {
	StatusMessage StatusMessage
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
	SessionToken *auth.Token

	// Websocker payload format, either json or msgpack (only these are supported by the client)
	Encoding string

	// The websocket connection
	//
	// +unstable
	WsConn *websocket.Conn

	// Notifications from the WS
	//
	// This is very low level and should not be used unless you know what you are doing
	//
	// +unstable
	NotifyChannel chan *NotifyPayload

	// Websocket state
	//
	// +unstable
	State WsState

	// This channel is fired when status updates are received
	//
	// +unstable
	StatusChannel broadcast.BroadcastServer[*StatusPayload]

	// Event handlers, set these to handle events
	EventHandlers EventHandlers

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

	if !w.StatusChannel.Open {
		w.StatusChannel = broadcast.NewBroadcastServer[*StatusPayload](w.Logger)
	}

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
	sub := w.StatusChannel.Subscribe()
	for payload := range sub {
		w.Logger.Debug("received statusChannel payload: ", payload)

		if payload == nil {
			continue
		}

		if payload.StatusMessage == DONE_StatusMessage {
			// Close the websocket
			w.StatusChannel.CancelSubscription(sub)
			return
		}
	}
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
				go w.heartbeat()
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

		w.StatusChannel.Broadcast(&StatusPayload{
			StatusMessage: WSEND_StatusMessage,
		})

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
		w.WsConn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "closeWsConn"), time.Now().Add(w.Deadline))
		w.WsConn.Close()
		w.StatusChannel.Broadcast(&StatusPayload{
			StatusMessage: DONE_StatusMessage,
		})
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

func (w *GatewayClient) heartbeat() {
	hbStartTime := time.Now().Nanosecond()
	w.Logger.Debug("starting heartbeat ", hbStartTime)
	// Create new ticker
	ticker := time.NewTicker(w.HeartbeatInterval)

	// Send heartbeat
	sub := w.StatusChannel.Subscribe()

	defer func() {
		ticker.Stop()
		w.StatusChannel.CancelSubscription(sub)
	}()

	for {
		select {
		case p := <-sub:
			if p == nil {
				// The status channel has closed, we should also die
				w.Logger.Debug("status channel closed, exiting heartbeat")
				return
			}

			if p.StatusMessage == DONE_StatusMessage || p.StatusMessage == WSEND_StatusMessage {
				w.Logger.Debug("stopping heartbeat")
				return
			}

		case <-ticker.C:
			if w.State != WsStateOpen {
				continue
			}

			w.Logger.Debug("sending heartbeat", hbStartTime)
			w.Send(map[string]any{
				"type": "Ping",
				"data": time.Now().Unix(),
			})
		}
	}
}
