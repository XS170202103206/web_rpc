// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OrderHTTPServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
	GetListOrder(context.Context, *GetListOrderRequest) (*GetListOrderReply, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderReply, error)
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/order/create", _Order_CreateOrder0_HTTP_Handler(srv))
	r.POST("/api/v1/order/update", _Order_UpdateOrder0_HTTP_Handler(srv))
	r.GET("/api/v1/order/get", _Order_GetOrder0_HTTP_Handler(srv))
	r.GET("/api/v1/order/list", _Order_ListOrder0_HTTP_Handler(srv))
	r.GET("/api/v1/order/all", _Order_GetListOrder0_HTTP_Handler(srv))
}

func _Order_CreateOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.order.v1.Order/CreateOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_UpdateOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.order.v1.Order/UpdateOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrder(ctx, req.(*UpdateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_GetOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.order.v1.Order/GetOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrder(ctx, req.(*GetOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_ListOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.order.v1.Order/ListOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_GetListOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetListOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.order.v1.Order/GetListOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetListOrder(ctx, req.(*GetListOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetListOrderReply)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	CreateOrder(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *CreateOrderReply, err error)
	GetListOrder(ctx context.Context, req *GetListOrderRequest, opts ...http.CallOption) (rsp *GetListOrderReply, err error)
	GetOrder(ctx context.Context, req *GetOrderRequest, opts ...http.CallOption) (rsp *GetOrderReply, err error)
	ListOrder(ctx context.Context, req *ListOrderRequest, opts ...http.CallOption) (rsp *ListOrderReply, err error)
	UpdateOrder(ctx context.Context, req *UpdateOrderRequest, opts ...http.CallOption) (rsp *UpdateOrderReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*CreateOrderReply, error) {
	var out CreateOrderReply
	pattern := "/api/v1/order/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.order.v1.Order/CreateOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) GetListOrder(ctx context.Context, in *GetListOrderRequest, opts ...http.CallOption) (*GetListOrderReply, error) {
	var out GetListOrderReply
	pattern := "/api/v1/order/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.order.v1.Order/GetListOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...http.CallOption) (*GetOrderReply, error) {
	var out GetOrderReply
	pattern := "/api/v1/order/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.order.v1.Order/GetOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...http.CallOption) (*ListOrderReply, error) {
	var out ListOrderReply
	pattern := "/api/v1/order/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.order.v1.Order/ListOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...http.CallOption) (*UpdateOrderReply, error) {
	var out UpdateOrderReply
	pattern := "/api/v1/order/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.order.v1.Order/UpdateOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
