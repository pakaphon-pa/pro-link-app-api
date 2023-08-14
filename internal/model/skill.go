package model

import (
	"pro-link-api/api"
	"time"
)

const (
	SkillTableName = "skill"
)

type Skill struct {
	SklID          int        `gorm:"column:skl_id;primaryKey;autoIncrement;"`
	AccID          int        `gorm:"column:acc_id;"`
	SklName        string     `gorm:"column:skl_name;"`
	SklCreatedDate time.Time  `gorm:"<-:create;column:skl_created_date;default:current_timestamp"`
	SklCreatedBy   int        `gorm:"<-:create;column:skl_created_by;"`
	SklUpdatedDate *time.Time `gorm:"column:skl_updated_date;"`
	SklUpdatedBy   int        `gorm:"column:skl_updated_by;"`
}

func (Skill) TableName() string {
	return SkillTableName
}

func ToSkillDomain(data *Skill) *api.Skill {
	return &api.Skill{
		Id:   data.SklID,
		Name: data.SklName,
	}
}

func ToSkillListDoamin(data []*Skill) []*api.Skill {
	result := make([]*api.Skill, 0)

	for _, edu := range data {
		result = append(result, ToSkillDomain(edu))
	}

	return result
}
