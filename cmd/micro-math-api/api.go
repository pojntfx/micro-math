package main

import (
	"github.com/micro/go-micro"
	"github.com/pojntfx/micro-math/math/api"
	proto "github.com/pojntfx/micro-math/math/svc/proto/math"
	"log"
)

func main() {
	server := micro.NewService(
		micro.Name("space.pojtinger.felicitas.api.math"),
	)

	server.Init()

	server.Server().Handle(
		server.Server().NewHandler(
			&api.Math{Client: proto.NewMathService("space.pojtinger.felicitas.api.svc", server.Client())},
		),
	)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}