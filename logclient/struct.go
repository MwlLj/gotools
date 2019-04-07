package logclient

type CRequest struct {
	Mode          string `json:"mode"`
	Identify      string `json:"identify"`
	ServerName    string `json:"serverName"`
	ServerVersion string `json:"serverVersion"`
	ServerNo      string `json:"serverNo"`
	Topic         string `json:"topic"`
	Data          string `json:"data"`
	StorageMode   string `json:"storageMode"`
	LogType       string `json:"logType"`
}
