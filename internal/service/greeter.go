package service

import (
	"context"

	v1 "web_rpc/api/helloworld/v1"
	"web_rpc/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// OrderService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.OrderUsecase
	log *log.Helper
}

// NewOrderService new a greeter service.
func NewGreeterService(uc *biz.OrderUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
