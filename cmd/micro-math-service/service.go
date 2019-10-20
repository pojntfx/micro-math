package main

import (
	"github.com/micro/go-micro"
	"github.com/pojntfx/micro-math/math/svc"
	proto "github.com/pojntfx/micro-math/math/svc/proto/math"
	"log"
	"time"
)

func main() {
	server := micro.NewService(
		micro.Name("space.pojtinger.felicitas.api.svc"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	server.Init()

	proto.RegisterMathHandler(server.Server(), new(svc.Math))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
