package server

import (
	"research-api/domain"
	"time"
)

type ArticleInputPort interface {
	FindArticle(id domain.ArticleId) (*ArticleResponse, error)
	FindArticles() (*ArticlesResponse, error)
	FindArticleContent(id domain.ArticleId) (*ArticleResponse, error)
	StoreEditor(editor domain.Editor) (*StoreEditorResponse, error)
}

type ArticleOutputPort interface {
	FindArticle(article domain.Article) (*ArticleResponse, error)
	FindArticles(articles domain.Articles) (*ArticlesResponse, error)
	StoreEditor(editor domain.Editor) (*StoreEditorResponse, error)
}

type ArticlesResponse []ArticleResponse

type ArticleResponse struct {
	ArticleId       string
	ArticleOverview ArticleOverview
	Content         map[string]interface{}
}

type StoreEditorResponse struct {
	EditorId   int
	EditorName string
	EditorIcon string
}

type ArticleOverview struct {
	Title        string
	Editor       string
	EditorIcon   string
	LastModified time.Time
	Thumbnail    string
	Description  string
}
