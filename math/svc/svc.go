package svc

import (
	"context"
	"fmt"
	proto "github.com/pojntfx/micro-math/math/svc/proto/math"
)

type Math struct{}

func (t *Math) Add(ctx context.Context, args *proto.MathAddArgs, reply *proto.MathAddReply) error {
	reply.Result = args.First + args.Second
	fmt.Println("Adding", args.First, args.Second, reply.Result)
	return nil
}

func (t *Math) Subtract(ctx context.Context, args *proto.MathSubtractArgs, reply *proto.MathSubtractReply) error {
	reply.Result = args.Second - args.First
	return nil
}
