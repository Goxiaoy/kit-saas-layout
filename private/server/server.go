package server

import (
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas/seed"
	api2 "github.com/goxiaoy/kit-saas-layout/api"
	"github.com/goxiaoy/uow"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewSeeder, wire.Value(ClientName))

var ClientName api.ClientName = api2.ServiceName

func NewSeeder( uow uow.Manager) seed.Seeder {
	var opt = seed.NewSeedOption()
	// seed host
	opt.TenantIds = []string{""}

	return seed.NewDefaultSeeder(opt.WithUow(uow), map[string]interface{}{})
}
