package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"test/models"
	"test/pkg/gb28181"
	"test/pkg/onvif"
	"test/pkg/setting"
	"test/routers"
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

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()

	log.Fatal(err)
}
