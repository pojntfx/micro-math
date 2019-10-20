package api

import (
	"context"
	"encoding/json"
	apiProto "github.com/micro/go-micro/api/proto"
	proto "github.com/pojntfx/micro-math/math/svc/proto/math"
	"log"
	"strconv"
)

type Math struct {
	Client proto.MathService
}

func (m *Math) Add(ctx context.Context, args *apiProto.Request, reply *apiProto.Response) error {
	log.Print("Received Math.Add request")

	first, _ := args.Get["first"]
	second, _ := args.Get["second"]

	firstAsInt, _ := strconv.ParseInt(first.Values[0], 0, 64)
	secondAsInt, _ := strconv.ParseInt(second.Values[0], 0, 64)

	response, _ := m.Client.Add(ctx, &proto.MathAddArgs{
		First:  firstAsInt,
		Second: secondAsInt,
	})

	reply.StatusCode = 200

	b, _ := json.Marshal(map[string]int64{
		"result": response.Result,
	})

	reply.Body = string(b)

	return nil
}
