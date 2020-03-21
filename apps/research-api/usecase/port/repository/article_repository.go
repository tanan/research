package repository

import "research-api/domain"

type ArticleRepository interface {
	FindById(id domain.ArticleId) (domain.Article, error)
}