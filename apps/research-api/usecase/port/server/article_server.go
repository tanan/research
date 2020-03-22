package server

import (
	"research-api/domain"
	"time"
)

type ArticleInputPort interface {
	FindArticle(id domain.ArticleId) (*ArticleResponse, error)
}

type ArticleOutputPort interface {
	FindArticle(article domain.Article) (*ArticleResponse, error)
}

type ArticleResponse struct {
	ArticleId    string
	ArticleOverview ArticleOverview
	Content string
}

type ArticleOverview struct {
	Title        string
	Editor       string
	LastModified time.Time
	Thumbnail    string
	Description  string
}