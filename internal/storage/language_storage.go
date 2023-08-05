package storage

import (
	"pro-link-api/internal/model"
)

type (
	ILanguageStorage interface {
		IStorage[*model.Language]
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
