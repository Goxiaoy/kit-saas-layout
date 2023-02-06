package service

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"github.com/go-chi/chi/v5"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	v12 "github.com/go-saas/kit-layout/api/post/v1"
	kitdi "github.com/go-saas/kit/pkg/di"
	kitgrpc "github.com/go-saas/kit/pkg/server/grpc"
	kithttp "github.com/go-saas/kit/pkg/server/http"
	"net/http"
)

//go:embed openapi/api.swagger.json
var spec []byte

// ProviderSet is service providers.
var ProviderSet = kitdi.NewSet(
	NewGrpcServerRegister,
	NewHttpServerRegister,
	NewPostService,
)

func NewHttpServerRegister(
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	post *PostService) kithttp.ServiceRegister {
	return kithttp.ServiceRegisterFunc(func(srv *khttp.Server, middleware ...middleware.Middleware) {
		v12.RegisterPostServiceHTTPServer(srv, post)

		swaggerRouter := chi.NewRouter()
		swaggerRouter.Use(
			kithttp.MiddlewareConvert(errEncoder, middleware...))
		const apiPrefix = "/v1/server/dev/swagger"
		swaggerRouter.Handle(apiPrefix+"*", http.StripPrefix(apiPrefix, swaggerui.Handler(spec)))
	})
}

func NewGrpcServerRegister(post *PostService) kitgrpc.ServiceRegister {
	return kitgrpc.ServiceRegisterFunc(func(srv *grpc.Server, middleware ...middleware.Middleware) {
		v12.RegisterPostServiceServer(srv, post)
	})
}
