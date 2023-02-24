package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"test/lib"
	"test/pkg/logging"
	"test/pkg/setting"
)

var addStreamProxyUrl = fmt.Sprintf("%s%s", setting.MediaSetting.Http, "/index/api/addStreamProxy")

// 启动拉流代理
func startStreamProxy(stream string) error {
	// values := map[string]string{"secret": setting.MediaSetting.Secret, "vhost": "__defaultVhost__", "app": "wukong", "stream": lib.MD5(stream), "url": stream}
	// json_data, err := json.Marshal(values)

	// if err != nil {
	// 	logging.Fatal(err)
	// 	return err
	// }

	params := "secret=" + setting.MediaSetting.Secret + "&" +
		"vhost=__defaultVhost__&app=wukong&stream=" + lib.MD5(stream) + "&url=" + stream
	path := fmt.Sprintf("%s?%s", addStreamProxyUrl, params)

	resp, err := http.Get(path)

	if err != nil {
		logging.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {

		logging.Fatal(err)
	}

	return nil
}
