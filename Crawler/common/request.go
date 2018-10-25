package common

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)
// 蘑菇代理
const appKey = "UzlRTTkxNnBMbDBJbXNvMjpud2RUSlNySXFWM3hjTWhW"
const transformIp = "transfer.mogumiao.com:9001"
func Get(targetUrl string, ifProxy bool) (*http.Response ,error){

	//const targetUrl = "http://album.zhenai.com/u/1572218980"
	if targetUrl == ""{
		targetUrl = "https://ip.cn"
	}

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil{
		panic(err)
	}
	// 初始化
	client := &http.Client{}

	if ifProxy{
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,

			Proxy: func(request *http.Request) (*url.URL, error) {
				return &url.URL{
					Host:transformIp,
				},nil
			},
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},

		}
		client = &http.Client{Transport: tr}

		// 增加蘑菇代理 服务器转发
		req.Header.Add("Proxy-Authorization", `Basic `+appKey)
	}

	// 增加必要头文件
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36`)

	return client.Do(req)
	//fmt.Printf("%s",rs)

}
