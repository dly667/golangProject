package parser

import (
	"Crawler/engine"
	"regexp"


)

const CityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]+>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult{
	result := engine.ParseResult{}
	re := regexp.MustCompile(CityListRe)
	//match := re.FindAll(contents,-1)
	match := re.FindAllSubmatch(contents,-1)

	for _,item := range match{
		result.Items = append(result.Items,string(item[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(item[1]),
			ParserFunc:ParserCity,

		})
		//fmt.Printf("City: %s Url: %s\n",item[2],item[1])
		break
	}
	return result
}
