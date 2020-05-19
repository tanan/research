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
			EditorIcon:   v.ArticleOverview.EditorIcon,
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
	return s.toGRPCArticleResponse(res)
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

func (s *server) toGRPCArticleResponse(res *p.ArticleResponse) (*article.ArticleResponse, error) {
	t, err := ptypes.TimestampProto(res.ArticleOverview.LastModified)
	content, err := s.toStruct(res.Content)
	if err != nil {
		return &article.ArticleResponse{}, err
	}
	return &article.ArticleResponse{
		ArticleId:    res.ArticleId,
		ArticleName:  res.ArticleOverview.Title,
		Editor:       res.ArticleOverview.Editor,
		EditorIcon:   res.ArticleOverview.EditorIcon,
		LastModified: t,
		Thumbnail:    res.ArticleOverview.Thumbnail,
		Description:  res.ArticleOverview.Description,
		Content:      content,
	}, nil
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
