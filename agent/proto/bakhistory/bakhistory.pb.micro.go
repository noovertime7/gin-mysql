// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: bakhistory.proto

package bakhistory

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for History service

func NewHistoryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for History service

type HistoryService interface {
	GetHistoryList(ctx context.Context, in *HistoryListInput, opts ...client.CallOption) (*HistoryListOutput, error)
	DeleteHistory(ctx context.Context, in *HistoryIDInput, opts ...client.CallOption) (*HistoryOneMessage, error)
	GetHistoryNumInfo(ctx context.Context, in *Empty, opts ...client.CallOption) (*HistoryNumInfoOut, error)
}

type historyService struct {
	c    client.Client
	name string
}

func NewHistoryService(name string, c client.Client) HistoryService {
	return &historyService{
		c:    c,
		name: name,
	}
}

func (c *historyService) GetHistoryList(ctx context.Context, in *HistoryListInput, opts ...client.CallOption) (*HistoryListOutput, error) {
	req := c.c.NewRequest(c.name, "History.GetHistoryList", in)
	out := new(HistoryListOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyService) DeleteHistory(ctx context.Context, in *HistoryIDInput, opts ...client.CallOption) (*HistoryOneMessage, error) {
	req := c.c.NewRequest(c.name, "History.DeleteHistory", in)
	out := new(HistoryOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyService) GetHistoryNumInfo(ctx context.Context, in *Empty, opts ...client.CallOption) (*HistoryNumInfoOut, error) {
	req := c.c.NewRequest(c.name, "History.GetHistoryNumInfo", in)
	out := new(HistoryNumInfoOut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for History service

type HistoryHandler interface {
	GetHistoryList(context.Context, *HistoryListInput, *HistoryListOutput) error
	DeleteHistory(context.Context, *HistoryIDInput, *HistoryOneMessage) error
	GetHistoryNumInfo(context.Context, *Empty, *HistoryNumInfoOut) error
}

func RegisterHistoryHandler(s server.Server, hdlr HistoryHandler, opts ...server.HandlerOption) error {
	type history interface {
		GetHistoryList(ctx context.Context, in *HistoryListInput, out *HistoryListOutput) error
		DeleteHistory(ctx context.Context, in *HistoryIDInput, out *HistoryOneMessage) error
		GetHistoryNumInfo(ctx context.Context, in *Empty, out *HistoryNumInfoOut) error
	}
	type History struct {
		history
	}
	h := &historyHandler{hdlr}
	return s.Handle(s.NewHandler(&History{h}, opts...))
}

type historyHandler struct {
	HistoryHandler
}

func (h *historyHandler) GetHistoryList(ctx context.Context, in *HistoryListInput, out *HistoryListOutput) error {
	return h.HistoryHandler.GetHistoryList(ctx, in, out)
}

func (h *historyHandler) DeleteHistory(ctx context.Context, in *HistoryIDInput, out *HistoryOneMessage) error {
	return h.HistoryHandler.DeleteHistory(ctx, in, out)
}

func (h *historyHandler) GetHistoryNumInfo(ctx context.Context, in *Empty, out *HistoryNumInfoOut) error {
	return h.HistoryHandler.GetHistoryNumInfo(ctx, in, out)
}
