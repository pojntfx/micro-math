package main

import (
	"github.com/micro/go-micro"
	"github.com/pojntfx/micro-math/math/api"
	"github.com/pojntfx/micro-math/math/proto"
	"log"
)

func main() {
	server := micro.NewService(
		micro.Name("go.micro.api.math"),
	)

	server.Init()

	server.Server().Handle(server.Server().NewHandler(&api.Math{
		Client: proto.NewMathService("go.micro.srv.math", server.Client()),
	}))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
