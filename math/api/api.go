package api

import (
	"context"
	apiProto "github.com/micro/go-micro/api/proto"
	"github.com/pojntfx/micro-math/math/proto"
	"log"
	"strconv"
)

type Math struct {
	Client proto.MathService
}

func (m *Math) Add(ctx *context.Context, args *apiProto.Request, reply *apiProto.Response) error {
	log.Print("Received Math.Add request")

	first, _ := args.Get["firstName"]
	second, _ := args.Get["secondName"]

	firstAsInt, _ := strconv.ParseInt(first.Values[0], 0, 64)
	secondAsInt, _ := strconv.ParseInt(second.Values[0], 0, 64)

	response, _ := m.Client.Add(*ctx, &proto.MathAddArgs{
		First:  firstAsInt,
		Second: secondAsInt,
	})

	reply.StatusCode = 200

	reply.Body = string(response.Result)

	return nil
}
