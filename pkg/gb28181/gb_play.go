package gb28181

import (
	"test/pkg/sip"
	"time"
)

// playParams 播放请求参数
type playParams struct {
	// 0  直播 1 历史
	T int
	//  开始结束时间，只有t=1 时有效
	S, E       time.Time
	SSRC       string
	Resp       *sip.Response
	DeviceID   string
	UserID     string
	ext        int64  // 推流等待的过期时间，用于判断是否请求成功但推流失败。超过还未接收到推流定义为失败，重新请求推流或者关闭此ssrc
	stream     bool   // 是否完成推流，用于web_hook 出现stream=false时等待推流，出现stream_not_found 且 stream=true表示推流过但已关闭。释放ssrc。
	streamType string //  pull 媒体服务器主动拉流，push 监控设备主动推流
}

// 请求播放
func GBPlay() interface{} {

	return nil
}

// 推送Media RTP Server
func gbPlayPush(data playParams) {
	name := "Play"
	protocal := "TCP/RTP/AVP"
	if data.T == 1 {
		name = "Playback"
		protocal = "RTP/RTCP"
	}
}
