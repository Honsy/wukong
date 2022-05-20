package types

import (
	"encoding/xml"
)

// ************** GetServices Start **************
type GetVideoSourcesEnvelope struct {
	Envelope xml.Name            `xml:"Envelope"`
	Body     GetVideoSourcesBody `xml:"Body"`
}

type GetVideoSourcesBody struct {
	Body                    xml.Name                `xml:"Body"`
	GetVideoSourcesResponse GetVideoSourcesResponse `xml:"GetVideoSourcesResponse"`
}

type GetVideoSourcesResponse struct {
	GetVideoSourcesResponse xml.Name `xml:"GetVideoSourcesResponse"`
	VideoSources            VideoSources
}

type VideoSources struct {
	Framerate  int                 `xml:"Framerate"`
	Resolution VideoResolution     `xml:"Resolution"`
	Imaging    VideoSourcesImaging `xml:"Imaging"`
	Token      string              `xml:"token,attr"`
}

type VideoSourcesImaging struct {
	BacklightCompensation BacklightCompensation `xml:"BacklightCompensation"`
	Brightness            int                   `xml:"Brightness"`
	ColorSaturation       int                   `xml:"ColorSaturation"`
	Contrast              int                   `xml:"Contrast"`
	Exposure              Exposure              `xml:"Exposure"`
	IrCutFilter           string                `xml:"IrCutFilter"`
	Sharpness             int                   `xml:"Sharpness"`
	WideDynamicRange      WideDynamicRange      `xml:"WideDynamicRange"`
	WhiteBalance          WhiteBalance          `xml:"WhiteBalance"`
}

type BacklightCompensation struct {
	Mode  string `xml:"Mode"`
	Level int    `xml:"Level"`
}

type Exposure struct {
	Mode            string `xml:"Mode"`
	Priority        string `xml:"Priority"`
	MinExposureTime int    `xml:"MinExposureTime"`
	MaxExposureTime int    `xml:"MaxExposureTime"`
	MinGain         int    `xml:"MinGain"`
	MaxGain         int    `xml:"MaxGain"`
	MinIris         int    `xml:"MinIris"`
	MaxIris         int    `xml:"MaxIris"`
	ExposureTime    int    `xml:"ExposureTime"`
	Gain            int    `xml:"Gain"`
	Iris            int    `xml:"Iris"`
}

type WideDynamicRange struct {
	Mode  string `xml:"Mode"`
	Level int    `xml:"Level"`
}

type WhiteBalance struct {
	Mode   string `xml:"Mode"`
	CrGain int    `xml:"CrGain"`
	CbGain int    `xml:"CbGain"`
}
