package database

import (
	"bytes"
	"research-api/domain"
	"research-api/infrastructure/database/model"
)

func (h SQLHandler) FindById(id domain.ArticleId) (domain.Article, error) {
	var article model.Article
	db := h.Conn.Where("id = ?", id).First(&article)
	if db.Error != nil {
		return domain.Article{}, nil
	}

	var editor model.Editor
	db = h.Conn.Where("id = ?", article.Editor).First(&editor)
	if db.Error != nil {
		return domain.Article{}, nil
	}

	return domain.Article{
		ArticleId:       domain.ArticleId(article.ArticleId),
		ArticleOverview: h.toArticleOverview(article, editor),
		Content:         nil,
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
		var editor model.Editor
		db = h.Conn.Where("id = ?", v.Editor).First(&editor)
		if db.Error != nil {
			continue
		}
		articles = append(articles, domain.Article{
			ArticleId:       domain.ArticleId(v.ArticleId),
			ArticleOverview: h.toArticleOverview(v, editor),
			Content:         nil,
		})
	}
	return articles, nil
}

func (h SQLHandler) StoreArticle(article domain.Article) (int, error) {
	m := model.Article{
		ArticleId:    string(article.ArticleId),
		Title:        article.ArticleOverview.Title,
		Description:  article.ArticleOverview.Description,
		Editor:       article.ArticleOverview.Editor.Id,
		Thumbnail:    article.ArticleOverview.Thumbnail,
		LastModified: article.ArticleOverview.LastModified,
	}
	db := h.Conn.Save(&m)
	if db.Error != nil {
		return 0, db.Error
	}
	return 1, nil
}

func (h SQLHandler) StoreEditor(editor domain.Editor) (domain.Editor, error) {
	m := model.Editor{
		Name: editor.Name,
		Icon: bytes.NewBufferString(editor.Icon).Bytes(),
	}
	db := h.Conn.Create(&m)
	if db.Error != nil {
		return domain.Editor{}, db.Error
	}
	return domain.Editor{
		Id:   m.EditorId,
		Name: m.Name,
		Icon: string(m.Icon),
	}, nil
}

func (h SQLHandler) toArticleOverview(article model.Article, editor model.Editor) domain.ArticleOverview {
	return domain.ArticleOverview{
		Title: article.Title,
		Editor: domain.Editor{
			Id:   editor.EditorId,
			Name: editor.Name,
			Icon: string(editor.Icon),
		},
		LastModified: article.LastModified,
		Thumbnail:    article.Thumbnail,
		Description:  article.Description,
	}
}
