package presenter

import (
	"research-api/domain"
	"research-api/usecase/port/server"
)

func (p *GRPCPresenter) FindArticle(article domain.Article) (*server.ArticleResponse, error) {
	return &server.ArticleResponse{
		ArticleId:    string(article.ArticleId),
		ArticleOverview: p.toArticleOverview(article.ArticleOverview),
		Content: article.Content,
	}, nil
}

func (p *GRPCPresenter) toArticleOverview(overview domain.ArticleOverview) server.ArticleOverview {
	return server.ArticleOverview{
		Title:        overview.Title,
		Editor:       overview.Editor,
		LastModified: overview.LastModified,
		Thumbnail:    overview.Thumbnail,
		Description:  overview.Description,
	}
}