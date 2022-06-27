package gb28181

import (
	"fmt"
	"strconv"
	"sync"
	"test/models"

	"gorm.io/gorm"
)

// DeviceStream DeviceStream
type DeviceStream struct {
	// 0  直播 1 历史
	T          int
	SSRC       string
	DeviceID   string
	UserID     string
	StreamType string            //  pull 媒体服务器主动拉流，push 监控设备主动推流
	Status     int               // 0正常 1关闭 -1 尚未开始
	Ftag       map[string]string // header from params
	Ttag       map[string]string // header to params
	CallID     string            // header callid
	Time       string
	Stop       bool
}

type playList struct {
	// key=ssrc value=PlayParams  播放对应的PlayParams 用来发送bye获取tag，callid等数据
	ssrcResponse *sync.Map
	// key=deviceid value={ssrc,path}  当前设备直播信息，防止重复直播
	devicesSucc *sync.Map
	ssrc        int
}

var _playList playList

func getSSRC(t int) string {
	r := false
	for {
		_playList.ssrc++
		key := fmt.Sprintf("%d%s%04d", t, gbConfig.GB28181.Region[3:8], _playList.ssrc)
		// 数据可查询
		if stream, err := models.GetStreamWithSSRC(ssrc2stream(key)); err == gorm.ErrRecordNotFound || stream.SSRC == "" {
			return key
		}
		if _playList.ssrc > 9000 && !r {
			_playList.ssrc = 0
			r = true
		}
	}
}

// zlm接收到的ssrc为16进制。发起请求的ssrc为10进制
func ssrc2stream(ssrc string) string {
	if ssrc[0:1] == "0" {
		ssrc = ssrc[1:]
	}
	num, _ := strconv.Atoi(ssrc)
	return fmt.Sprintf("%X", num)
}
