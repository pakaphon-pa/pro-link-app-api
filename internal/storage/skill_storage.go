package storage

import (
	"context"
	"pro-link-api/internal/model"
)

type (
	ISkillStorage interface {
		IStorage[*model.Skill]
		FindByAccId(c context.Context, accId int) (data []*model.Skill, err error)
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

func (s *SkillStorage) FindByAccId(c context.Context, accId int) (data []*model.Skill, err error) {
	err = s.db.WithContext(c).Table(s.tableName).Where("acc_id = ?", accId).Find(&data).Error
	return data, err
}
