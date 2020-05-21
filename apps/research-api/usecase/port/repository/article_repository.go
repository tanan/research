package repository

import "research-api/domain"

type ArticleRepository interface {
	FindById(id domain.ArticleId) (domain.Article, error)
	FindLatest(size int) (domain.Articles, error)
	StoreArticle(article domain.Article) (int, error)
	StoreEditor(editor domain.Editor) (domain.Editor, error)
}
