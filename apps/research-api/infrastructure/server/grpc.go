package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"research-api/adapter/controller"
	"research-api/domain"
	"research-api/proto"
	p "research-api/usecase/port/server"
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
	return s.toGRPCArticleResponse(res)
}

func (s *server) FindArticles(c context.Context, req *article.ArticlesRequest) (*article.ArticlesResponse, error) {
	res, err := s.controller.FindArticles()
	if err != nil {
		return &article.ArticlesResponse{}, err
	}
	var articlesResponse article.ArticlesResponse
	for _, v := range *res {
		t, err := ptypes.TimestampProto(v.ArticleOverview.LastModified)
		if err != nil {
			continue
		}
		articlesResponse.Articles = append(articlesResponse.Articles, &article.ArticleResponse{
			ArticleId:    v.ArticleId,
			ArticleName:  v.ArticleOverview.Title,
			Editor:       v.ArticleOverview.Editor,
			LastModified: t,
			Thumbnail:    v.ArticleOverview.Thumbnail,
			Description:  v.ArticleOverview.Description,
		})
	}
	return &articlesResponse, nil
}

func (s *server) FindArticleContent(c context.Context, req *article.ArticleRequest) (*article.ArticleResponse, error) {
	res, err := s.controller.FindArticleContent(domain.ArticleId(req.ArticleId))
	if err != nil {
		return &article.ArticleResponse{}, err
	}
	fmt.Println(res.Content)
	return s.toGRPCArticleResponse(res)
}

func (s *server) toGRPCArticleResponse(res *p.ArticleResponse) (*article.ArticleResponse, error) {
	t, err := ptypes.TimestampProto(res.ArticleOverview.LastModified)
	if err != nil {
		return &article.ArticleResponse{}, err
	}
	return &article.ArticleResponse{
		ArticleId:    res.ArticleId,
		ArticleName:  res.ArticleOverview.Title,
		Editor:       res.ArticleOverview.Editor,
		LastModified: t,
		Thumbnail:    res.ArticleOverview.Thumbnail,
		Description:  res.ArticleOverview.Description,
		Content:      res.Content,
	}, nil
}
