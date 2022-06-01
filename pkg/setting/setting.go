package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

type Server struct {
	RunMode  string
	HttpPort int
}

type GB28181 struct {
	Lid    string
	Region string
	Uid    string
	Did    string
	Unum   int
	Dnum   int
}

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var GBSetting = &GB28181{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("app.ini")

	if err != nil {
		log.Fatalf("setting read error")
	}

	cfg.Section("server").MapTo(ServerSetting)
	cfg.Section("database").MapTo(DatabaseSetting)
	cfg.Section("gb28181").MapTo(GBSetting)
}
