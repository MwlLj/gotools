package logclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
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
	isConnect     bool
	mutex         sync.Mutex
	conn          *net.TCPConn
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
	this.connect()
	go func(self *CClient) {
		for {
			if !self.isConnect {
				self.connect()
			}
			time.Sleep(3 * time.Second)
		}
	}(this)
}

func (this *CClient) connect() {
	buffer := bytes.Buffer{}
	buffer.WriteString(this.host)
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatInt(int64(this.port), 10))
	addr, err := net.ResolveTCPAddr("tcp4", buffer.String())
	if err != nil {
		this.mutex.Lock()
		this.isConnect = false
		this.mutex.Unlock()
		return
	}
	this.conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		this.mutex.Lock()
		this.isConnect = false
		this.mutex.Unlock()
		return
	}
	this.conn.SetKeepAlive(true)
	request := CRequest{
		Mode:          requestModeConnect,
		Identify:      requestIdentifyPublish,
		ServerName:    this.serverName,
		ServerVersion: this.serverVersion,
		ServerNo:      this.serverNo,
	}
	err = this.send(&request)
	if err != nil {
		this.mutex.Lock()
		this.isConnect = false
		this.mutex.Unlock()
		return
	}
	this.mutex.Lock()
	this.isConnect = true
	this.mutex.Unlock()
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
	this.send(&request)
}

func (this *CClient) send(request *CRequest) error {
	if this.conn == nil {
		this.mutex.Lock()
		this.isConnect = false
		this.mutex.Unlock()
		log.Println("connect is nil", this.isConnect)
		return errors.New("not connect")
	}
	b, err := json.Marshal(request)
	if err != nil {
		this.mutex.Lock()
		this.isConnect = false
		this.mutex.Unlock()
		log.Printf("encode request error, err: %v", err)
		return err
	}
	buf := bytes.Buffer{}
	buf.WriteString(string(b))
	buf.WriteString("\n")
	_, err = io.WriteString(this.conn, buf.String())
	// _, err = this.conn.Write(buf.Bytes())
	if err != nil {
		this.mutex.Lock()
		this.isConnect = false
		this.mutex.Unlock()
		log.Printf("send error, err: %v", err)
		return err
	}
	this.mutex.Lock()
	this.isConnect = true
	this.mutex.Unlock()
	return nil
}
