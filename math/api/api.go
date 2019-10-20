package api

import (
	"context"
	"encoding/json"
	apiProto "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/pojntfx/micro-math/math/proto"
	"log"
	"strconv"
)

type Math struct {
	Client proto.MathService
}

func (m *Math) Add(ctx context.Context, args *apiProto.Request, reply *apiProto.Response) error {
	log.Print("Received Math.Add request")

	first, ok := args.Get["first"]
	if !ok {
		return errors.BadRequest("space.pojtinger.felicitas.api.math", "`first` cannot be blank")
	}

	second, ok := args.Get["second"]
	if !ok {
		return errors.BadRequest("space.pojtinger.felicitas.api.math", "`second` cannot be blank")
	}

	firstAsInt, _ := strconv.ParseInt(first.Values[0], 0, 64)
	secondAsInt, _ := strconv.ParseInt(second.Values[0], 0, 64)

	response, err := m.Client.Add(ctx, &proto.MathAddArgs{
		First:  firstAsInt,
		Second: secondAsInt,
	})
	if err != nil {
		return errors.InternalServerError("space.pojtinger.felicitas.api.math", "Could not call `space.pojtinger.felicitas.svc.math`", err)
	}

	reply.StatusCode = 200

	b, _ := json.Marshal(map[string]int64{
		"result": response.Result,
	})

	reply.Body = string(b)

	return nil
}

func (m *Math) Subtract(ctx context.Context, args *apiProto.Request, reply *apiProto.Response) error {
	log.Print("Received Math.Subtract request")

	first, ok := args.Get["first"]
	if !ok {
		return errors.BadRequest("space.pojtinger.felicitas.api.math", "`first` cannot be blank")
	}

	second, ok := args.Get["second"]
	if !ok {
		return errors.BadRequest("space.pojtinger.felicitas.api.math", "`second` cannot be blank")
	}

	firstAsInt, _ := strconv.ParseInt(first.Values[0], 0, 64)
	secondAsInt, _ := strconv.ParseInt(second.Values[0], 0, 64)

	response, err := m.Client.Subtract(ctx, &proto.MathSubtractArgs{
		First:  secondAsInt,
		Second: firstAsInt,
	})
	if err != nil {
		return errors.InternalServerError("space.pojtinger.felicitas.api.math", "Could not call `space.pojtinger.felicitas.svc.math`", err)
	}
	reply.StatusCode = 200

	b, _ := json.Marshal(map[string]int64{
		"result": response.Result,
	})

	reply.Body = string(b)

	return nil
}
