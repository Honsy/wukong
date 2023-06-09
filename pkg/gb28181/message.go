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

// SDP
var ContentTypeSDP = sip.ContentType("application/sdp")

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
	var active = 0
	message := &MessageNotify{}
	if err := lib.XMLDecode([]byte(body), message); err != nil {
		logging.Error("Message Unmarshal xml err:", err, "body:", body)
		return err
	}

	// 更新全局变量
	if message.Status == "OK" {
		active = 1
		activeDevices.Store(u.DeviceId, u)
	} else {
		active = 0
		activeDevices.Delete(u.DeviceId)
	}
	// 更新数据库状态
	models.UpdateDevice(u.DeviceId, map[string]interface{}{
		"active": active,
	})
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
		for _, camera := range message.Item {
			if dbCamera, err := models.GetCamera(camera.DeviceID); err == nil {
				dbCamera.PDID = u.DeviceId
				dbCamera.Active = time.Now().Unix()
				dbCamera.URIStr = fmt.Sprintf("sip:%s@%s", camera.DeviceID, gbConfig.GB28181.Region)
				models.UpdateCamera(dbCamera.ID, dbCamera)
			} else {
				camera.PDID = u.DeviceId
				// 此处自动接收NVR或者DVR子设备直接注册
				models.InsertCamera(camera)
				logging.Info("设备未找到", camera.DeviceID)
			}
		}
	}
	return nil
}

// 查询设备列表请求
func sipCatalog(to models.Device) {
	// 如果ContactMap存在使用map存放数据
	var contactUri sip.ContactUri
	if deviceContactMap[to.DeviceId] != nil {
		contactUri = deviceContactMap[to.DeviceId]
	} else {
		contactUri = to.Addr.Uri
	}

	req := sip.NewRequest("", sip.MESSAGE, contactUri, DefaultSipVersion, []sip.Header{}, GetCatalogXML(to.DeviceId), nil)
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
