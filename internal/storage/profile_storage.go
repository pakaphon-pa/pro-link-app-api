package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	IProfileStorage interface {
		IStorage[*model.Profile]
		FindByAccId(ctx context.Context, id int) (*model.Profile, error)
	}

	ProfileStorage struct {
		AbstractStorage[*model.Profile]
	}
)

func NewProfileStorage(s *Storage) *ProfileStorage {
	return &ProfileStorage{
		AbstractStorage[*model.Profile]{
			db:        s.db,
			tableName: model.ProfileTableName,
		},
	}
}

func (s *ProfileStorage) FindByAccId(ctx context.Context, id int) (data *model.Profile, err error) {
	err = s.db.WithContext(ctx).Table(s.tableName).Where("acc_id = ?", id).Find(&data).Error
	return data, err
}
