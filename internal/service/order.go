package service

import (
	"context"
	"errors"
	v1 "web_rpc/api/order/v1"
	"web_rpc/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// OrderService is a greeter service.
type OrderService struct {
	v1.UnimplementedOrderServer
	uc        *biz.OrderUsecase
	log       *log.Helper
}

// NewOrderService new a greeter service.
func NewOrderService(uc *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{
		uc: uc,
		log: log.NewHelper(logger),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (*v1.CreateOrderReply, error) {
	//userName := req.GetUserName()
	//if userName == "" {
	//	return &v1.CreateOrderReply{}, errors.New("user name can not be empty")
	//}
	order := biz.Order{
		UserName:  req.GetUserName(),
		OrderNo:   req.GetOrderNo(),
		Amount:    req.GetAmount(),
		Status:    req.GetStatus(),
		FileUrl:   req.GetFileUrl(),
	}
	return &v1.CreateOrderReply{}, s.uc.CreateOrder(ctx, &order)
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *v1.UpdateOrderRequest) (*v1.UpdateOrderReply, error) {
	id := req.GetId()
	if id <= 0 {
		return  &v1.UpdateOrderReply{}, nil
	}
	_, err := s.uc.GetOrder(ctx, uint(id))
	if err != nil {
		return &v1.UpdateOrderReply{}, err
	}
	order := biz.Order{
		ID:        uint(id),
		UserName:  req.GetUserName(),
		OrderNo:   req.GetOrderNo(),
		Amount:    req.GetAmount(),
		Status:    req.GetStatus(),
		FileUrl:   req.GetFileUrl(),
	}
	return &v1.UpdateOrderReply{}, s.uc.UpdateOrder(ctx, &order)
}

func (s *OrderService) GetOrder(ctx context.Context, req *v1.GetOrderRequest) (*v1.GetOrderReply, error) {
	var reply v1.GetOrderReply
	id := req.GetId()

	if id < 0 {
		return &reply, errors.New("bad id")
	}
	order, err := s.uc.GetOrder(ctx, uint(id))
	if err != nil {
		return &reply, err
	}

	reply.Order = &v1.OrderModel{
		Id:      uint32(order.ID),
		Name:    order.UserName,
		OrderNo: order.OrderNo,
		Amount:  order.Amount,
		Status:  order.Status,
		FileUrl: order.FileUrl,
	}

	return &reply, nil
}

func (s *OrderService)ListOrder(ctx context.Context, req *v1.ListOrderRequest) (*v1.ListOrderReply, error)  {
	sequence := req.GetSequence()
	if sequence != "AMOUNT" && sequence != "CREATED" {
		return &v1.ListOrderReply{}, nil
	}
	if sequence == "CREATED" {
		sequence = "created_at"
	}

	by := req.GetBy()
	if by != "DESC" && by != "ASC" {
		return &v1.ListOrderReply{}, nil
	}
	query := req.GetQuery()

	orders, err := s.uc.ListOrder(ctx, query, sequence, by)
	if err != nil {
		return &v1.ListOrderReply{}, nil
	}

	var reply v1.ListOrderReply
	reply.Orders = make([]*v1.OrderModel, len(orders))
	for i, v := range orders {
		reply.Orders[i] = &v1.OrderModel{
			Id:      uint32(v.ID),
			Name:    v.UserName,
			OrderNo: v.OrderNo,
			Amount:  v.Amount,
			Status:  v.Status,
			FileUrl: v.FileUrl,
		}
	}
	return &reply, nil
}

func (s *OrderService)GetListOrder(ctx context.Context, _ *v1.GetListOrderRequest) (*v1.GetListOrderReply, error) {
	orders, err := s.uc.GetListOrder(ctx)
	if err != nil {
		return &v1.GetListOrderReply{}, nil
	}

	var reply v1.GetListOrderReply
	reply.Orders = make([]*v1.OrderModel, len(orders))
	for i, v := range orders {
		reply.Orders[i] = &v1.OrderModel{
			Id:      uint32(v.ID),
			Name:    v.UserName,
			OrderNo: v.OrderNo,
			Amount:  v.Amount,
			Status:  v.Status,
			FileUrl: v.FileUrl,
		}
	}
	return &reply, nil
}
