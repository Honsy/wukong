package onvif

import (
	"encoding/xml"
	"io/ioutil"
	"test/pkg/onvif/types"
)

type StreamOptions struct {
	stream       string
	protocol     string
	profileToken string
}

// 获取onvif对应rtsp url
func GetStreamUri(options StreamOptions) (string, error) {
	sopts := StreamOptions{
		stream:   "RTP-Unicast",
		protocol: "RTSP",
	}

	if options.protocol != "" {
		sopts.protocol = options.protocol
	}
	if options.stream != "" {
		sopts.stream = options.stream
	}

	if sopts.protocol == "HTTP" {
		sopts.protocol = "RtspOverHttp"
	}
	if sopts.protocol == "UDP" && sopts.stream == "RTP-Unicast" {
		sopts.protocol = "RtspUnicast"
	}
	if sopts.protocol == "UDP" && sopts.stream == "RTP-Multicast" {
		sopts.protocol = "RtspMulticast"
	}

	if options.profileToken == "" {
		options.profileToken = activeSource.profileToken
	}

	if media2Support {
		res, err := _request(RequsetOptions{
			service: "media2",
			body: _envelopeHeader() +
				"<GetStreamUri xmlns='http://www.onvif.org/ver20/media/wsdl'>" +
				"<Protocol>" + sopts.protocol + "</Protocol>" +
				"<ProfileToken>" + options.profileToken + "</ProfileToken>" +
				"</GetStreamUri>" +
				_envelopeFooter(),
		})
		if err != nil {
			logger.Fatalf("GetStreamUri Error %s", err)
			return "", err
		}
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return "", err
		}

		streamUriEnvelope := new(types.GetStreamUriEnvelope)

		err = xml.Unmarshal(data, streamUriEnvelope)

		if err != nil {
			logger.Fatalf("GetStreamUriEnvelope XML Parse Error!")
			return "", err
		}

		return streamUriEnvelope.Body.GetStreamUriResponse.Uri, nil
	} else {
		_request(RequsetOptions{
			service: "media",
		})

		return "", nil
	}
}

/**
* Receive video sources
 */
func getProfiles() (string, error) {
	// Media2
	if media2Support {
		res, err := _request(RequsetOptions{
			service: "media2",
			body: _envelopeHeader() +
				"<GetProfiles xmlns='http://www.onvif.org/ver20/media/wsdl'>" +
				"<Type>All</Type>" +
				"</GetProfiles>" +
				_envelopeFooter(),
		})
		if err != nil {
			logger.Fatalf("GetProfiles Error!")
			return "", nil
		}

		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return "", err
		}

		profileEnvelope := new(types.GetProfilesEnvelope)
		err = xml.Unmarshal(data, profileEnvelope)
		if err != nil {
			logger.Fatalf("GetVideoSourcesEnvelope XML Parse Error! %v", err)
			return "", err
		}

		profiles = profileEnvelope.Body.GetProfilesResponse.Profiles

		return "", nil
	} else {
		res, err := _request(RequsetOptions{
			service: "media",
			body: _envelopeHeader() +
				"<GetProfiles xmlns='http://www.onvif.org/ver10/media/wsdl'/>" +
				_envelopeFooter(),
		})

		if err != nil {
			logger.Fatalf("GetProfiles Error!")
			return "", nil
		}

		_, err = ioutil.ReadAll(res.Body)

		if err != nil {
			return "", err
		}

		// log.Printf("getProfiles %s", string(data))
		return "", nil
	}

}

func getVideoSources() (string, error) {
	res, err := _request(RequsetOptions{
		service: "media",
		body: _envelopeHeader() +
			"<GetVideoSources xmlns='http://www.onvif.org/ver10/media/wsdl'/>" +
			_envelopeFooter(),
	})

	if err != nil {
		logger.Fatalf("GetVideoSources Error!")
		return "", err
	}
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	videoSourceEnvelope := new(types.GetVideoSourcesEnvelope)
	err = xml.Unmarshal(data, videoSourceEnvelope)
	if err != nil {
		logger.Fatalf("GetVideoSourcesEnvelope XML Parse Error!")
		return "", err
	}

	videoSources = []types.VideoSources{videoSourceEnvelope.Body.GetVideoSourcesResponse.VideoSources}

	return "", nil
}
