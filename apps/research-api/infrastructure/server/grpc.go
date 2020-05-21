package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"research-api/adapter/controller"
	"research-api/domain"
	"research-api/proto"
	p "research-api/usecase/port/server"
	"time"
)

type server struct {
	controller controller.ArticleController
}

func NewServer(s *grpc.Server, articleController controller.ArticleController) {
	server := &server{
		controller: articleController,
	}
	article.RegisterArticleServiceServer(s, server)
	reflection.Register(s)
}

func (s *server) FindArticle(c context.Context, req *article.Article) (*article.Article, error) {
	res, err := s.controller.FindArticle(domain.ArticleId(req.ArticleId))
	if err != nil {
		return &article.Article{}, err
	}
	return s.toGRPCArticleResponse(res)
}

func (s *server) FindArticles(c context.Context, req *article.Articles) (*article.Articles, error) {
	res, err := s.controller.FindArticles()
	if err != nil {
		return &article.Articles{}, err
	}
	var articlesResponse article.Articles
	for _, v := range *res {
		if err != nil {
			continue
		}
		articlesResponse.Articles = append(articlesResponse.Articles, &article.Article{
			ArticleId: v.ArticleId,
			Overview:  s.toOverview(v.ArticleOverview),
		})
	}
	return &articlesResponse, nil
}

func (s *server) FindArticleContent(c context.Context, req *article.Article) (*article.Article, error) {
	res, err := s.controller.FindArticleContent(domain.ArticleId(req.ArticleId))
	if err != nil {
		return &article.Article{}, err
	}
	return s.toGRPCArticleResponse(res)
}

func (s *server) StoreArticle(ctx context.Context, req *article.StoreArticleRequest) (*article.StoreArticleResponse, error) {
	res, err := s.controller.StoreArticle(domain.Article{
		ArticleId: domain.ArticleId(req.Article.ArticleId),
		ArticleOverview: domain.ArticleOverview{
			Title: req.Article.Overview.ArticleName,
			Editor: domain.Editor{
				Id:   int(req.Article.Overview.Editor.EditorId),
				Name: req.Article.Overview.Editor.EditorName,
				Icon: req.Article.Overview.Editor.EditorIcon,
			},
			LastModified: time.Unix(req.Article.Overview.LastModified.GetSeconds(), int64(req.Article.Overview.LastModified.GetNanos())),
			Thumbnail:    req.Article.Overview.Thumbnail,
			Description:  req.Article.Overview.Description,
		},
		Content: nil,
	})
	if err != nil {
		return &article.StoreArticleResponse{}, err
	}
	return &article.StoreArticleResponse{
		Message: res.Message,
	}, nil
}

func (s *server) StoreEditor(ctx context.Context, req *article.StoreEditorRequest) (*article.StoreEditorResponse, error) {
	res, err := s.controller.StoreArticleEditor(domain.Editor{
		Id:   int(req.Editor.EditorId),
		Name: req.Editor.EditorName,
		Icon: req.Editor.EditorIcon,
	})
	if err != nil {
		return &article.StoreEditorResponse{}, err
	}
	return &article.StoreEditorResponse{
		Editor: &article.Editor{
			EditorId:   int32(res.EditorId),
			EditorName: res.EditorName,
			EditorIcon: res.EditorIcon,
		},
	}, nil
}

func (s *server) toGRPCArticleResponse(res *p.ArticleResponse) (*article.Article, error) {
	content, err := s.toStruct(res.Content)
	if err != nil {
		return &article.Article{}, err
	}
	return &article.Article{
		ArticleId: res.ArticleId,
		Overview:  s.toOverview(res.ArticleOverview),
		Content:   content,
	}, nil
}

func (s *server) toOverview(overview p.ArticleOverview) *article.ArticleOverview {
	t, err := ptypes.TimestampProto(overview.LastModified)
	if err != nil {
		return &article.ArticleOverview{}
	}
	return &article.ArticleOverview{
		ArticleName: overview.Title,
		Editor: &article.Editor{
			EditorName: overview.Editor,
			EditorIcon: overview.EditorIcon,
		},
		LastModified: t,
		Thumbnail:    overview.Thumbnail,
		Description:  overview.Description,
	}
}

func (s *server) toStruct(msg map[string]interface{}) (*structpb.Struct, error) {

	byteArray, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(byteArray)

	pbs := &structpb.Struct{}
	if err = jsonpb.Unmarshal(reader, pbs); err != nil {
		return nil, err
	}

	return pbs, nil
}
