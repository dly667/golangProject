package main

import (
	"github.com/micro/go-log"
	"mytest/handle"

	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("cn.hmalls.api.v1.mytest1"),
		micro.Version("latest"),
	)

	// Register Handler
	service.Server().Handle(service.Server().NewHandler(&handle.User{}))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		//micro.WrapHandler(wrapper.LoginWrapper),
	)


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

