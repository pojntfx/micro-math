// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: math/proto/proto.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Math service

type MathService interface {
	Add(ctx context.Context, in *MathAddArgs, opts ...client.CallOption) (*MathAddReply, error)
	Subtract(ctx context.Context, in *MathSubtractArgs, opts ...client.CallOption) (*MathSubtractReply, error)
}

type mathService struct {
	c    client.Client
	name string
}

func NewMathService(name string, c client.Client) MathService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &mathService{
		c:    c,
		name: name,
	}
}

func (c *mathService) Add(ctx context.Context, in *MathAddArgs, opts ...client.CallOption) (*MathAddReply, error) {
	req := c.c.NewRequest(c.name, "Math.Add", in)
	out := new(MathAddReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathService) Subtract(ctx context.Context, in *MathSubtractArgs, opts ...client.CallOption) (*MathSubtractReply, error) {
	req := c.c.NewRequest(c.name, "Math.Subtract", in)
	out := new(MathSubtractReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Math service

type MathHandler interface {
	Add(context.Context, *MathAddArgs, *MathAddReply) error
	Subtract(context.Context, *MathSubtractArgs, *MathSubtractReply) error
}

func RegisterMathHandler(s server.Server, hdlr MathHandler, opts ...server.HandlerOption) error {
	type math interface {
		Add(ctx context.Context, in *MathAddArgs, out *MathAddReply) error
		Subtract(ctx context.Context, in *MathSubtractArgs, out *MathSubtractReply) error
	}
	type Math struct {
		math
	}
	h := &mathHandler{hdlr}
	return s.Handle(s.NewHandler(&Math{h}, opts...))
}

type mathHandler struct {
	MathHandler
}

func (h *mathHandler) Add(ctx context.Context, in *MathAddArgs, out *MathAddReply) error {
	return h.MathHandler.Add(ctx, in, out)
}

func (h *mathHandler) Subtract(ctx context.Context, in *MathSubtractArgs, out *MathSubtractReply) error {
	return h.MathHandler.Subtract(ctx, in, out)
}
