package server

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"research-api/adapter/controller"
	"research-api/domain"
	"research-api/proto/article"
)

type server struct {
	controller controller.ArticleController
}

func NewServer(s *grpc.Server, articleController controller.ArticleController) {
	server := &server{
		controller: articleController,
	}
	article.RegisterArticleServer(s, server)
	reflection.Register(s)
}

func (s *server) FindArticle(c context.Context, req *article.ArticleRequest) (*article.ArticleResponse, error) {
	res, err := s.controller.FindArticle(domain.ArticleId(req.ArticleId))
	if err != nil {
		return &article.ArticleResponse{}, err
	}
	t, err := ptypes.TimestampProto(res.LastModified)
	if err != nil {
		return &article.ArticleResponse{}, err
	}
	return &article.ArticleResponse{
		ArticleId:    res.ArticleId,
		ArticleName:  res.Title,
		Editor:       res.Editor,
		LastModified: t,
		Thumbnail:    res.Thumbnail,
		Description:  res.Description,
	}, nil
}
