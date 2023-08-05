package model

import (
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
