package configs

import (
	"encoding/json"
)

type CMysqlDbData struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Username string `json:"username"`
	Userpwd  string `json:"userpwd"`
	Dbname   string `json:"dbname"`
}

type CMysqlDbConfig interface {
	GetMysqlDbData() *CMysqlDbData
}

type mysqlDbConfig struct {
	CConfigBase
	m_data *CMysqlDbData
}

func NewMysqlDbConfig(path string) (CMysqlDbConfig, error) {
	data, err := json.Marshal(&CMysqlDbData{"localhost", 3306, "root", "123456", ""})
	if err != nil {
		return nil, err
	}
	mb := &mysqlDbConfig{m_data: &CMysqlDbData{}}
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

func (this *mysqlDbConfig) GetMysqlDbData() *CMysqlDbData {
	return this.m_data
}
