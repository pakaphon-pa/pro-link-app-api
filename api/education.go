package api

type (
	Education struct {
		Id           int        `json:"id"`
		School       string     `json:"school"`
		Degree       string     `json:"degree"`
		FieldOfStudy string     `json:"field_of_study"`
		Start        *YearMonth `json:"start"`
		End          *YearMonth `json:"end"`
		Grade        string     `json:"grade"`
		Description  string     `json:"description"`
	}

	YearMonth struct {
		Year  int  `json:"year"`
		Month *int `json:"month"`
	}
)
