package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)
const param = "com.hgmalls.api.v1.user"
var serverIdRe = fmt.Sprintf("<td>(%s-[^<]+)</td>",param)
func main() {

	resp,err := http.Get("http://localhost:8082/registry?service="+param)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	newReader := bufio.NewReader(resp.Body)

	contents ,err:= ioutil.ReadAll(newReader)

	re := regexp.MustCompile(serverIdRe)

	match := re.FindAllSubmatch(contents,-1)

	client := &http.Client{}

	for _, item := range match{

		req, err := http.NewRequest("PUT", "http://127.0.0.1:8500/v1/agent/service/deregister/"+string(item[1]), nil)
		if err != nil{
			panic(err)
		}
		resp, err := client.Do(req)
		if err != nil{
			panic(err)
		}
		fmt.Println(resp,err)

	}

	//fmt.Printf("%s",match)
}
