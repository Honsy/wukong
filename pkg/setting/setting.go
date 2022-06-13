package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

// 数据库配置
type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

// 服务器配置
type Server struct {
	RunMode  string
	HttpPort int
}

// GB配置
type GB28181 struct {
	Ip     string
	Port   string
	Lid    string
	Region string
	Uid    string
	Did    string
	Unum   int
	Dnum   int
}

// Media服务器配置
type Media struct {
	Restful string
	Http    string
	WS      string
	Rtmp    string
	Rtsp    string
	Rtp     string
	Secret  string
}

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var GBSetting = &GB28181{}
var MediaSetting = &Media{}

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
	cfg.Section("media").MapTo(MediaSetting)
}
