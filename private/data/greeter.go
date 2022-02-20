package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/kit-saas-layout/private/biz"
)

type greeterRepo struct {
	Repo
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(dbProvider gorm.DbProvider, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		Repo: Repo{DbProvider: dbProvider},
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
