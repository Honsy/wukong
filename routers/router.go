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
		onvif := apiv1.Group("/onvif")
		{
			onvif.GET("/device", v1.GetOnvifDevice)
			onvif.GET("/getRtsp", v1.GetOnvifRtsp)
		}
		user := apiv1.Group("/user")
		{
			user.POST("/login", v1.Login)
			user.POST("/register", v1.Register)
		}
	}
	return r
}
