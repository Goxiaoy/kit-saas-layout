package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	api2 "github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/blob"
	conf2 "github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/server"
	"github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas/common"
	shttp "github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/kratos/saas"
	"github.com/goxiaoy/kit-saas-layout/api"
	v1 "github.com/goxiaoy/kit-saas-layout/api/helloworld/v1"
	"github.com/goxiaoy/kit-saas-layout/private/conf"
	"github.com/goxiaoy/kit-saas-layout/private/service"
	uow2 "github.com/goxiaoy/uow"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf2.Services,
	sCfg *conf2.Security,
	tokenizer jwt.Tokenizer,
	ts common.TenantStore,
	uowMgr uow2.Manager,
	reqDecoder http.DecodeRequestFunc,
	resEncoder http.EncodeResponseFunc,
	errEncoder http.EncodeErrorFunc,
	factory blob.Factory,
	dataCfg *conf.Data,
	mOpt *shttp.WebMultiTenancyOption,
	apiOpt *api2.Option,
	greeter *service.GreeterService,
	logger log.Logger) *http.Server {
	var opts []http.ServerOption
	opts = server.PatchHttpOpts(logger, opts, api.ServiceName, c, sCfg, reqDecoder, resEncoder, errEncoder)
	opts = append(opts, []http.ServerOption{
		http.Middleware(
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
	}...)
	srv := http.NewServer(opts...)

	server.HandleBlobs("", dataCfg.Blobs, srv, factory)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
