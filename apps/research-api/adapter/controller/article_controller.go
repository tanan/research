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

func NewArticleController(c *config.Config) ArticleController {
	repo, err := database.NewSQLHandler("host=127.0.0.1 port=5432 user=dolphin dbname=research password=dolphin sslmode=disable", c.Verbose)
	if err != nil {
		log.Fatalln(err)
	}
	api, err := contentful.NewContentfulApi("9dmsobov5wik")
	if err != nil {
		log.Fatalln(err)
	}
	return ArticleController{
		InputPort: interactor.NewArticleInteractor(c, presenter.NewGRPCPresenter(c), repo, api),
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
