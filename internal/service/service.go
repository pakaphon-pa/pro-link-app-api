package service

import "pro-link-api/internal/storage"

type Service struct {
}

func New(storage *storage.Storage) *Service {
	return &Service{}
}
