package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"test/models"
	"test/pkg/gb28181"
	"test/pkg/logging"
	"test/pkg/onvif"
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
	gb28181.Setup()
	// onvif
	onvif.Setup(onvif.CameraOption{
		Hostname: "10.0.30.33",
		Port:     80,
		Username: "admin",
		Passowrd: "admin1234",
	})

	routersInit := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logging.Warn("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()

	logging.Fatal(err)
}
