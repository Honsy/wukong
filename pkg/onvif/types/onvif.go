package types

type OnvifVersion struct {
	Major int
	Minor int
}

type IntRectangle struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type VideoResolution struct {
	Width  int `xml:"Width"`
	Height int `xml:"Height"`
}

type VideoRateControl struct {
	FrameRateLimit   float32 `xml:"FrameRateLimit"`
	EncodingInterval int     `xml:"EncodingInterval"`
	BitrateLimit     int     `xml:"BitrateLimit"`
}

type MulticastConfiguration struct {
	Address   IPAddress `xml:"Address"`
	Port      int       `xml:"Port"`
	TTL       int       `xml:"TTL"`
	AutoStart bool      `xml:"AutoStart"`
}

type IPAddress struct {
	Type        string `xml:"Type"`
	IPv4Address string `xml:"IPv4Address"`
	IPv6Address string `xml:"IPv6Address"`
}

type StringAttrList struct {
	AttrList []string
}
