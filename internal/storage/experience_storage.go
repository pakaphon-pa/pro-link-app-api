package storage

import (
	"pro-link-api/internal/model"
)

type (
	IExperienceStorage interface {
		IStorage[*model.Experience]
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
