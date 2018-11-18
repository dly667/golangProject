package main

import (
	"Crawler/engine"
	"Crawler/zhenai/parser"
)

const targetUrl = "http://www.zhenai.com/zhenghun"
func main()  {


	engine.SimpleEngine{}.Run(engine.Request{
		Url:targetUrl,
		ParserFunc:parser.ParseCityList,
	})
}


