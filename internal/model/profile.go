package model

import (
	"time"
)

const (
	ProfileTableName = "profile"
)

type Profile struct {
	PrfID          int        `gorm:"column:prf_id;primaryKey;autoIncrement;"`
	AccID          int        `gorm:"column:acc_id;"`
	PrfFirstName   string     `gorm:"column:prf_first_name;"`
	PrfLastName    string     `gorm:"column:prf_last_name;"`
	PrfAbout       *string    `gorm:"column:prf_about;"`
	PrfPhoneNumber *string    `gorm:"column:prf_phone_number;"`
	PrfPhoneType   *string    `gorm:"column:prf_phone_type;"`
	PrfAddress     *string    `gorm:"column:prf_address;"`
	PrfBirthMonth  *string    `gorm:"column:prf_birth_month;"`
	PrfBirthDate   *string    `gorm:"column:prf_birth_date;"`
	PrfCreatedDate time.Time  `gorm:"<-:create;column:prf_created_date;default:current_timestamp"`
	PrfCreatedBy   int        `gorm:"<-:create;column:prf_created_by;"`
	PrfUpdatedDate *time.Time `gorm:"column:prf_updated_date;"`
	PrfUpdatedBy   int        `gorm:"column:prf_updated_by;"`
}

func (Profile) TableName() string {
	return ProfileTableName
}
