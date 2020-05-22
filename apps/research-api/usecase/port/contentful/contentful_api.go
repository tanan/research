package contentful

import "research-api/domain"

type ContentfulApi interface {
	FindById(id domain.ArticleId) (map[string]interface{}, error)
	FindThumbnailUrl(id domain.ThumbnailId) (domain.ThumbnailUrl, error)
}
