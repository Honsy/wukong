package v1

import (
	"net/http"
	"test/pkg/app"
	"test/pkg/enum"
	"test/pkg/logging"
	"test/service"
	gb28181service "test/service/gb28181_service"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取所有支持GB的设备（NVR，DVR，支持GB的摄像头）
func GetDeviceList(c *gin.Context) {
	appG := app.Gin{C: c}
	ds := gb28181service.DeviceService{
		PageCommon: service.PageCommon{
			PageNumber: com.StrTo(c.Query("pageNumber")).MustInt(),
			PageSize:   com.StrTo(c.Query("pageSize")).MustInt(),
		},
	}

	total, err := ds.Count()
	if err != nil {
		logging.Error("获取设备列表总数失败！", err)
		appG.Response(http.StatusInternalServerError, enum.ERROR_GET_DEVICE_LIST, nil)
		return
	}

	devices, err := ds.GetDeviceList()

	if err != nil {
		logging.Error("获取设备列表失败！", err)
		appG.Response(http.StatusInternalServerError, enum.ERROR_GET_DEVICE_LIST, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = devices
	data["total"] = total

	appG.Response(http.StatusOK, enum.SUCCESS, data)
}

// 查询设备下所有的子设备
func GetCamerasWithDeivceId(c *gin.Context) {
	appG := app.Gin{C: c}
	ds := gb28181service.DeviceService{
		DeviceId: c.Query("deviceId"),
	}

	cameras, err := ds.GetCamerasWithDeivceId()

	if err != nil {
		logging.Error("获取子设备列表失败！", err)
		appG.Response(http.StatusInternalServerError, enum.ERROR_GET_DEVICE_LIST, nil)
		return
	}
	data := make(map[string]interface{})
	data["list"] = cameras

	appG.Response(http.StatusOK, enum.SUCCESS, data)
}
