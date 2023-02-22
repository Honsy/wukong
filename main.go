package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"test/models"
	"test/pkg/logging"
	"test/pkg/setting"
	"test/pkg/sip/sipserver"
	"test/routers"
)

var (
	server sipserver.Server
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	// gb28181.Setup()
	// onvif
	// onvif.Setup(onvif.CameraOption{
	// 	Hostname: "10.0.16.111",
	// 	Port:     80,
	// 	Username: "admin",
	// 	Passowrd: "l1234567",
	// })

	routersInit := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	logging.Warn("start http server listening %s", endPoint)

	routersInit.Run(endPoint)
}
