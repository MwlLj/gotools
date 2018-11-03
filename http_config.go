package configs

import (
	"encoding/json"
)

type CHttpConfigData struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	DomainName string `json:"domainname"`
}

type CHttpConfig interface {
	GetHttpConfigData() *CHttpConfigData
}

type httpConfig struct {
	CConfigBase
	m_data *CHttpConfigData
}

func NewHttpConfig(path string) (CHttpConfig, error) {
	data, err := json.Marshal(&CHttpConfigData{"0.0.0.0", 8080, ""})
	if err != nil {
		return nil, err
	}
	hp := &httpConfig{m_data: &CHttpConfigData{}}
	buf, e := hp.Load(path, string(data))
	if e != nil {
		return nil, err
	}
	e = json.Unmarshal([]byte(buf), hp.m_data)
	if e != nil {
		return nil, err
	}
	return hp, nil
}

func (this *httpConfig) GetHttpConfigData() *CHttpConfigData {
	return this.m_data
}
