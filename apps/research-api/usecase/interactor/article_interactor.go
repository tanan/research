package interactor

import (
	"research-api/config"
	"research-api/domain"
	"research-api/usecase/port/repository"
	"research-api/usecase/port/server"
)

type ArticleInteractor struct {
	Config            *config.Config
	OutputPort        server.ArticleOutputPort
	ArticleRepository repository.ArticleRepository
}

func NewArticleInteractor(
	config *config.Config,
	outputPort server.ArticleOutputPort,
	articleRepository repository.ArticleRepository,
) ArticleInteractor {
	return ArticleInteractor{
		Config:            config,
		OutputPort:        outputPort,
		ArticleRepository: articleRepository,
	}
}

func (i ArticleInteractor) FindArticle(id domain.ArticleId) (*server.ArticleResponse, error) {
	article, err := i.ArticleRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.FindArticle(article)
}
