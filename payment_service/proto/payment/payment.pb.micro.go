
// source: proto/payment/payment.proto

package payment

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

// Client API for PaymentService service

type PaymentService interface {
	Pay(ctx context.Context, in *Payment, opts ...client.CallOption) (*Response, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Pay(ctx context.Context, in *Payment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PaymentService.Pay", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentService service

type PaymentServiceHandler interface {
	Pay(context.Context, *Payment, *Response) error
}

func RegisterPaymentServiceHandler(s server.Server, hdlr PaymentServiceHandler, opts ...server.HandlerOption) error {
	type paymentService interface {
		Pay(ctx context.Context, in *Payment, out *Response) error
	}
	type PaymentService struct {
		paymentService
	}
	h := &paymentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PaymentService{h}, opts...))
}

type paymentServiceHandler struct {
	PaymentServiceHandler
}

func (h *paymentServiceHandler) Pay(ctx context.Context, in *Payment, out *Response) error {
	return h.PaymentServiceHandler.Pay(ctx, in, out)
}
