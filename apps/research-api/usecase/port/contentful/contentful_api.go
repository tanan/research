package contentful

import "research-api/domain"

type ContentfulApi interface {
	FindById(id domain.ArticleId) (fields map[string]interface{}, includes map[string]interface{}, err error)
	FindThumbnailUrl(id domain.ThumbnailId) (domain.ThumbnailUrl, error)
}
