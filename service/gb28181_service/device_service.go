package gb28181service

import (
	"test/models"
	"test/service"
)

type DeviceService struct {
	service.PageCommon
	DeviceId string
	maps     interface{}
}

func (u *DeviceService) GetDeviceList() ([]*models.Device, error) {
	return models.GetDeivces(u.PageNumber, u.PageSize, u.getMaps())
}

func (u *DeviceService) GetCamerasWithDeivceId() ([]*models.Camera, error) {
	return models.GetCamerasWithDeviceId(u.DeviceId)
}

func (u *DeviceService) Count() (int64, error) {
	return models.GetDevicesTotal()
}

func (a *DeviceService) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	return maps
}
