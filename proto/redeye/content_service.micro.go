// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/redeye/content_service.proto

package redeye

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ContentService service

type ContentServiceClient interface {
	Create(ctx context.Context, in *CreateContentRequest, opts ...client.CallOption) (*CreateContentResponse, error)
	Get(ctx context.Context, in *GetContentRequest, opts ...client.CallOption) (*GetContentResponse, error)
	List(ctx context.Context, in *ListContentRequest, opts ...client.CallOption) (*ListContentResponse, error)
	Update(ctx context.Context, in *UpdateContentRequest, opts ...client.CallOption) (*UpdateContentResponse, error)
	Delete(ctx context.Context, in *DeleteContentRequest, opts ...client.CallOption) (*DeleteContentResponse, error)
}

type contentServiceClient struct {
	c           client.Client
	serviceName string
}

func NewContentServiceClient(serviceName string, c client.Client) ContentServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "video.quartz.famethyst.srv.redeye"
	}
	return &contentServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *contentServiceClient) Create(ctx context.Context, in *CreateContentRequest, opts ...client.CallOption) (*CreateContentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ContentService.Create", in)
	out := new(CreateContentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) Get(ctx context.Context, in *GetContentRequest, opts ...client.CallOption) (*GetContentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ContentService.Get", in)
	out := new(GetContentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) List(ctx context.Context, in *ListContentRequest, opts ...client.CallOption) (*ListContentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ContentService.List", in)
	out := new(ListContentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) Update(ctx context.Context, in *UpdateContentRequest, opts ...client.CallOption) (*UpdateContentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ContentService.Update", in)
	out := new(UpdateContentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) Delete(ctx context.Context, in *DeleteContentRequest, opts ...client.CallOption) (*DeleteContentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ContentService.Delete", in)
	out := new(DeleteContentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentService service

type ContentServiceHandler interface {
	Create(context.Context, *CreateContentRequest, *CreateContentResponse) error
	Get(context.Context, *GetContentRequest, *GetContentResponse) error
	List(context.Context, *ListContentRequest, *ListContentResponse) error
	Update(context.Context, *UpdateContentRequest, *UpdateContentResponse) error
	Delete(context.Context, *DeleteContentRequest, *DeleteContentResponse) error
}

func RegisterContentServiceHandler(s server.Server, hdlr ContentServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ContentService{hdlr}, opts...))
}

type ContentService struct {
	ContentServiceHandler
}

func (h *ContentService) Create(ctx context.Context, in *CreateContentRequest, out *CreateContentResponse) error {
	return h.ContentServiceHandler.Create(ctx, in, out)
}

func (h *ContentService) Get(ctx context.Context, in *GetContentRequest, out *GetContentResponse) error {
	return h.ContentServiceHandler.Get(ctx, in, out)
}

func (h *ContentService) List(ctx context.Context, in *ListContentRequest, out *ListContentResponse) error {
	return h.ContentServiceHandler.List(ctx, in, out)
}

func (h *ContentService) Update(ctx context.Context, in *UpdateContentRequest, out *UpdateContentResponse) error {
	return h.ContentServiceHandler.Update(ctx, in, out)
}

func (h *ContentService) Delete(ctx context.Context, in *DeleteContentRequest, out *DeleteContentResponse) error {
	return h.ContentServiceHandler.Delete(ctx, in, out)
}
