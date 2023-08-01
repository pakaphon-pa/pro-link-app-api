package service

import (
	"pro-link-api/internal/config"
	"pro-link-api/internal/storage"
)

type Service struct {
	Config *config.Config
}

func New(storage *storage.Storage, config *config.Config) *Service {
	return &Service{
		Config: config,
	}
}
