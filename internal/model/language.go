package model

import (
	"pro-link-api/api"
	"time"
)

const (
	LanguageTableName = "language"
)

type Language struct {
	LanID          int        `gorm:"column:lan_id;primaryKey;autoIncrement;"`
	AccID          int        `gorm:"column:acc_id;"`
	LanName        string     `gorm:"column:lan_name;"`
	LanProficiency string     `gorm:"column:lan_proficiency;"`
	LanCreatedDate time.Time  `gorm:"<-:create;column:lan_created_date;default:current_timestamp"`
	LanCreatedBy   int        `gorm:"<-:create;column:lan_created_by;"`
	LanUpdatedDate *time.Time `gorm:"column:lan_updated_date;"`
	LanUpdatedBy   int        `gorm:"column:lan_updated_by;"`
}

func (Language) TableName() string {
	return LanguageTableName
}

func ToLanguageDomain(data *Language) *api.Language {
	return &api.Language{
		Id:          data.LanID,
		Name:        data.LanName,
		Proficiency: data.LanProficiency,
	}
}

func ToLanguageListDoamin(data []*Language) []*api.Language {
	result := make([]*api.Language, 0)

	for _, edu := range data {
		result = append(result, ToLanguageDomain(edu))
	}

	return result
}
