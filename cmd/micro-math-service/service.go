package main

import (
	"github.com/micro/go-micro"
	"github.com/pojntfx/micro-math/math/proto"
	"github.com/pojntfx/micro-math/math/service"
	"log"
	"time"
)

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.math"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	server.Init()

	proto.RegisterMathHandler(server.Server(), new(service.Math))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
