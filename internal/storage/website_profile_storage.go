package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	IWebsiteProfileStorage interface {
		IStorage[*model.WebsiteProfile]
		FindByPrfId(ctx context.Context, id int) ([]*model.WebsiteProfile, error)
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

func (s *WebsiteProfileStorage) FindByPrfId(ctx context.Context, id int) (data []*model.WebsiteProfile, err error) {
	err = s.db.WithContext(ctx).Table(s.tableName).Where("prf_id = ?", id).Find(&data).Error
	return data, err
}
