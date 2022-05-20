package types

import "encoding/xml"

// ************** GetCapabilities Start **************
type GetCapabilitiesEnvelope struct {
	Envelope xml.Name            `xml:"Envelope"`
	Body     GetCapabilitiesBody `xml:"Body"`
}

type GetCapabilitiesBody struct {
	Body                    xml.Name                `xml:"Body"`
	GetCapabilitiesResponse GetCapabilitiesResponse `xml:"GetCapabilitiesResponse"`
}

type GetCapabilitiesResponse struct {
	GetCapabilitiesResponse xml.Name `xml:"GetCapabilitiesResponse"`
	Capabilities            Capabilities
}

type Capabilities struct {
	Capabilities xml.Name              `xml:"Capabilities"`
	Analytics    CapabilitiesAnalytics `xml:"Analytics"`
	Device       CapabilitiesDevice    `xml:"Device"`
	Events       CapabilitiesEvent     `xml:"Events"`
	Imaging      CapabilitiesImaging   `xml:"Imaging"`
}

type CapabilitiesAnalytics struct {
	XAddr                  string `xml:"XAddr"`
	RuleSupport            bool   `xml:"RuleSupport"`
	AnalyticsModuleSupport bool   `xml:"AnalyticsModuleSupport"`
}

type CapabilitiesDevice struct {
	XAddr    string                     `xml:"XAddr"`
	Network  CapabilitiesDeviceNetwork  `xml:"Network"`
	System   CapabilitiesDeviceSystem   `xml:"System"`
	IO       CapabilitiesDeviceIO       `xml:"IO"`
	Security CapabilitiesDeviceSecurity `xml:"Security"`
}

type CapabilitiesDeviceNetwork struct {
	IPFilter          bool `xml:"IPFilter"`
	ZeroConfiguration bool `xml:"ZeroConfiguration"`
	IPVersion6        bool `xml:"IPVersion6"`
	DynDNS            bool `xml:"DynDNS"`
}

type CapabilitiesDeviceSystem struct {
	DiscoveryResolve  bool           `xml:"DiscoveryResolve"`
	DiscoveryBye      bool           `xml:"DiscoveryBye"`
	RemoteDiscovery   bool           `xml:"RemoteDiscovery"`
	SystemBackup      bool           `xml:"SystemBackup"`
	SystemLogging     bool           `xml:"SystemLogging"`
	FirmwareUpgrade   bool           `xml:"FirmwareUpgrade"`
	SupportedVersions []OnvifVersion `xml:"SupportedVersions"`
}

type CapabilitiesDeviceIO struct {
	InputConnectors int `xml:"InputConnectors"`
	RelayOutputs    int `xml:"RelayOutputs"`
}

type CapabilitiesDeviceSecurity struct {
	TLS1                 bool `xml:"TLS1.1"`
	TLS2                 bool `xml:"TLS1.2"`
	OnboardKeyGeneration bool `xml:"OnboardKeyGeneration"`
	Token                bool `xml:"X.509Token"`
	SAMLToken            bool `xml:"SAMLToken"`
	KerberosToken        bool `xml:"KerberosToken"`
	RELToken             bool `xml:"RELToken"`
}

type CapabilitiesEvent struct {
	XAddr                                         string `xml:"XAddr"`
	WSSubscriptionPolicySupport                   bool   `xml:"WSSubscriptionPolicySupport"`
	WSPullPointSupport                            bool   `xml:"WSPullPointSupport"`
	WSPausableSubscriptionManagerInterfaceSupport bool   `xml:"WSPausableSubscriptionManagerInterfaceSupport"`
}

type CapabilitiesImaging struct {
	XAddr string `xml:"XAddr"`
}

type CapabilitiesMedia struct {
	XAddr                 string                            `xml:"XAddr"`
	StreamingCapabilities CapabilitiesStreamingCapabilities `xml:"StreamingCapabilities"`
}

type CapabilitiesStreamingCapabilities struct {
	RTPMulticast bool `xml:"RTPMulticast"`
	RTP_TCP      bool `xml:"RTP_TCP"`
	RTP_RTSP_TCP bool `xml:"RTP_RTSP_TCP"`
}
