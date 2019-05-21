package websocket

import (
	"golang.org/x/net/websocket"
)

type CReadMessage struct {
	Length int
	Body   []byte
}

type CWebsocket struct {
	conn          *websocket.Conn
	readChan      chan *CReadMessage
	bodyMaxLength int
}

func (this *CWebsocket) init() {
	go func() {
		body := make([]byte, this.bodyMaxLength)
		for {
			readLen, err := this.conn.Read(body)
			if err != nil {
				this.readChan <- nil
				close(this.readChan)
				break
			}
			message := CReadMessage{
				Length: readLen,
				Body:   body,
			}
			this.readChan <- &message
			body = make([]byte, this.bodyMaxLength)
		}
	}()
}

func (this *CWebsocket) Read() <-chan *CReadMessage {
	return this.readChan
}

func (this *CWebsocket) Write(body []byte) {
	this.conn.Write(body)
}

func (this *CWebsocket) Close() {
	this.conn.Close()
}

func New(conn *websocket.Conn, bodyMaxLength int) *CWebsocket {
	ws := CWebsocket{
		conn:          conn,
		readChan:      make(chan *CReadMessage, 1),
		bodyMaxLength: bodyMaxLength,
	}
	ws.init()
	return &ws
}
