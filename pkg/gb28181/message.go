package gb28181

import (
	"fmt"
	"math/rand"
	"test/lib"
	"test/models"
	"test/pkg/logging"
	"test/pkg/sip"
	"time"
)

var DefaultSipVersion = "SIP/2.0"

// ContentTypeXML XML contenttype
var ContentTypeXML = sip.ContentType("Application/MANSCDP+xml")

var (
	// CatalogXML 获取设备列表xml样式
	CatalogXML = `<?xml version="1.0"?>
<Query>
<CmdType>Catalog</CmdType>
<SN>17430</SN>
<DeviceID>%s</DeviceID>
</Query>`
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

// MessageDeviceListResponse 设备明细列表返回结构
type MessageDeviceListResponse struct {
	CmdType  string          `xml:"CmdType"`
	SN       int             `xml:"SN"`
	DeviceID string          `xml:"DeviceID"`
	SumNum   int             `xml:"SumNum"`
	Item     []models.Camera `xml:"DeviceList>Item"`
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

// 设备列表数据处理
func sipMessageOnCatalog(u models.Device, body string) error {
	message := &MessageDeviceListResponse{}
	if err := lib.XMLDecode([]byte(body), message); err != nil {
		logging.Error("Message Unmarshal xml err:", err, "body:", body)
		return err
	}
	if message.SumNum > 0 {
		for _, d := range message.Item {
			if camera, err := models.GetCamera(d.DeviceID); err == nil {
				camera.Active = time.Now().Unix()
				models.UpdateCamera(camera.ID, camera)
			} else {
				logging.Info("设备未找到", d.DeviceID)
			}
		}
	}
	return nil
}

// 查询设备列表请求
func sipCatalog(to models.Device) {
	req := sip.NewRequest("", sip.MESSAGE, to.Addr.Uri, DefaultSipVersion, []sip.Header{}, GetCatalogXML(to.DeviceId), nil)
	rand.Seed(time.Now().UnixNano())
	callId := sip.CallID(fmt.Sprintf("%10f", rand.Float64()))
	req.AppendHeader(&callId)
	req.AppendHeader(&sip.CSeq{
		SeqNo:      20,
		MethodName: "MESSAGE",
	})
	req.AppendHeader(&ContentTypeXML)
	req.AppendHeader(&sip.FromHeader{
		Address: serverDevices.Addr.Uri,
	})
	req.AppendHeader(&sip.ToHeader{
		Address: to.Addr.Uri,
	})
	_, err := server.Request(req)
	if err != nil {
		logging.Error("sipCatalog  error,", err)
		return
	}
}
