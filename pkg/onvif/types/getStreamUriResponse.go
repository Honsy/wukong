package types

import (
	"encoding/xml"
)

// ************** GetStreamUriResponse Start **************
type GetStreamUriEnvelope struct {
	Envelope xml.Name         `xml:"Envelope"`
	Body     GetStreamUriBody `xml:"Body"`
}

type GetStreamUriBody struct {
	Body                 xml.Name             `xml:"Body"`
	GetStreamUriResponse GetStreamUriResponse `xml:"GetStreamUriResponse"`
}

type GetStreamUriResponse struct {
	GetStreamUriResponse xml.Name `xml:"GetStreamUriResponse"`
	Uri                  string   `xml:"Uri"`
}
