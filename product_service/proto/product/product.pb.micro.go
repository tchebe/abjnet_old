// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product/product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProductService service

type ProductService interface {
	Get(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetClientProducts(ctx context.Context, in *Client, opts ...client.CallOption) (*Response, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) Get(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetClientProducts(ctx context.Context, in *Client, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetClientProducts", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	Get(context.Context, *Product, *Response) error
	GetAll(context.Context, *Request, *Response) error
	GetClientProducts(context.Context, *Client, *Response) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		Get(ctx context.Context, in *Product, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
		GetClientProducts(ctx context.Context, in *Client, out *Response) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) Get(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.Get(ctx, in, out)
}

func (h *productServiceHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.ProductServiceHandler.GetAll(ctx, in, out)
}

func (h *productServiceHandler) GetClientProducts(ctx context.Context, in *Client, out *Response) error {
	return h.ProductServiceHandler.GetClientProducts(ctx, in, out)
}
