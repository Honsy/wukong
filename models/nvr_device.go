package models

import (
	"test/pkg/sip"

	"gorm.io/gorm"
)

// NVR 对象模型
type Device struct {
	Model
	Name       string       `json:"name"`
	DeviceId   string       `json:"device_id"`
	Region     string       `json:"region"`
	Host       string       `json:"host"`
	Port       string       `json:"port"`
	Proto      string       `json:"proto"`
	DeviceType string       `json:"device_type"`
	ModelType  string       `json:"model_type"`
	Active     int          `json:"active"`
	Regist     int          `json:"regist"`
	Rport      string       `json:"rport"`
	RAddr      string       `json:"raddr"`
	TransPort  string       `json:"transport"`
	URIStr     string       `json:"uristr"`
	Addr       *sip.Address `gorm:"-"`
}

// Devices 摄像头信息
type Camera struct {
	Model
	// DeviceID 设备编号
	DeviceID string `xml:"DeviceID" json:"device_id"`
	// Name 设备名称
	Name         string `xml:"Name" json:"name"`
	Manufacturer string `xml:"Manufacturer" json:"manufacturer"`
	Owner        string `xml:"Owner" json:"owner"`
	CivilCode    string `xml:"CivilCode" json:"civil_code"`
	// Address ip地址
	Address     string `xml:"Address" json:"address"`
	Parental    int    `xml:"Parental" json:"parental"`
	SafetyWay   int    `xml:"SafetyWay" json:"safety_way"`
	RegisterWay int    `xml:"RegisterWay" json:"registerway"`
	Secrecy     int    `xml:"Secrecy" json:"secrecy"`
	// Status 状态  on 在线
	Status string `xml:"Status" json:"status"`
	// PDID 所属用户id
	PDID string `json:"pd_id"`
	// Active 最后活跃时间
	Active int64  `json:"active"`
	URIStr string `json:"uri"`

	addr *sip.Address `gorm:"-"`
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

// 查询NVR设备下所有的子设备
func GetCamerasWithDeviceId(deviceId string) ([]*Camera, error) {
	var cameras []*Camera
	err := db.Debug().Select("*").Where("device_id = ?", deviceId).Find(&cameras).Error

	if err != nil {
		return nil, err
	}

	return cameras, nil
}

// 根据DeivceId查询单个子设备
func GetCamera(deviceId string) (Camera, error) {
	var camera Camera
	err := db.Debug().Select("*").Where("device_id = ?", deviceId).First(&camera).Error

	if err != nil {
		return Camera{}, err
	}

	return camera, nil
}

// 更新摄像头
func UpdateCamera(id int, data interface{}) error {
	if err := db.Debug().Model(&Camera{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// 插入摄像头
func InsertCamera(camera Camera) error {
	if err := db.Create(&camera).Error; err != nil {
		return err
	}
	return nil
}
