package svc

import (
	"context"
	"github.com/pojntfx/micro-math/math/proto"
	"log"
)

type Math struct{}

func (t *Math) Add(ctx context.Context, args *proto.MathAddArgs, reply *proto.MathAddReply) error {
	reply.Result = args.First + args.Second
	log.Println("Adding", args.First, "to", args.Second)
	return nil
}

func (t *Math) Subtract(ctx context.Context, args *proto.MathSubtractArgs, reply *proto.MathSubtractReply) error {
	reply.Result = args.Second - args.First
	log.Println("Subtracting", args.First, "from", args.Second)
	return nil
}
