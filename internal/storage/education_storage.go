package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	IEducationStorage interface {
		IStorage[*model.Education]
		FindByAccId(c context.Context, accId int) (data []*model.Education, err error)
	}

	EducationStorage struct {
		AbstractStorage[*model.Education]
	}
)

func NewEducationStorage(s *Storage) *EducationStorage {
	return &EducationStorage{
		AbstractStorage[*model.Education]{
			db:        s.db,
			tableName: model.EducationTableName,
		},
	}
}

func (s *EducationStorage) FindByAccId(c context.Context, accId int) (data []*model.Education, err error) {
	err = s.db.WithContext(c).Table(s.tableName).Where("acc_id = ?", accId).Find(&data).Error
	return data, err
}
