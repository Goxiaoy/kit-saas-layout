package biz

import (
	"github.com/goxiaoy/go-saas-kit/pkg/data"
	"github.com/goxiaoy/go-saas-kit/pkg/gorm"
	v1 "github.com/goxiaoy/kit-saas-layout/api/post/v1"
)

type Post struct {
	gorm.UIDBase
	gorm.AuditedModel
	Name string
}

type PostRepo interface {
	data.Repo[Post, string, v1.ListPostRequest]
}
