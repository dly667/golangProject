package fetcher

import (
	"Crawler/common"

	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte,error){
	//time.Sleep(1e11)
	// 直接Get可能会出现返回403错误
	//resp,err := http.Get(url)
	resp,err := common.Get(url,false)

	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	newReader := bufio.NewReader(resp.Body)

	//更换爬取过来的文件编码
	encode := determineEncoding(newReader)
	utf8Reader:=transform.NewReader(newReader,
		encode.NewDecoder())
	if resp.StatusCode != http.StatusOK{
		//fmt.Println("Error :status code",resp.StatusCode)
		return nil,fmt.Errorf("Wrong status code :%d",resp.StatusCode)
	}

	return ioutil.ReadAll(utf8Reader)

	//fmt.Printf("%s",all)
	//printCityList([]byte(all))
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes,err := r.Peek(1024)
	if err != nil{
		log.Println("Fetch error:%v",err)
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(bytes,"")

	return e
}


