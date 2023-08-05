package storage

import (
	"pro-link-api/internal/model"
)

type (
	ISkillStorage interface {
		IStorage[*model.Skill]
	}

	SkillStorage struct {
		AbstractStorage[*model.Skill]
	}
)

func NewSkillStorage(s *Storage) *SkillStorage {
	return &SkillStorage{
		AbstractStorage[*model.Skill]{
			db:        s.db,
			tableName: model.SkillTableName,
		},
	}
}
