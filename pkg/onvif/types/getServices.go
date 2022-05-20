package types

import (
	"encoding/xml"
)

// ************** GetServices Start **************
type GetServicesEnvelope struct {
	Envelope xml.Name        `xml:"Envelope"`
	Body     GetServicesBody `xml:"Body"`
}

type GetServicesBody struct {
	Body                xml.Name            `xml:"Body"`
	GetServicesResponse GetServicesResponse `xml:"GetServicesResponse"`
}

type GetServicesResponse struct {
	GetServicesResponse xml.Name `xml:"GetServicesResponse"`
	Service             []Service
}

type Service struct {
	Namespace    string
	XAddr        string
	Capabilities ServiceCapabilities
	Version      OnvifVersion
}

type ServiceCapabilities struct {
	// Any string
	DeviceCapabilities DeviceCapabilities `xml:"tds:Capabilities"`
	MediaCapabilities  MediaCapabilities  `xml:"trt:Capabilities"`
	TevCapabilities    TevCapabilities    `xml:"tev:Capabilities"`
	TimgCapabilities   TimgCapabilities   `xml:"timg:Capabilities"`
	Any                string
}

type TimgCapabilities struct {
	ImageStabilization bool `xml:"ImageStabilization,attr"`
}

type TevCapabilities struct {
	WSSubscriptionPolicySupport                   bool `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport                            bool `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
	MaxNotificationProducers                      int  `xml:"MaxNotificationProducers,attr"`
	MaxPullPoints                                 int  `xml:"MaxPullPoints,attr"`
	PersistentNotificationStorage                 bool `xml:"PersistentNotificationStorage,attr"`
}

type MediaCapabilities struct {
	SnapshotUri           bool `xml:"SnapshotUri,attr"`
	Rotation              bool `xml:"Rotation,attr"`
	VideoSourceMode       bool `xml:"VideoSourceMode,attr"`
	OSD                   bool `xml:"OSD,attr"`
	TemporaryOSDText      bool `xml:"TemporaryOSDText,attr"`
	EXICompression        bool `xml:"EXICompression,attr"`
	ProfileCapabilities   ProfileCapabilities
	StreamingCapabilities StreamingCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
}

type StreamingCapabilities struct {
	RTPMulticast        bool `xml:"RTPMulticast,attr"`
	RTP_TCP             bool `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP        bool `xml:"RTP_RTSP_TCP,attr"`
	NonAggregateControl bool `xml:"NonAggregateControl,attr"`
	NoRTSPStreaming     bool `xml:"NoRTSPStreaming,attr"`
}

type DeviceCapabilities struct {
	Network  NetworkCapabilities
	Security SecurityCapabilities
	System   SystemCapabilities
}

type DeviceServiceCapabilities struct {
	Network  NetworkCapabilities
	Security SecurityCapabilities
	System   SystemCapabilities
	Misc     MiscCapabilities
}

type NetworkCapabilities struct {
	IPFilter            bool `xml:"IPFilter,attr"`
	ZeroConfiguration   bool `xml:"ZeroConfiguration,attr"`
	IPVersion6          bool `xml:"IPVersion6,attr"`
	DynDNS              bool `xml:"DynDNS,attr"`
	Dot11Configuration  bool `xml:"Dot11Configuration,attr"`
	Dot1XConfigurations int  `xml:"Dot1XConfigurations,attr"`
	HostnameFromDHCP    bool `xml:"HostnameFromDHCP,attr"`
	NTP                 int  `xml:"NTP,attr"`
	DHCPv6              bool `xml:"DHCPv6,attr"`
}

type SecurityCapabilities struct {
	TLS1_0               bool           `xml:"TLS1_0,attr"`
	TLS1_1               bool           `xml:"TLS1_1,attr"`
	TLS1_2               bool           `xml:"TLS1_2,attr"`
	OnboardKeyGeneration bool           `xml:"OnboardKeyGeneration,attr"`
	AccessPolicyConfig   bool           `xml:"AccessPolicyConfig,attr"`
	DefaultAccessPolicy  bool           `xml:"DefaultAccessPolicy,attr"`
	Dot1X                bool           `xml:"Dot1X,attr"`
	RemoteUserHandling   bool           `xml:"RemoteUserHandling,attr"`
	X_509Token           bool           `xml:"X_509Token,attr"`
	SAMLToken            bool           `xml:"SAMLToken,attr"`
	KerberosToken        bool           `xml:"KerberosToken,attr"`
	UsernameToken        bool           `xml:"UsernameToken,attr"`
	HttpDigest           bool           `xml:"HttpDigest,attr"`
	RELToken             bool           `xml:"RELToken,attr"`
	SupportedEAPMethods  EAPMethodTypes `xml:"SupportedEAPMethods,attr"`
	MaxUsers             int            `xml:"MaxUsers,attr"`
	MaxUserNameLength    int            `xml:"MaxUserNameLength,attr"`
	MaxPasswordLength    int            `xml:"MaxPasswordLength,attr"`
}

type EAPMethodTypes struct {
	Types []int
}

type SystemCapabilities struct {
	DiscoveryResolve         bool           `xml:"DiscoveryResolve,attr"`
	DiscoveryBye             bool           `xml:"DiscoveryBye,attr"`
	RemoteDiscovery          bool           `xml:"RemoteDiscovery,attr"`
	SystemBackup             bool           `xml:"SystemBackup,attr"`
	SystemLogging            bool           `xml:"SystemLogging,attr"`
	FirmwareUpgrade          bool           `xml:"FirmwareUpgrade,attr"`
	HttpFirmwareUpgrade      bool           `xml:"HttpFirmwareUpgrade,attr"`
	HttpSystemBackup         bool           `xml:"HttpSystemBackup,attr"`
	HttpSystemLogging        bool           `xml:"HttpSystemLogging,attr"`
	HttpSupportInformation   bool           `xml:"HttpSupportInformation,attr"`
	StorageConfiguration     bool           `xml:"StorageConfiguration,attr"`
	MaxStorageConfigurations int            `xml:"MaxStorageConfigurations,attr"`
	GeoLocationEntries       int            `xml:"GeoLocationEntries,attr"`
	AutoGeo                  StringAttrList `xml:"AutoGeo,attr"`
}

type MiscCapabilities struct {
	AuxiliaryCommands StringAttrList `xml:"AuxiliaryCommands,attr"`
}
