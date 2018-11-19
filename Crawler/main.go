package main

import (
	"Crawler/engine"
	"Crawler/scheduler"
	"Crawler/zhenai/parser"
)

const targetUrl = "http://www.zhenai.com/zhenghun"
func main()  {
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,

	}
	e.Run(engine.Request{
		Url:targetUrl,
		ParserFunc:parser.ParseCityList,
	})
}


