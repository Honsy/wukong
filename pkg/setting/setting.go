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

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("app.ini")

	if err != nil {
		log.Fatalf("setting read error")
	}

	cfg.Section("server").MapTo(ServerSetting)
	cfg.Section("database").MapTo(DatabaseSetting)
}
