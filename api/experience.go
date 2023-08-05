package api

type (
	Experience struct {
		Id              int        `json:"id"`
		Start           *YearMonth `json:"start"`
		End             *YearMonth `json:"end"`
		IsCurrent       bool       `json:"is_current"`
		Title           string     `json:"title"`
		EmployeeType    string     `json:"employee_type"`
		Company         string     `json:"company"`
		CompanyLocation string     `json:"company_location"`
		LocationType    string     `json:"location_type"`
		Industry        string     `json:"industry"`
		Description     string     `json:"description"`
	}
)
