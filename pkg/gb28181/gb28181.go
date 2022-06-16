package gb28181

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"test/lib"
	"test/log"
	"test/models"
	"test/pkg/logging"
	"test/pkg/setting"
	"test/pkg/sip"
	"test/pkg/sip/parser"
	"test/pkg/sip/sipserver"

	"github.com/sirupsen/logrus"
)

var (
	logger           log.Logger
	server           sipserver.Server
	defaultAlgorithm string = "MD5"
	gbConfig         Config
	serverDevices    models.Device
	addr             string
	deviceContactMap map[string]sip.ContactUri
)

func init() {
	deviceContactMap = make(map[string]sip.ContactUri, 0)

	logger = log.NewDefaultLogrusLogger().WithPrefix("SipServer")
	// logger.SetLevel(log.TraceLevel)
	logger.SetLevel(log.DebugLevel)
}

// 启动sip服务器
func Setup() {

	server = sipserver.NewServer(sipserver.ServerConfig{}, nil, nil, logger)
	server.Listen("udp", "0.0.0.0:5061", nil)
	server.OnRequest(sip.INVITE, INVITE)
	server.OnRequest(sip.REGISTER, REGISTER)
	server.OnRequest(sip.MESSAGE, MESSAGE)

	loadConfig()
}

func loadConfig() {
	gbConfig = Config{
		GB28181: SysInfo{
			IP:     setting.GBSetting.Ip,
			Port:   setting.GBSetting.Port,
			Region: setting.GBSetting.Region,
			LID:    setting.GBSetting.Lid,
			UID:    setting.GBSetting.Uid,
			DID:    setting.GBSetting.Did,
			UNUM:   setting.GBSetting.Unum,
			DNUM:   setting.GBSetting.Dnum,
		},
		Media: *setting.MediaSetting,
	}

	// 初始化媒体服务相关配置
	url, err := url.Parse(setting.MediaSetting.Rtp)
	if err != nil {
		logrus.Fatalf("media rtp url error,url:%s,err:%v", setting.MediaSetting.Rtp, err)
	}
	ipaddr, err := net.ResolveIPAddr("ip", url.Hostname())
	if err != nil {
		logrus.Fatalf("media rtp url error,url:%s,err:%v", setting.MediaSetting.Rtp, err)
	}
	gbConfig.mediaServerRtpIP = ipaddr.IP
	gbConfig.mediaServerRtpPort, _ = strconv.Atoi(url.Port())

	addr = fmt.Sprintf("%s:%s", gbConfig.GB28181.IP, gbConfig.GB28181.Port)
	uri, _ := parser.ParseSipUri(fmt.Sprintf("sip:%s@%s", gbConfig.GB28181.LID, addr))

	serverDevices = models.Device{
		DeviceId: gbConfig.GB28181.DID,
		Region:   gbConfig.GB28181.Region,
		Addr: &sip.Address{
			DisplayName: sip.String{Str: "sipserver"},
			Uri:         &uri,
			Params:      sip.NewParams(),
		},
	}

}

func INVITE(req sip.Request, tx sip.ServerTransaction) {

}

//
func MESSAGE(req sip.Request, tx sip.ServerTransaction) {
	fromDevice, ok := parserDevicesFromReqeust(req)

	if !ok {
		// 未解析出来源用户返回错误
		tx.Respond(sip.NewResponseFromRequest("", req, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), ""))
		return
	}
	// 判断是否存在body数据
	if len, have := req.ContentLength(); !have || len.Equals(0) {
		// 不存在就直接返回的成功
		tx.Respond(sip.NewResponseFromRequest("", req, http.StatusOK, "", ""))
		return
	}

	// 判断当前设备是否存在于数据库中 不存在则返回401
	if db_device, err := models.GetDeviceByDeviceId(fromDevice.DeviceId); err != nil {
		tx.Respond(sip.NewResponseFromRequest("", req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), ""))
		return
	} else {
		// 存在的话 生成contact_uri
		if deviceContactMap[fromDevice.DeviceId] == nil {
			uri, _ := parser.ParseSipUri(fmt.Sprintf("%s", db_device.Contact))
			deviceContactMap[fromDevice.DeviceId] = &uri
		}
	}

	body := req.Body()
	message := &MessageReceive{}

	if err := lib.XMLDecode([]byte(body), message); err != nil {
		logger.Errorf("Message Unmarshal xml err:", err, "body:", body)
		tx.Respond(sip.NewResponseFromRequest("", req, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), ""))
		return
	}
	switch message.CmdType {

	case "Catalog":
		// 设备列表
		sipMessageOnCatalog(fromDevice, body)
		tx.Respond(sip.NewResponseFromRequest("", req, http.StatusOK, "", ""))
		return
	case "Keepalive":
		// 心跳
		if err := sipMessageOnKeepAlive(fromDevice, body); err == nil {
			tx.Respond(sip.NewResponseFromRequest("", req, http.StatusOK, "", ""))
			// 心跳后同步注册设备列表信息
			sipCatalog(fromDevice)
			return
		}
	}
}

// 设备注册
func REGISTER(req sip.Request, tx sip.ServerTransaction) {

	// 判断请求头是否哦包含设备认证消息
	if len(req.GetHeaders("Authorization")) == 0 {
		resp := sip.NewResponseFromRequest("", req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "")
		resp.AppendHeader(&sip.GenericHeader{HeaderName: "WWW-Authenticate", Contents: fmt.Sprintf("Digest Nonce=\"%s\", algorithm=MD5, Realm=\"%s\",qop=\"auth\"", req.MessageID(), req.MessageID())})
		tx.Respond(resp)
		return
	}
	auth := sip.AuthFromValue(req.GetHeaders("Authorization")[0].Value())

	// 判断设备认证消息是否包含User
	if auth.Username() == "" {
		resp := sip.NewResponseFromRequest("", req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "")
		resp.AppendHeader(&sip.GenericHeader{HeaderName: "WWW-Authenticate", Contents: fmt.Sprintf("Digest Nonce=\"%s\", algorithm=MD5, Realm=\"%s\",qop=\"auth\"", req.MessageID(), req.MessageID())})
		tx.Respond(resp)
		return
	}

	fromDevice, ok := parserDevicesFromReqeust(req)
	if !ok {
		return
	}
	// 查询库中是否存在该NVR设备
	if device, err := models.GetDeviceByDeviceId(fromDevice.DeviceId); err == nil {
		deviceContactMap[fromDevice.DeviceId] = *fromDevice.ContactUri
		if device.Regist == 1 {
			logging.Debug("该设备已经注册过了~")
			// 执行信息修改操作
			device.Contact = fromDevice.Contact
			models.UpdateDevice(device.DeviceId, device)
		} else {
			models.AddDevice(map[string]interface{}{
				"name":        fromDevice.Name,
				"device_id":   fromDevice.DeviceId,
				"region":      fromDevice.Region,
				"host":        fromDevice.Host,
				"port":        fromDevice.Port,
				"proto":       fromDevice.Proto,
				"contact":     fromDevice.Contact,
				"device_type": "",
				"model_type":  "",
				"regist":      1,
				"active":      1,
			})
		}
	}

	// 响应注册成功
	resp := sip.NewResponseFromRequest("", req, http.StatusOK, http.StatusText(http.StatusOK), "")
	tx.Respond(resp)
}
