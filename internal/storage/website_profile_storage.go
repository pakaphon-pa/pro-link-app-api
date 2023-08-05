package storage

import (
	"pro-link-api/internal/model"
)

type (
	IWebsiteProfileStorage interface {
		IStorage[*model.WebsiteProfile]
	}

	WebsiteProfileStorage struct {
		AbstractStorage[*model.WebsiteProfile]
	}
)

func NewWebsiteProfileStorage(s *Storage) *WebsiteProfileStorage {
	return &WebsiteProfileStorage{
		AbstractStorage[*model.WebsiteProfile]{
			db:        s.db,
			tableName: model.WebsiteProfileTableName,
		},
	}
}
