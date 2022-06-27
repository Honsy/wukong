package models

type Stream struct {
	Model
	Type       int    `json:"type"`        // 播放类型
	SSRC       string `json:"ssrc"`        // rtp标志位
	CameraId   string `json:"camera_id"`   // 摄像头ID
	DeviceId   string `json:"device_id"`   // 用户设备ID(NVR, GB摄像头)
	StreamType string `json:"stream_type"` // 流类型
	Status     int    `json:"status"`      // 状态
	Time       string `json:"time"`
}

// 插入流
func InsertStream(stream Stream) error {
	if err := db.Create(&stream).Error; err != nil {
		return err
	}
	return nil
}

// 更新流
func UpdateStream(id int, data interface{}) error {
	if err := db.Debug().Model(&Stream{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// 根据id查询流
func GetStreamWithSSRC(ssrc string) (Stream, error) {
	var stream Stream
	err := db.Debug().Select("*").Where("ssrc = ?", ssrc).First(&stream).Error

	if err != nil {
		return Stream{}, err
	}

	return stream, nil
}
