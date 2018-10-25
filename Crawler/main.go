package main

import (
	"Crawler/engine"
	"Crawler/zhenai/parser"
)

const targetUrl = "http://www.zhenai.com/zhenghun"
func main()  {


	engine.Run(engine.Request{
		Url:targetUrl,
		ParserFunc:parser.ParseCityList,
	})
}


