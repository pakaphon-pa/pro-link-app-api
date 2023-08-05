package storage

import "pro-link-api/internal/model"

type (
	IEducationStorage interface {
		IStorage[*model.Education]
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
