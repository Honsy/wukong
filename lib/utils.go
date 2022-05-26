package lib

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

func UnicodeDecode(v string) (string, error) {
	quoted := strconv.QuoteToASCII(v)
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(quoted), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

func ResolveSelfIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip, nil
		}
	}
	return nil, errors.New("server not connected to any network")
}

// Check two uint16 pointers for equality as follows:
// - If neither pointer is nil, check equality of the underlying uint16s.
// - If either pointer is nil, return true if and only if they both are.
func Uint16PtrEq(a *uint16, b *uint16) bool {
	if a == nil || b == nil {
		return a == b
	}

	return *a == *b
}

// XMLDecode XMLDecode
func XMLDecode(data []byte, v interface{}) error {
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	decoder.CharsetReader = charset.NewReaderLabel
	return decoder.Decode(v)
}
