package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"web_rpc/internal/data/model"
)

type Order struct {
	ID uint
	OrderNo string
	UserName string
	Amount float64
	Status string
	FileUrl string
}

type OrderRepo interface {
	CreateOrder(context.Context, *Order) error
	UpdateOrder(context.Context, *Order) error
	GetOrder(context.Context, uint) (*Order, error)
	ListOrder(context.Context, string, string, string) ([]*Order, error)
	//GetListOrder(context.Context) ([]*Order, error)
	GetListOrder(context.Context) ([]*model.OrderModel, error)
}

type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (o *OrderUsecase) CreateOrder(ctx context.Context,  order *Order) error {
	return o.repo.CreateOrder(ctx, order)
}

func (o *OrderUsecase) UpdateOrder(ctx context.Context, order *Order) error {
	return o.repo.UpdateOrder(ctx, order)
}

func (o *OrderUsecase) GetOrder(ctx context.Context, id uint) (*Order, error) {
	return o.repo.GetOrder(ctx, id)
}

func (o *OrderUsecase) ListOrder(ctx context.Context, query, sequence, by string) ([]*Order, error) {
	return o.repo.ListOrder(ctx, query, sequence, by)
}
func (o *OrderUsecase) GetListOrder(ctx context.Context) ([]*model.OrderModel, error) {
	return o.repo.GetListOrder(ctx)
}


