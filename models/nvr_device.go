package models

import (
	"test/pkg/sip"

	"gorm.io/gorm"
)

// NVR 对象模型
type Device struct {
	Model
	Name       string  `json:"name"`
	DeviceId   string  `json:"device_id"`
	Region     string  `json:"region"`
	Host       string  `json:"host"`
	Port       string  `json:"port"`
	Proto      string  `json:"proto"`
	DeviceType string  `json:"device_type"`
	ModelType  string  `json:"model_type"`
	Active     int     `json:"active"`
	Regist     int     `json:"regist"`
	Rport      string  `json:"rport"`
	RAddr      string  `json:"raddr"`
	TransPort  string  `json:"transport"`
	URIStr     string  `json:"uristr"`
	Addr       sip.Uri `gorm:"-"`
}

// 查询Device根据deviceid
func GetDeviceByDeviceId(id string) (Device, error) {
	var nvr_device Device
	err := db.Debug().Select("*").Where("device_id = ?", id).First(&nvr_device).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Device{}, err
	}

	return nvr_device, nil
}

// 增加Device
func AddDevice(data map[string]interface{}) error {
	device := Device{
		Name:       data["name"].(string),
		DeviceId:   data["device_id"].(string),
		Region:     data["region"].(string),
		Host:       data["host"].(string),
		Port:       data["port"].(string),
		Proto:      data["proto"].(string),
		DeviceType: data["device_type"].(string),
		ModelType:  data["model_type"].(string),
		Active:     data["active"].(int),
		Regist:     data["regist"].(int),
	}
	if err := db.Create(&device).Error; err != nil {
		return err
	}

	return nil
}

// 查询所有的设备列表
func GetDeivces(pageNum int, pageSize int, maps interface{}) ([]*Device, error) {
	var devices []*Device

	err := db.Debug().Where(maps).Offset(pageNum - 1).Limit(pageSize).Find(&devices).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return devices, nil
}

// 查询设备列表的总数
func GetDevicesTotal() (int64, error) {
	var count int64
	if err := db.Debug().Model(&Device{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
