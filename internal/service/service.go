package service

import (
	"pro-link-api/internal/config"
	"pro-link-api/internal/storage"
)

type (
	Service struct {
		Config *config.Config
	}

	AuthService struct {
		*Service
	}
)

func New(storage *storage.Storage, config *config.Config) *Service {
	return &Service{
		Config: config,
	}
}

func NewAuthService(service *Service) *AuthService {
	return &AuthService{
		Service: service,
	}
}
