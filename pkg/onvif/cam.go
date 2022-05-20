package onvif

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"test/pkg/onvif/types"

	"github.com/elgs/gostrgen"
)

// 小写字母开头的字段为不公开字段
type CameraOption struct {
	useSecure       bool
	secureOpts      string
	Hostname        string
	Username        string
	Passowrd        string
	Port            int
	path            string
	timeout         int
	agent           bool
	preserveAddress bool
}

type RequsetOptions struct {
	url     string // Url
	service string // service
	body    string // SOAP body
}

type PasswordDigest struct {
	passdigest string
	nonce      string
	timestamp  string
}

type ActiveSource struct {
	sourceToken                   string
	profileToken                  string
	videoSourceConfigurationToken string
	encoding                      string
	width                         int
	height                        int
	fps                           float32
	bitrate                       int
}

// props
var media2Support = false
var cameraOption CameraOption
var uri map[string]*url.URL
var capabilities types.Capabilities
var videoSources []types.VideoSources
var activeSource ActiveSource
var activeSources []ActiveSource
var profiles []types.Profile

func Setup(opts CameraOption) {
	uri = make(map[string]*url.URL)
	cameraOption = opts
	connect()
}

// 开始连接
func connect() {
	_, err := getSystemDateAndTime()
	if err != nil {
		log.Fatalf("getSystemDateAndTime Error %v", err)
	}
	_, err = getServices()
	if err != nil {
		log.Fatalf("getServices Error %v", err)
	}
	_, err = getCapabilities()
	if err != nil {
		log.Fatalf("getCapabilities Error %v", err)
	}

	// 全部成功判断
	if uri["media"] != nil {
		getProfiles()
		getVideoSources()
	}

	getActiveSources()

	log.Printf("ONVIF CONNECTED")
}

// 获取系统时间
func getSystemDateAndTime() (string, error) {
	res, err := _request(RequsetOptions{
		body: "<s:Envelope xmlns:s='http://www.w3.org/2003/05/soap-envelope'>" +
			"<s:Body xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance' xmlns:xsd='http://www.w3.org/2001/XMLSchema'>" +
			"<GetSystemDateAndTime xmlns='http://www.onvif.org/ver10/device/wsdl'/>" +
			"</s:Body>" +
			"</s:Envelope>",
	})

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	dataEnvelope := new(DataEnvelope)
	err = xml.Unmarshal(data, dataEnvelope)
	if err != nil {
		return "", err
	}

	dateTime := dataEnvelope.Body.GetSystemDateAndTimeResponse.SystemDateAndTime.UTCDateTime
	localDateTime := dataEnvelope.Body.GetSystemDateAndTimeResponse.SystemDateAndTime.LocalDateTime
	var currentTime UTCDateTime
	if dateTime != (UTCDateTime{}) {
		currentTime = dateTime
	} else {
		currentTime = localDateTime
	}

	time := fmt.Sprintf("%s-%s-%s %s:%s:%s", currentTime.Date.Year,
		currentTime.Date.Month,
		currentTime.Date.Day,
		currentTime.Time.Hour,
		currentTime.Time.Minute,
		currentTime.Time.Second,
	)

	return time, nil
}

// 获取onvif服务
func getServices() (string, error) {
	res, err := _request(RequsetOptions{
		body: _envelopeHeader() +
			"<GetServices xmlns='http://www.onvif.org/ver10/device/wsdl'>" +
			"<IncludeCapability>true</IncludeCapability>" +
			"</GetServices>" +
			_envelopeFooter(),
	})

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	serviceEnvelope := new(types.GetServicesEnvelope)
	err = xml.Unmarshal(data, serviceEnvelope)
	if err != nil {
		return "", err
	}

	for _, value := range serviceEnvelope.Body.GetServicesResponse.Service {
		// uri解析
		parsedNameSpace, err := url.Parse(value.Namespace)
		if err != nil {
			log.Println("Url Parse Error! ")
			break
		}
		host := strings.Split(parsedNameSpace.Host, ":")[0]
		if host == "www.onvif.org" {
			namespaceSplitted := strings.Split(parsedNameSpace.Path, "/")

			if namespaceSplitted[2] == "media" && namespaceSplitted[1] == "ver20" {
				media2Support = true
				namespaceSplitted[2] = "media2"
			}
			uri[namespaceSplitted[2]], err = url.Parse(value.XAddr)
		}
	}

	return "", nil
}

// 获取Capabilities
func getCapabilities() (string, error) {
	res, err := _request(RequsetOptions{
		body: _envelopeHeader() +
			"<GetCapabilities xmlns='http://www.onvif.org/ver10/device/wsdl'>" +
			"<Category>All</Category>" +
			"</GetCapabilities>" +
			_envelopeFooter(),
	})

	if err != nil {
		log.Fatalf("GetCapabilities SOAP Error !")
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("GetCapabilities IO Read Error !")
		return "", err
	}

	capabilitiesEnvelope := new(types.GetCapabilitiesEnvelope)
	err = xml.Unmarshal(data, capabilitiesEnvelope)
	if err != nil {
		return "", err
	}

	capabilities = capabilitiesEnvelope.Body.GetCapabilitiesResponse.Capabilities

	uri["device"], _ = url.Parse(capabilities.Device.XAddr)
	uri["imaging"], _ = url.Parse(capabilities.Imaging.XAddr)
	uri["events"], _ = url.Parse(capabilities.Events.XAddr)

	return "", nil
}

// 生成Token
func generateToken(Username string, Nonce string, Created string, Password string) string {
	sDec, _ := base64.StdEncoding.DecodeString(Nonce)

	hasher := sha1.New()
	hasher.Write([]byte(string(sDec) + Created + Password))

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func _passwordDigest() PasswordDigest {
	charsToGenerate := 32
	charSet := gostrgen.Lower | gostrgen.Digit
	nonce, _ := gostrgen.RandGen(charsToGenerate, charSet, "", "")
	timestamp := time.Now().Format(time.RFC3339)
	passdigest := generateToken(cameraOption.Username, nonce, timestamp, cameraOption.Passowrd)

	pdg := PasswordDigest{
		passdigest: passdigest,
		nonce:      nonce,
		timestamp:  timestamp,
	}
	return pdg
}

func getActiveSources() {
	defaultProfile := types.Profile{}
	defaultProfiles := make([]types.Profile, len(videoSources))
	activeSources = make([]ActiveSource, len(videoSources))
	appropriateProfiles := []types.Profile{}

	for idx, videoSource := range videoSources {
		for _, profile := range profiles {
			if profile.Configurations.VideoSource != (types.VideoSourceConfiguration{}) && profile.Configurations.VideoEncoder != (types.VideoEncoderConfiguration{}) && profile.Configurations.VideoSource.SourceToken == videoSource.Token {
				appropriateProfiles = append(appropriateProfiles, profile)
			}
		}
		if len(appropriateProfiles) == 0 {
			log.Printf("No Profiles")
			return
		}
		if idx == 0 {
			defaultProfile = appropriateProfiles[0]
		}
		defaultProfiles[idx] = appropriateProfiles[0]
		tmpActiveSource := ActiveSource{
			sourceToken:                   videoSource.Token,
			profileToken:                  defaultProfile.Token,
			videoSourceConfigurationToken: defaultProfile.Configurations.VideoSource.Token,
		}
		if defaultProfile.Configurations.VideoEncoder != (types.VideoEncoderConfiguration{}) {
			tmpActiveSource.encoding = defaultProfile.Configurations.VideoEncoder.Encoding
			tmpActiveSource.width = defaultProfile.Configurations.VideoEncoder.Resolution.Width
			tmpActiveSource.height = defaultProfile.Configurations.VideoEncoder.Resolution.Height
			tmpActiveSource.fps = defaultProfile.Configurations.VideoEncoder.RateControl.FrameRateLimit
			tmpActiveSource.bitrate = defaultProfile.Configurations.VideoEncoder.RateControl.BitrateLimit
		}

		activeSources[idx] = tmpActiveSource

		if idx == 0 {
			activeSource = activeSources[idx]
		}
	}
}

func _envelopeHeader() string {
	header := "<s:Envelope xmlns:s='http://www.w3.org/2003/05/soap-envelope' xmlns:a='http://www.w3.org/2005/08/addressing'>" +
		"<s:Header>"
	if cameraOption.Username != "" && cameraOption.Passowrd != "" {
		req := _passwordDigest()
		header += "<Security s:mustUnderstand='1' xmlns='http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd'>" +
			"<UsernameToken>" +
			"<Username>" + cameraOption.Username + "</Username>" +
			"<Password Type='http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest'>" + req.passdigest + "</Password>" +
			"<Nonce EncodingType='http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary'>" + req.nonce + "</Nonce>" +
			"<Created xmlns='http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd'>" + req.timestamp + "</Created>" +
			"</UsernameToken>" +
			"</Security>" +
			"</s:Header>" +
			"<s:Body xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance' xmlns:xsd='http://www.w3.org/2001/XMLSchema'>"
		return header
	}

	return ""
}

func _envelopeFooter() string {
	return "</s:Body>" +
		"</s:Envelope>"
}

// SOAP requset
func _request(requestOpts RequsetOptions) (*http.Response, error) {
	path := "/onvif/device_service"
	if uri[requestOpts.service] != nil {
		path = uri[requestOpts.service].Path
	}

	bufferBody := []byte(requestOpts.body)
	url := fmt.Sprintf("http://%s:%d%s", cameraOption.Hostname, cameraOption.Port, path)
	// Http client
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// 建立请求
	req, err := http.NewRequest("POST", url, strings.NewReader(requestOpts.body))
	if err != nil {
		log.Fatalln("Create Request Failed!")
		return nil, err
	}
	req.Header.Add("Content-Type", "application/soap+xml")
	req.Header.Add("Content-Length", string(len(bufferBody)))
	req.Header.Add("charset", "utf-8")

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln("Request Failed", err)
		return nil, err
	}

	return res, nil
}
