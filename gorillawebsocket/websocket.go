package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	TextMessage int = 1
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
		this.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
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

func (this *CWebsocket) Close() {
	this.conn.Close()
}

func New(w http.ResponseWriter, r *http.Request) *CWebsocket {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
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
