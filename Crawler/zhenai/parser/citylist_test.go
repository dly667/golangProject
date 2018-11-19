package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents ,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("test_citylist_html.html")
	if  err != nil{
		panic(err)
	}

	//fmt.Printf("%s",contents)
	result := ParseCityList(contents)


	const resultSize  = 470
	expectedUrls  := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedClients  := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	for i,url:=range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d: %s;but"+"was %s",
				i,url, result.Requests[i].Url)
		}
	}
	for i,city:=range expectedClients{
		if result.Items[i] != city{
			t.Errorf("expected items #%d: %s;but"+"was %s",
				i,city, result.Items[i])
		}
	}
	if len(result.Requests) != resultSize{
		t.Errorf("result should hava %d"+"requests;but had %d",resultSize,len(result.Requests))
	}

	if len(result.Items) != resultSize{
		t.Errorf("result should hava %d"+"requests;but had %d",resultSize,len(result.Items))
	}

}
