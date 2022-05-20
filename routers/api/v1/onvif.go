package v1

import (
	"net/http"
	"test/pkg/app"
	"test/pkg/enum"
	"test/pkg/onvif"

	"github.com/gin-gonic/gin"
)

func GetOnvifRtsp(c *gin.Context) {
	appG := app.Gin{C: c}

	res, err := onvif.GetStreamUri(onvif.StreamOptions{})
	if err != nil {
		appG.Response(http.StatusInternalServerError, enum.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, enum.SUCCESS, map[string]interface{}{
		"url": res,
	})

}

func GetOnvifDevice(c *gin.Context) {

}
