package presenter

import (
	"research-api/usecase/port/server"
)

type GRPCPresenter struct {
	//cfg *config.Config
	server.ArticleOutputPort
}

func NewGRPCPresenter() *GRPCPresenter {
	return &GRPCPresenter{
		//cfg:               cfg,
	}
}
