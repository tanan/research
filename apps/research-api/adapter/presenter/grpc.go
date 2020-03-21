package presenter

import (
	"research-api/config"
	"research-api/usecase/port/server"
)

type GRPCPresenter struct {
	cfg *config.Config
	server.ArticleOutputPort
}

func NewGRPCPresenter(cfg *config.Config) *GRPCPresenter {
	return &GRPCPresenter{
		cfg:               cfg,
	}
}