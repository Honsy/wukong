package types

import (
	"encoding/xml"
)

// ************** GetServices Start **************
type GetProfilesEnvelope struct {
	Envelope xml.Name        `xml:"Envelope"`
	Body     GetProfilesBody `xml:"Body"`
}

type GetProfilesBody struct {
	Body                xml.Name            `xml:"Body"`
	GetProfilesResponse GetProfilesResponse `xml:"GetProfilesResponse"`
}

type GetProfilesResponse struct {
	GetProfilesResponse xml.Name  `xml:"GetProfilesResponse"`
	Profiles            []Profile `xml:"Profiles"`
}

type Profile struct {
	Name           string
	Configurations ProfileConfigurations
	Token          string `xml:"token,attr"`
}
type ProfileConfigurations struct {
	Name         string
	VideoSource  VideoSourceConfiguration
	AudioSource  AudioSourceConfiguration
	VideoEncoder VideoEncoderConfiguration
	AudioEncoder AudioEncoderConfiguration
	Analytics    AnalyticsConfiguration
}

type VideoSourceConfiguration struct {
	Name        string       `xml:"Name"`
	UseCount    int          `xml:"UseCount"`
	SourceToken string       `xml:"SourceToken"`
	Bounds      IntRectangle `xml:"Bounds"`
	Token       string       `xml:"token,attr"`
}

type AudioSourceConfiguration struct {
	Name        string `xml:"Name"`
	UseCount    int    `xml:"UseCount"`
	SourceToken string `xml:"SourceToken"`
}

type VideoEncoderConfiguration struct {
	Name        string                 `xml:"Name"`
	UseCount    int                    `xml:"UseCount"`
	Encoding    string                 `xml:"Encoding"`
	Resolution  VideoResolution        `xml:"Resolution"`
	RateControl VideoRateControl       `xml:"RateControl"`
	Multicast   MulticastConfiguration `xml:"Multicast"`
	Quality     string                 `xml:"Quality"`
}

type AudioEncoderConfiguration struct {
	Name       string                 `xml:"Name"`
	UseCount   int                    `xml:"UseCount"`
	Encoding   string                 `xml:"Encoding"`
	Multicast  MulticastConfiguration `xml:"Multicast"`
	Bitrate    int                    `xml:"Bitrate"`
	SampleRate int                    `xml:"SampleRate"`
}

type AnalyticsConfiguration struct {
	Name                         string                       `xml:"Name"`
	UseCount                     int                          `xml:"UseCount"`
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"AnalyticsEngineConfiguration"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule []AnalyticsModuleConfiguration
}

type AnalyticsModuleConfiguration struct {
	Parameters AnalyticsModuleParameters `xml:"Parameters"`
}

type AnalyticsModuleParameters struct {
	SimpleItem  SimpleItem  `xml:"SimpleItem"`
	ElementItem ElementItem `xml:"ElementItem"`
}

type SimpleItem struct {
	Name  string `xml:"Name,attr"`
	Value int    `xml:"Value,attr"`
}
type ElementItem struct {
	Name string `xml:"Name,attr"`
}
