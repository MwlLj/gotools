package orgxwebsocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type CReadMessage struct {
	MessageType int
	Body        *[]byte
}

type CWebsocket struct {
	conn     *websocket.Conn
	readChan chan *CReadMessage
}

func (this *CWebsocket) init() {
	go func() {
		for {
			messageType, body, err := this.conn.ReadMessage()
			if err != nil {
				close(this.readChan)
				break
			}
			message := CReadMessage{
				MessageType: messageType,
				Body:        &body,
			}
			this.readChan <- &message
		}
	}()
}

func (this *CWebsocket) Read() <-chan *CReadMessage {
	return this.readChan
}

func (this *CWebsocket) Write(messageType int, body *[]byte) error {
	return this.conn.WriteMessage(messageType, *body)
}

func New(w http.ResponseWriter, r *http.Request) *CWebsocket {
	conn, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		return nil
	}
	ws := CWebsocket{
		conn:     conn,
		readChan: make(chan *CReadMessage),
	}
	ws.init()
	return &ws
}
