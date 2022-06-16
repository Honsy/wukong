package gb28181

import (
	"net"
	"test/pkg/setting"
)

// Config Config
type Config struct {
	GB28181            SysInfo `json:"gb28181"`
	mediaServerRtpIP   net.IP
	mediaServerRtpPort int
	Media              setting.Media
}
type SysInfo struct {
	// Region 当前域
	Region string `json:"region"`
	// UID 用户id固定头部
	UID string `json:"uid"`
	// UNUM 当前用户数
	UNUM int `json:"unum"`
	// DID 设备id固定头部
	DID string `json:"did"`
	// DNUM 当前设备数
	DNUM int `json:"dnum"`
	// LID 当前服务id
	LID string `json:"lid"`
	// Port
	Port string `json:"port"`
	// Ip
	IP string `json:"ip"`
}
