package model

import (
	"pro-link-api/api"
	"time"
)

const (
	EducationTableName = "education"
)

type Education struct {
	EduID           int        `gorm:"column:edu_id;primaryKey;autoIncrement;"`
	AccID           int        `gorm:"column:acc_id;"`
	EduSchool       string     `gorm:"column:edu_school;"`
	EduDegree       string     `gorm:"column:edu_degree;"`
	EduFieldOfStudy string     `gorm:"column:edu_field_of_study;"`
	EduGrade        *string    `gorm:"column:edu_grade;"`
	EduDescription  *string    `gorm:"column:edu_description;"`
	EduStartYear    int        `gorm:"column:edu_start_year;"`
	EduStartMonth   *int       `gorm:"column:edu_start_month;"`
	EduEndYear      int        `gorm:"column:edu_end_year;"`
	EduEndMonth     *int       `gorm:"column:edu_end_month;"`
	EduCreatedDate  time.Time  `gorm:"<-:create;column:edu_created_date;default:current_timestamp"`
	EduCreatedBy    int        `gorm:"<-:create;column:edu_created_by;"`
	EduUpdatedDate  *time.Time `gorm:"column:edu_updated_date;"`
	EduUpdatedBy    int        `gorm:"column:edu_updated_by;"`
}

func (Education) TableName() string {
	return EducationTableName
}

func ToEducationDomain(data *Education) *api.Education {
	return &api.Education{
		Id:           data.EduID,
		School:       data.EduSchool,
		Degree:       data.EduDegree,
		FieldOfStudy: data.EduFieldOfStudy,
		Start: &api.YearMonth{
			Year:  data.EduStartYear,
			Month: data.EduStartMonth,
		},
		End: &api.YearMonth{
			Year:  data.EduEndYear,
			Month: data.EduEndMonth,
		},
		Grade:       *data.EduGrade,
		Description: *data.EduDescription,
	}
}

func ToEducationListDoamin(data []*Education) []*api.Education {
	result := make([]*api.Education, 0)

	for _, edu := range data {
		result = append(result, ToEducationDomain(edu))
	}

	return result
}
