package interactor

import (
	"research-api/domain"
	"research-api/usecase/port/contentful"
	"research-api/usecase/port/repository"
	"research-api/usecase/port/server"
)

type ArticleInteractor struct {
	//Config            *config.Config
	OutputPort        server.ArticleOutputPort
	ArticleRepository repository.ArticleRepository
	ContentfulApi     contentful.ContentfulApi
}

func NewArticleInteractor(
	//config *config.Config,
	outputPort server.ArticleOutputPort,
	articleRepository repository.ArticleRepository,
	contentfulApi contentful.ContentfulApi,
) ArticleInteractor {
	return ArticleInteractor{
		//Config:            config,
		OutputPort:        outputPort,
		ArticleRepository: articleRepository,
		ContentfulApi:     contentfulApi,
	}
}

func (i ArticleInteractor) FindArticle(id domain.ArticleId) (*server.ArticleResponse, error) {
	article, err := i.ArticleRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.FindArticle(article)
}

func (i ArticleInteractor) FindArticles() (*server.ArticlesResponse, error) {
	articles, err := i.ArticleRepository.FindLatest(10)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.FindArticles(articles)
}

func (i ArticleInteractor) FindArticleContent(id domain.ArticleId) (*server.ArticleResponse, error) {
	article, err := i.ArticleRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	content, err := i.ContentfulApi.FindById(id)
	if err != nil {
		return nil, err
	}
	article.Content = content
	return i.OutputPort.FindArticle(article)
}

func (i ArticleInteractor) StoreArticle(article domain.Article) (*server.StoreArticleResponse, error) {
	return i.OutputPort.StoreArticle(i.ArticleRepository.StoreArticle(article))
}

func (i ArticleInteractor) StoreEditor(e domain.Editor) (*server.StoreEditorResponse, error) {
	editor, err := i.ArticleRepository.StoreEditor(e)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.StoreEditor(editor)
}
