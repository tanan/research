package presenter

import (
	"research-api/domain"
	"research-api/usecase/port/server"
)

func (p *GRPCPresenter) FindArticle(article domain.Article) (*server.ArticleResponse, error) {
	return &server.ArticleResponse{
		ArticleId:       string(article.ArticleId),
		ArticleOverview: p.toArticleOverview(article.ArticleOverview),
		Content:         article.Content,
	}, nil
}

func (p *GRPCPresenter) FindArticles(articles domain.Articles) (*server.ArticlesResponse, error) {
	var res server.ArticlesResponse
	for _, v := range articles {
		res = append(res, server.ArticleResponse{
			ArticleId:       string(v.ArticleId),
			ArticleOverview: p.toArticleOverview(v.ArticleOverview),
			Content:         v.Content,
		})
	}
	return &res, nil
}

func (p *GRPCPresenter) StoreEditor(editor domain.Editor) (*server.StoreEditorResponse, error) {
	return &server.StoreEditorResponse{
		EditorId:   editor.Id,
		EditorName: editor.Name,
		EditorIcon: editor.Icon,
	}, nil
}

func (p *GRPCPresenter) toArticleOverview(overview domain.ArticleOverview) server.ArticleOverview {
	return server.ArticleOverview{
		Title:        overview.Title,
		Editor:       overview.Editor.Name,
		EditorIcon:   overview.Editor.Icon,
		LastModified: overview.LastModified,
		Thumbnail:    overview.Thumbnail,
		Description:  overview.Description,
	}
}
