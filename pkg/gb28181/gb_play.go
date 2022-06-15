package gb28181

import (
	"fmt"
	"net/http"
	"test/models"
	"test/pkg/logging"
	"test/pkg/sdp"
	"test/pkg/sip"
	"test/pkg/sip/parser"
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
func gbPlayPush(data playParams, camera models.Camera, device models.Device) (playParams, error) {
	var (
		sdpSession sdp.Session
		byteData   []byte
	)
	name := "Play"
	protocal := "TCP/RTP/AVP"
	if data.T == 1 {
		name = "Playback"
		protocal = "RTP/RTCP"
	}

	// body体
	video := sdp.Media{
		Description: sdp.MediaDescription{
			Type:     "video",
			Port:     gbConfig.mediaServerRtpPort,
			Formats:  []string{"96", "98", "97"},
			Protocol: protocal,
		},
	}
	video.AddAttribute("recvonly")
	if data.T == 0 {
		video.AddAttribute("setup", "passive")
		video.AddAttribute("connection", "new")
	}
	video.AddAttribute("rtpmap", "96", "PS/90000")
	video.AddAttribute("rtpmap", "98", "H264/90000")
	video.AddAttribute("rtpmap", "97", "MPEG4/90000")

	// sdp消息体
	sdpMessage := &sdp.Message{
		Origin: sdp.Origin{
			Username: serverDevices.DeviceId, // 媒体服务器id
			Address:  gbConfig.mediaServerRtpIP.String(),
		},
		Name: name,
		Connection: sdp.ConnectionData{
			IP:  gbConfig.mediaServerRtpIP,
			TTL: 0,
		},
		Timing: []sdp.Timing{
			{
				Start: data.S,
				End:   data.E,
			},
		},
		Medias: []sdp.Media{video},
	}
	if data.T == 1 {
		sdpMessage.URI = fmt.Sprintf("%s:0", data.DeviceID)
	}

	// appending message to session
	sdpSession = sdpMessage.Append(sdpSession)
	// appending session to byte buffer
	byteData = sdpSession.AppendTo(byteData)

	deviceURI, _ := parser.ParseUri(device.URIStr)
	device.Addr = &sip.Address{Uri: deviceURI}

	hdrs := make([]sip.Header, 0)
	// To
	hdrs = append(hdrs, &sip.ToHeader{
		Address: device.Addr.Uri,
	})

	// From
	var contactUri sip.ContactUri
	if deviceContactMap[device.DeviceId] != nil {
		contactUri = deviceContactMap[device.DeviceId]
	} else {
		contactUri = device.Addr.Uri
	}
	hdrs = append(hdrs, &sip.FromHeader{
		Address: contactUri,
	})
	// Via
	hdrs = append(hdrs, &sip.ViaHeader{
		&sip.ViaHop{
			Params: sip.NewParams().Add("branch", sip.String{Str: sip.GenerateBranch()}),
		},
	})
	// Contact
	hdrs = append(hdrs, &sip.ContactHeader{
		Address: contactUri,
	})
	// ContentType
	hdrs = append(hdrs, &ContentTypeSDP)
	// CSeq Method
	hdrs = append(hdrs, &sip.CSeq{
		MethodName: sip.INVITE,
	})

	req := sip.NewRequest("", sip.INVITE, device.Addr.Uri, DefaultSipVersion, hdrs, string(byteData), nil)

	tx, err := server.Request(req)

	if err != nil {
		logging.Warn("sipPlayPush fail.id:", device.DeviceId, "err:", err)
		return data, err
	}

	return data, nil
}

func sipResponse(ctx sip.ClientTransaction) sip.Response {
	for {
		res := <-ctx.Responses()
		if res == nil {
			return res
		}

		logging.Debug("response tx", ctx.Key(), time.Now().Format("2006-01-02 15:04:05"))
		if res.StatusCode() == http.StatusContinue || res.statusCode == http.StatusSwitchingProtocols {
			// Trying and Dialog Establishement 等待下一个返回
			continue
		}
		return res
	}
}