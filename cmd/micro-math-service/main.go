package main

import (
	"github.com/micro/go-micro"
	"github.com/pojntfx/micro-math/math/proto"
	"github.com/pojntfx/micro-math/math/service"
	"log"
	"time"
)

func main() {
	microService := micro.NewService(
		micro.Name("com.github.pojntfx.micro-math.math.svc"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	microService.Init()

	// Register Handlers
	proto.RegisterMathHandler(microService.Server(), new(service.Math))

	// Run server
	if err := microService.Run(); err != nil {
		log.Fatal(err)
	}
}
