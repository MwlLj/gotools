package logclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	requestModeConnect       string = "connect"
	requestModeSending       string = "sending"
	requestIdentifyPublish   string = "publish"
	requestIdentifySubscribe string = "subscribe"
	storageModeNone          string = "none"
	storageModeFile          string = "file"
)

var (
	logTypeMessage string = "message"
	logTypeDebug   string = "debug"
	logTypeInfo    string = "info"
	logTypeWarn    string = "warn"
	logTypeError   string = "error"
	logTypeFatal   string = "fatal"
)

type CClient struct {
	host          string
	port          int
	serverName    string
	serverVersion string
	serverNo      string
	ch            chan *CRequest
}

func (this *CClient) Println(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	this.sendRequest("", &logTypeMessage, &s)
}

func (this *CClient) Debugln(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	this.sendRequest("", &logTypeDebug, &s)
}

func (this *CClient) Infoln(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	this.sendRequest("", &logTypeInfo, &s)
}

func (this *CClient) Warnln(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	this.sendRequest("", &logTypeWarn, &s)
}

func (this *CClient) Errorln(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	this.sendRequest("", &logTypeError, &s)
}

func (this *CClient) Fatalln(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	this.sendRequest("", &logTypeFatal, &s)
}

func (this *CClient) init() {
	this.ch = make(chan *CRequest)
	go func() {
		isConnect := false
		var conn *net.TCPConn = nil
		for {
			request := <-this.ch
			if !isConnect {
				conn, isConnect = this.connect()
			}
			if !isConnect {
				time.Sleep(1 * time.Second)
				continue
			}
			err := this.send(conn, request)
			if err != nil {
				isConnect = false
			}
		}
	}()
}

func (this *CClient) connect() (*net.TCPConn, bool) {
	buffer := bytes.Buffer{}
	buffer.WriteString(this.host)
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatInt(int64(this.port), 10))
	addr, err := net.ResolveTCPAddr("tcp4", buffer.String())
	if err != nil {
		return nil, false
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, false
	}
	request := CRequest{
		Mode:          requestModeConnect,
		Identify:      requestIdentifyPublish,
		ServerName:    this.serverName,
		ServerVersion: this.serverVersion,
		ServerNo:      this.serverNo,
	}
	err = this.send(conn, &request)
	if err != nil {
		return nil, false
	}
	return conn, true
}

func (this *CClient) sendRequest(topic string, logType *string, data *string) {
	request := CRequest{
		Mode:          requestModeSending,
		Identify:      requestIdentifyPublish,
		ServerName:    this.serverName,
		ServerVersion: this.serverVersion,
		ServerNo:      this.serverNo,
		Topic:         topic,
		Data:          *data,
		StorageMode:   storageModeFile,
		LogType:       *logType,
	}
	this.ch <- &request
}

func (this *CClient) send(conn *net.TCPConn, request *CRequest) error {
	b, err := json.Marshal(request)
	if err != nil {
		log.Printf("encode request error, err: %v", err)
		return err
	}
	buf := bytes.Buffer{}
	buf.WriteString(string(b))
	buf.WriteString("\n")
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		log.Printf("send error, err: %v", err)
		return err
	}
	return nil
}
