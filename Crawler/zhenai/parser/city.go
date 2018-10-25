package parser

import (
	"Crawler/engine"
	"regexp"
)

const CityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`
func ParserCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	re := regexp.MustCompile(CityRe)
	//match := re.FindAll(contents,-1)
	match := re.FindAllSubmatch(contents,-1)

	for _,item := range match{
		name := item[2]
		result.Items = append(result.Items,string(item[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(item[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return Profile(bytes, string(name))
			},
		})

		//fmt.Printf("User: %s Url: %s\n",item[2],item[1])
	}
	return result

}
