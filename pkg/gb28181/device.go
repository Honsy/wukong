package gb28181

import (
	"test/models"
	"test/pkg/logging"
	"test/pkg/sip"
)

// 从请求中解析出设备信息
func parserDevicesFromReqeust(req sip.Request) (models.Device, bool) {
	u := models.Device{}
	header, ok := req.From()
	if !ok {
		logging.Warn("not found from header from request", req.String())
		return u, false
	}
	if header.Address == nil {
		logging.Warn("not found from user from request", req.String())
		return u, false
	}
	if header.Address.User() == nil {
		logging.Warn("not found from user from request", req.String())
		return u, false
	}
	u.DeviceId = header.Address.User().String()
	u.Region = header.Address.Host()
	via, ok := req.ViaHop()
	if !ok {
		logging.Info("not found ViaHop from request", req.String())
		return u, false
	}
	u.Proto = via.ProtocolName
	u.DeviceType = req.GetHeaders("User-Agent")[0].String()
	u.Host = via.Host
	u.Port = via.Port.String()
	contact, ok := req.Contact()
	if ok {
		u.Contact = contact.Address.String()
		u.ContactUri = &contact.Address
	}

	report, ok := via.Params.Get("rport")
	if ok && report != nil {
		u.Rport = report.String()
	}
	raddr, ok := via.Params.Get("received")
	if ok && raddr != nil {
		u.RAddr = raddr.String()
	}
	u.TransPort = via.Transport
	u.URIStr = header.Address.String()
	u.Addr = sip.NewAddressFromFromHeader(header)
	return u, true
}
