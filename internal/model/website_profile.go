package model

import (
	"time"
)

const (
	WebsiteProfileTableName = "website_profile"
)

type WebsiteProfile struct {
	WebID          int        `gorm:"column:web_id;primaryKey;autoIncrement;"`
	PrfID          int        `gorm:"column:prf_id;"`
	WebName        string     `gorm:"column:web_name;"`
	WebType        string     `gorm:"column:web_type;"`
	WebCreatedDate time.Time  `gorm:"<-:create;column:web_created_date;default:current_timestamp"`
	WebCreatedBy   int        `gorm:"<-:create;column:web_created_by;"`
	WebUpdatedDate *time.Time `gorm:"column:web_updated_date;"`
	WebUpdatedBy   int        `gorm:"column:web_updated_by;"`
}

func (WebsiteProfile) TableName() string {
	return WebsiteProfileTableName
}
