package storage

import (
	"pro-link-api/internal/model"
)

type (
	IAccountStorage interface {
		IStorage[*model.Account]
	}

	AccountStorage struct {
		AbstractStorage[*model.Account]
	}
)

func NewAccountStorage(s *Storage) *AccountStorage {
	return &AccountStorage{
		AbstractStorage[*model.Account]{
			db:        s.db,
			tableName: model.AccountTableName,
		},
	}
}
