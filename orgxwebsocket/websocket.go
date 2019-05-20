package orgxwebsocket

import (
	"golang.org/x/net/websocket"
)

type CReadMessage struct {
	Length int
	Body   []byte
}

type CWebsocket struct {
	conn     *websocket.Conn
	readChan chan *CReadMessage
}

func (this *CWebsocket) init() {
	go func() {
		for {
			body := []byte{}
			readLen, err := this.conn.Read(body)
			if err != nil {
				continue
			}
			message := CReadMessage{
				Length: readLen,
				Body:   body,
			}
			this.readChan <- &message
		}
	}()
}

func (this *CWebsocket) CloseReadChannel() {
	if _, ok := this.readChan; ok {
		close(this.readChan)
	}
}

func (this *CWebsocket) Read() <-chan *CReadMessage {
	return this.readChan
}

func New(conn *websocket.Conn) *CWebsocket {
	ws := CWebsocket{
		conn:     conn,
		readChan: make(chan *CReadMessage),
	}
	ws.init()
	return &ws
}
