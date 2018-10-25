package engine

import (
	"Crawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	//requests = make([]Request)
	for _,r:=range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0{
		r:= requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher:error"+" fetching url %s:%v \n",r.Url,err)


		}
		ParseResult := r.ParserFunc(body)

		requests = append(requests, ParseResult.Requests...)
		for _,item:=range ParseResult.Items{
			log.Printf("Got item %v",item)
		}

	}

}