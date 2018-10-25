package request

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main(){
	getRequestByProxy()
}

func getRequestByProxy(){
	const appKey = "UzlRTTkxNnBMbDBJbXNvMjpud2RUSlNySXFWM3hjTWhW"
	//const targetUrl = "http://album.zhenai.com/u/1572218980"
	const targetUrl = "https://ip.cn"
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		Proxy: func(request *http.Request) (*url.URL, error) {
			return &url.URL{
				Host:"transfer.mogumiao.com:9001",
			},nil
		},
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},

	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil{
		panic(err)
	}
	// 增加蘑菇代理 服务器转发
	req.Header.Add("Proxy-Authorization", `Basic `+appKey)
	// 增加必要头文件
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36`)


	resp, err := client.Do(req)


	if err != nil{
		panic(err)
	}
	// 解析响应文件
	rs ,err:= ioutil.ReadAll(resp.Body)
	fmt.Printf("%s",rs)

}
