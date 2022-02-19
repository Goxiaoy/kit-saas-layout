package api

import (
	grpc2 "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	v1 "github.com/goxiaoy/kit-saas-layout/api/helloworld/v1"
	"google.golang.org/grpc"
)

type GrpcConn grpc.ClientConnInterface
type HttpClient *http.Client

const ServiceName = "github.com/goxiaoy/kit-saas-layout"

func NewGrpcConn(clientName api.ClientName,services *conf.Services, opt *api.Option, tokenMgr api.TokenManager, opts ...grpc2.ClientOption) (GrpcConn, func()) {
	return api.NewGrpcConn(clientName, ServiceName, services, true, opt, tokenMgr, opts...)
}

func NewHttpClient(clientName api.ClientName,services *conf.Services, opt *api.Option, tokenMgr api.TokenManager, opts ...http.ClientOption) (HttpClient, func()) {
	return api.NewHttpClient(clientName, ServiceName, services, opt, tokenMgr, opts...)
}

var GrpcProviderSet = wire.NewSet(NewGrpcConn,NewGreetGrpcClient)
var HttpProviderSet = wire.NewSet(NewHttpClient,NewGreetHttpClient)

func NewGreetGrpcClient(conn GrpcConn) v1.GreeterClient {
	return v1.NewGreeterClient(conn)
}

func NewGreetHttpClient(http HttpClient) v1.GreeterHTTPClient {
	return v1.NewGreeterHTTPClient(http)
}
