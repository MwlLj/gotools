package configs

import (
	"encoding/json"
)

type CMessagebusData struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Userpwd  string `json:"userpwd"`
}

type CMessagebusConfig interface {
	GetMessagebusData() *CMessagebusData
}

type messagebusConfig struct {
	CConfigBase
	m_data *CMessagebusData
}

func NewMessagebusConfig(path string) (CMessagebusConfig, error) {
	data, err := json.Marshal(&CMessagebusData{"127.0.0.1", 1883, "", ""})
	if err != nil {
		return nil, err
	}
	mb := &messagebusConfig{m_data: &CMessagebusData{}}
	buf, e := mb.Load(path, string(data))
	if e != nil {
		return nil, err
	}
	e = json.Unmarshal([]byte(buf), mb.m_data)
	if e != nil {
		return nil, err
	}
	return mb, nil
}

func (this *messagebusConfig) GetMessagebusData() *CMessagebusData {
	return this.m_data
}
