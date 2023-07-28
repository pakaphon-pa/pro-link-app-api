package service

import "pro-link-api/internal/storage"

type Service struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
