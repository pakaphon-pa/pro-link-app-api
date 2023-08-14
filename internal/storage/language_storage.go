package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	ILanguageStorage interface {
		IStorage[*model.Language]
		FindByAccId(c context.Context, accId int) ([]*model.Language, error)
	}

	LanguageStorage struct {
		AbstractStorage[*model.Language]
	}
)

func NewLanguageStorage(s *Storage) *LanguageStorage {
	return &LanguageStorage{
		AbstractStorage[*model.Language]{
			db:        s.db,
			tableName: model.LanguageTableName,
		},
	}
}

func (s *LanguageStorage) FindByAccId(c context.Context, accId int) (data []*model.Language, err error) {
	err = s.db.WithContext(c).Table(s.tableName).Where("acc_id = ?", accId).Find(&data).Error
	return data, err
}
