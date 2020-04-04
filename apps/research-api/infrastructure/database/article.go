package database

import (
	"research-api/domain"
	"research-api/infrastructure/database/model"
)

func (h SQLHandler) FindById(id domain.ArticleId) (domain.Article, error) {
	var m model.Article
	db := h.Conn.Where("id = ?", id).First(&m)
	if db.Error != nil {
		return domain.Article{}, nil
	}
	return domain.Article{
		ArticleId:       domain.ArticleId(m.ArticleId),
		ArticleOverview: h.toArticleOverview(m),
		Content:         "",
	}, nil
}

func (h SQLHandler) FindLatest(size int) (domain.Articles, error) {
	var m []model.Article
	db := h.Conn.Order("last_modified desc").Limit(size).Find(&m)
	if db.Error != nil {
		return domain.Articles{}, db.Error
	}
	var articles domain.Articles
	for _, v := range m {
		articles = append(articles, domain.Article{
			ArticleId:       domain.ArticleId(v.ArticleId),
			ArticleOverview: h.toArticleOverview(v),
			Content:         "",
		})
	}
	return articles, nil
}

func (h SQLHandler) toArticleOverview(m model.Article) domain.ArticleOverview {
	return domain.ArticleOverview{
		Title:        m.Title,
		Editor:       m.Editor,
		LastModified: m.LastModified,
		Thumbnail:    m.Thumbnail,
		Description:  m.Description,
	}
}
