package service

import (
	"context"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authz"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/goxiaoy/kit-saas-layout/api/helloworld/v1"
	"github.com/goxiaoy/kit-saas-layout/private/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc   *biz.GreeterUsecase
	log  *log.Helper
	auth authz.Service
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, auth authz.Service, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, auth: auth, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
