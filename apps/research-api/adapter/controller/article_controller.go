package controller

import (
	"log"
	"research-api/adapter/presenter"
	"research-api/config"
	"research-api/domain"
	"research-api/infrastructure/contentful"
	"research-api/infrastructure/database"
	"research-api/usecase/interactor"
	"research-api/usecase/port/server"
)

type ArticleController struct {
	InputPort server.ArticleInputPort
}

func NewArticleController() ArticleController {
	repo, err := database.NewSQLHandler(config.GetDatabaseConnInfo(), config.GetVerbose())
	if err != nil {
		log.Fatalln(err)
	}
	api, err := contentful.NewContentfulApi("9dmsobov5wik")
	if err != nil {
		log.Fatalln(err)
	}
	return ArticleController{
		InputPort: interactor.NewArticleInteractor(presenter.NewGRPCPresenter(), repo, api),
	}
}

func (c *ArticleController) FindArticle(id domain.ArticleId) (*server.ArticleResponse, error) {
	return c.InputPort.FindArticle(id)
}

func (c *ArticleController) FindArticles() (*server.ArticlesResponse, error) {
	return c.InputPort.FindArticles()
}

func (c *ArticleController) FindArticleContent(id domain.ArticleId) (*server.ArticleResponse, error) {
	return c.InputPort.FindArticleContent(id)
}

func (c *ArticleController) StoreArticleEditor(editor domain.Editor) (*server.StoreEditorResponse, error) {
	return c.InputPort.StoreEditor(editor)
}

func (c *ArticleController) StoreArticle(article domain.Article) (*server.StoreArticleResponse, error) {
	return c.InputPort.StoreArticle(article)
}
