package v1

import (
	"net/http"
	"test/pkg/app"
	"test/pkg/enum"
	"test/pkg/rtsp"

	"github.com/gin-gonic/gin"
)

func TestRtspUrl(c *gin.Context) {
	appG := app.Gin{C: c}

	rtspUrl := c.Query("url")
	rtspClient := rtsp.NewRtspClient(rtspUrl)

	rtspClient.ContinueAfterClientCreation1()

	appG.Response(http.StatusOK, enum.SUCCESS, map[string]interface{}{
		"url": "",
	})

}
