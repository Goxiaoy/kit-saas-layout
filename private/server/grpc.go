package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	api2 "github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	conf2 "github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/server"
	"github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/kratos/saas"
	"github.com/goxiaoy/kit-saas-layout/api"
	v1 "github.com/goxiaoy/kit-saas-layout/api/helloworld/v1"
	"github.com/goxiaoy/kit-saas-layout/private/service"
	uow2 "github.com/goxiaoy/uow"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf2.Services, tokenizer jwt.Tokenizer, ts common.TenantStore, uowMgr uow2.Manager,
	mOpt *http.WebMultiTenancyOption, apiOpt *api2.Option, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			jwt.ServerExtractAndAuth(tokenizer, logger),
			saas.Server(mOpt, ts),
			api2.ServerMiddleware(apiOpt),
			uow.Uow(logger, uowMgr),
		),
	}
	opts = server.PatchGrpcOpts(logger, opts, api.ServiceName, c)
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
