package configs

import (
	"fmt"
	"testing"
)

func TestMessagebusConfig(t *testing.T) {
	mb, err := NewMessagebusConfig("obj/messagebus_cfg.txt")
	if err != nil {
		panic("create messagebus config failed")
	}
	data := mb.GetMessagebusData()
	fmt.Printf("%#v\n", data)
}

func TestMysqlDbConfig(t *testing.T) {
	md, err := NewMysqlDbConfig("obj/mysql_cfg.txt")
	if err != nil {
		panic("mysql db config config error")
	}
	data := md.GetMysqlDbData()
	fmt.Printf("%#v\n", data)
}
