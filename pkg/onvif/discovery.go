package onvif

import (
	"log"
	"net"

	"github.com/gofrs/uuid"
)

func probe() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(10, 0, 32, 32),
		Port: 80,
	}
	socket, err := net.DialUDP("upd", nil, addr)

	if err != nil {
		log.Fatalln("upd connect error !")
		return
	}
	messageID, err := uuid.NewV4()

	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	requestStr :=
		"<Envelope xmlns='http://www.w3.org/2003/05/soap-envelope' xmlns:dn='http://www.onvif.org/ver10/network/wsdl'>" +
			"<Header>" +
			"<wsa:MessageID xmlns:wsa='http://schemas.xmlsoap.org/ws/2004/08/addressing'>" + messageID.String() + "</wsa:MessageID>" +
			"<wsa:To xmlns:wsa='http://schemas.xmlsoap.org/ws/2004/08/addressing'>urn:schemas-xmlsoap-org:ws:2005:04:discovery</wsa:To>" +
			"<wsa:Action xmlns:wsa='http://schemas.xmlsoap.org/ws/2004/08/addressing'>http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</wsa:Action>" +
			"</Header>" +
			"<Body>" +
			"<Probe xmlns='http://schemas.xmlsoap.org/ws/2005/04/discovery' xmlns:xsd='http://www.w3.org/2001/XMLSchema' xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'>" +
			"<Types>dn:NetworkVideoTransmitter</Types>" +
			"<Scopes />" +
			"</Probe>" +
			"</Body>" +
			"</Envelope>"

	socket.WriteTo([]byte(requestStr), addr)

	for {
		var data [1024]byte
		_, _, read_err := socket.ReadFromUDP(data[:])
		if read_err != nil {
			log.Fatalf("read upd error !")
		}

		log.Fatalln(data)
	}
}
