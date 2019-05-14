package logclient

type IClient interface {
	Println(format string, a ...interface{})
	Debugln(format string, a ...interface{})
	Infoln(format string, a ...interface{})
	Warnln(format string, a ...interface{})
	Errorln(format string, a ...interface{})
	Fatalln(format string, a ...interface{})
}

func New(host string, port int, serverName string, serverVersion string, serverNo string) CClient {
	client := CClient{
		host:          host,
		port:          port,
		serverName:    serverName,
		serverVersion: serverVersion,
		serverNo:      serverNo,
	}
	client.init()
	return client
}
