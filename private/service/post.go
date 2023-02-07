package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-saas/kit-layout/api"
	pb "github.com/go-saas/kit-layout/api/post/v1"
	"github.com/go-saas/kit-layout/private/biz"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostService struct {
	pb.UnimplementedPostServiceServer
	repo biz.PostRepo
	auth authz.Service
}

func NewPostService(repo biz.PostRepo, auth authz.Service) *PostService {
	return &PostService{repo: repo, auth: auth}
}

func (s *PostService) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePost, "*"), authz.ReadAction); err != nil {
		return nil, err
	}
	ret := &pb.ListPostReply{}

	totalCount, filterCount, err := s.repo.Count(ctx, req)
	ret.TotalSize = int32(totalCount)
	ret.FilterSize = int32(filterCount)

	if err != nil {
		return ret, err
	}
	items, err := s.repo.List(ctx, req)
	if err != nil {
		return ret, err
	}
	rItems := lo.Map(items, func(g *biz.Post, _ int) *pb.Post {
		b := &pb.Post{}
		MapBizPost2Pb(g, b)
		return b
	})

	ret.Items = rItems
	return ret, nil
}
func (s *PostService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePost, req.GetId()), authz.ReadAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}
	res := &pb.Post{}
	MapBizPost2Pb(g, res)
	return res, nil
}
func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePost, "*"), authz.WriteAction); err != nil {
		return nil, err
	}
	e := &biz.Post{}
	MapCreatePbPost2Biz(req, e)
	err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	res := &pb.Post{}
	MapBizPost2Pb(e, res)
	return res, nil
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePost, req.Post.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Post.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	MapUpdatePbPost2Biz(req.Post, g)
	if err := s.repo.Update(ctx, g.ID.String(), g, nil); err != nil {
		return nil, err
	}
	res := &pb.Post{}
	MapBizPost2Pb(g, res)
	return res, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	if _, err := s.auth.Check(ctx, authz.NewEntityResource(api.ResourcePost, req.Id), authz.WriteAction); err != nil {
		return nil, err
	}
	g, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, errors.NotFound("", "")
	}

	err = s.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePostReply{Id: g.ID.String(), Name: g.Name}, nil
}
func MapBizPost2Pb(a *biz.Post, b *pb.Post) {
	b.Id = a.ID.String()
	b.Name = a.Name
	b.CreatedAt = timestamppb.New(a.CreatedAt)
}

func MapUpdatePbPost2Biz(a *pb.UpdatePost, b *biz.Post) {
	b.Name = a.Name
}
func MapCreatePbPost2Biz(a *pb.CreatePostRequest, b *biz.Post) {
	b.Name = a.Name
}
