package model

import (
	"time"
)

const (
	ExperienceTableName = "experience"
)

type Experience struct {
	ExpID              int        `gorm:"column:exp_id;primaryKey;autoIncrement;"`
	AccID              int        `gorm:"column:acc_id;"`
	ExpTitle           string     `gorm:"column:exp_title;"`
	ExpEmployeeType    string     `gorm:"column:exp_employee_type;"`
	ExpCompany         string     `gorm:"column:exp_company;"`
	ExpCompanyLocaiton *string    `gorm:"column:exp_company_locaiton;"`
	ExpLocationType    *string    `gorm:"column:exp_location_type;"`
	ExpIndustry        *string    `gorm:"column:exp_industry;"`
	ExpDescription     *string    `gorm:"column:exp_description;"`
	ExpStartYear       int        `gorm:"column:exp_start_year;"`
	ExpStartMonth      *int       `gorm:"column:exp_start_month;"`
	ExpIsCurrent       bool       `gorm:"column:exp_is_current;"`
	ExpEndYear         *int       `gorm:"column:exp_end_year;"`
	ExpEndMonth        *int       `gorm:"column:exp_end_month;"`
	ExpCreatedDate     time.Time  `gorm:"<-:create;column:exp_created_date;default:current_timestamp"`
	ExpCreatedBy       int        `gorm:"<-:create;column:exp_created_by;"`
	ExpUpdatedDate     *time.Time `gorm:"column:exp_updated_date;"`
	ExpUpdatedBy       int        `gorm:"column:exp_updated_by;"`
}

func (Experience) TableName() string {
	return ExperienceTableName
}
