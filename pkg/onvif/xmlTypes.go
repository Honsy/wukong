package onvif

import "encoding/xml"

// ************** GetSystemDateAndTime Start**************
type DataEnvelope struct {
	Envelope xml.Name `xml:"Envelope"`
	Body     Body     `xml:"Body"`
}

type Body struct {
	Body                         xml.Name                     `xml:"Body"`
	GetSystemDateAndTimeResponse GetSystemDateAndTimeResponse `xml:"GetSystemDateAndTimeResponse"`
}

type GetSystemDateAndTimeResponse struct {
	GetSystemDateAndTimeResponse xml.Name          `xml:"GetSystemDateAndTimeResponse"`
	SystemDateAndTime            SystemDateAndTime `xml:"SystemDateAndTime"`
}

type SystemDateAndTime struct {
	XMLName         xml.Name    `xml:"SystemDateAndTime"`
	DateTimeType    string      `xml:"DateTimeType"`
	DaylightSavings string      `xml:"DaylightSavings"`
	TimeZone        TimeZone    `xml:"TimeZone"`
	UTCDateTime     UTCDateTime `xml:"UTCDateTime"`
	LocalDateTime   UTCDateTime `xml:"LocalDateTime"`
}

type TimeZone struct {
	XMLName xml.Name `xml:"TimeZone"`
	TZ      string   `xml:"TZ"`
}

type UTCDateTime struct {
	Time Time `xml:"Time"`
	Date Date `xml:"Date"`
}

type Time struct {
	XMLName xml.Name `xml:"Time"`
	Hour    string   `xml:"Hour"`
	Minute  string   `xml:"Minute"`
	Second  string   `xml:"Second"`
}

type Date struct {
	XMLName xml.Name `xml:"Date"`
	Year    string   `xml:"Year"`
	Month   string   `xml:"Month"`
	Day     string   `xml:"Day"`
}

// ************** GetSystemDateAndTime End **************
