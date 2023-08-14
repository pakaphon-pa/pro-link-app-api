package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	IExperienceStorage interface {
		IStorage[*model.Experience]
		FindByAccId(c context.Context, accId int) (data []*model.Experience, err error)
	}

	ExperienceStorage struct {
		AbstractStorage[*model.Experience]
	}
)

func NewExperienceStorage(s *Storage) *ExperienceStorage {
	return &ExperienceStorage{
		AbstractStorage[*model.Experience]{
			db:        s.db,
			tableName: model.ExperienceTableName,
		},
	}
}

func (s *ExperienceStorage) FindByAccId(c context.Context, accId int) (data []*model.Experience, err error) {
	err = s.db.WithContext(c).Table(s.tableName).Where("acc_id = ?", accId).Find(&data).Error
	return data, err
}
