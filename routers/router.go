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

	root := r.Group("/wukong")
	root.POST("/auth", api.GetAuth)
	// V1版本
	apiv1 := root.Group("/api/v1")
	// onvif
	onvif := apiv1.Group("/onvif")
	{
		onvif.GET("/device", v1.GetOnvifDevice)
		onvif.GET("/getRtsp", v1.GetOnvifRtsp)
	}
	// 用户
	user := apiv1.Group("/user")
	{
		user.POST("/login", v1.Login)
		user.POST("/register", v1.Register)
	}
	return r
}
