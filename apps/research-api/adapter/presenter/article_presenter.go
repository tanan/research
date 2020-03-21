package presenter

import (
	"research-api/domain"
	"research-api/usecase/port/server"
)

func (p *GRPCPresenter) FindArticle(article domain.Article) (*server.ArticleResponse, error) {
	return &server.ArticleResponse{
		ArticleId:    string(article.ArticleId),
		Editor:       article.ArticleOverview.Editor,
		LastModified: article.ArticleOverview.LastModified,
		Thumbnail:    article.ArticleOverview.Thumbnail,
		Description:  article.ArticleOverview.Description,
	}, nil
}
