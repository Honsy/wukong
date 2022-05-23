package routers

import (
	"test/log"
	"test/routers/api"
	v1 "test/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	// 自定义Gin Logger
	r.Use(log.GinLogger())
	r.Use(gin.Recovery())
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/onvif", v1.GetOnvifDevice)
		apiv1.GET("/getOnvifRtsp", v1.GetOnvifRtsp)
	}
	return r
}
