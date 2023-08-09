package model

import (
	"time"
)

const (
	AccountTableName = "account"
)

type Account struct {
	AccID               int        `gorm:"column:acc_id;primaryKey;autoIncrement;"`
	AccEmail            string     `gorm:"column:acc_email;"`
	AccUsername         string     `gorm:"column:acc_username;"`
	AccPassword         string     `gorm:"column:acc_password;"`
	AccLastLogin        *time.Time `gorm:"column:acc_last_login;"`
	AccVerificationCode *string    `gorm:"column:acc_verification_code;"`
	AccIsVerified       bool       `gorm:"column:acc_is_verified;"`
	AccCreatedDate      time.Time  `gorm:"<-:create;column:acc_created_date;default:current_timestamp"`
	AccCreatedBy        int        `gorm:"<-:create;column:acc_created_by;"`
	AccUpdatedDate      *time.Time `gorm:"column:acc_updated_date;"`
	AccUpdatedBy        int        `gorm:"column:acc_updated_by;"`
}

func (Account) TableName() string {
	return AccountTableName
}
