package model

import (
	"pro-link-api/api"
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

func ToExperienceDomain(data *Experience) *api.Experience {
	return &api.Experience{
		Id:              data.ExpID,
		Title:           data.ExpTitle,
		EmployeeType:    data.ExpEmployeeType,
		LocationType:    *data.ExpCompanyLocaiton,
		Company:         data.ExpCompany,
		CompanyLocation: *data.ExpCompanyLocaiton,
		Industry:        *data.ExpIndustry,
		IsCurrent:       data.ExpIsCurrent,
		Start: &api.YearMonth{
			Year:  data.ExpStartYear,
			Month: data.ExpStartMonth,
		},
		End: &api.YearMonth{
			Year:  *data.ExpEndYear,
			Month: data.ExpEndMonth,
		},
		Description: *data.ExpDescription,
	}
}

func ToExperienceListDoamin(data []*Experience) []*api.Experience {
	result := make([]*api.Experience, 0)

	for _, edu := range data {
		result = append(result, ToExperienceDomain(edu))
	}

	return result
}
