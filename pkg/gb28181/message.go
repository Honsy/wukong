package gb28181

import (
	"fmt"
	"test/lib"
	"test/models"
	"test/pkg/logging"
	"test/pkg/sip"
)

var DefaultSipVersion = "SIP/2.0"

var (
	// CatalogXML 获取设备列表xml样式
	CatalogXML = `<?xml version="1.0"?>
<Query>
<CmdType>Catalog</CmdType>
<SN>17430</SN>
<DeviceID>%s</DeviceID>
</Query>
	`
	// RecordInfoXML 获取录像文件列表xml样式
	RecordInfoXML = `<?xml version="1.0"?>
<Query>
<CmdType>RecordInfo</CmdType>
<SN>17430</SN>
<DeviceID>%s</DeviceID>
<StartTime>%s</StartTime>
<EndTime>%s</EndTime>
<Secrecy>0</Secrecy>
<Type>time</Type>
</Query>
`
	// DeviceInfoXML 查询设备详情xml样式
	DeviceInfoXML = `<?xml version="1.0"?>
<Query>
<CmdType>DeviceInfo</CmdType>
<SN>17430</SN>
<DeviceID>%s</DeviceID>
</Query>
`
)

// MessageNotify 心跳包xml结构
type MessageNotify struct {
	CmdType  string `xml:"CmdType"`
	SN       int    `xml:"SN"`
	DeviceID string `xml:"DeviceID"`
	Status   string `xml:"Status"`
	Info     string `xml:"Info"`
}

// GetCatalogXML 获取NVR下设备列表指令
func GetCatalogXML(id string) string {
	return fmt.Sprintf(CatalogXML, id)
}

// 心跳逻辑处理
func sipMessageOnKeepAlive(u models.Device, body string) error {
	message := &MessageNotify{}
	if err := lib.XMLDecode([]byte(body), message); err != nil {
		logging.Error("Message Unmarshal xml err:", err, "body:", body)
		return err
	}

	// 更新数据库设备状态

	return nil
}

// 查询设备列表请求
func sipCatalog(to models.Device) {
	req := sip.NewRequest("", sip.MESSAGE, to.Addr, DefaultSipVersion, []sip.Header{}, GetCatalogXML(to.DeviceId), nil)
	_, err := server.Request(req)
	if err != nil {
		logging.Warn("sipCatalog  error,", err)
		return
	}
}
