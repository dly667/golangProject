package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "myhello"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}


	//// create a new function
	//fnc := micro.NewFunction(
	//	micro.Name("greeter"),
	//)
	//
	//// init the command line
	//fnc.Init()
	//
	//// register a handler
	//fnc.Handle(new(Greeter))
	//
	//// run the function
	//fnc.Run()
}
