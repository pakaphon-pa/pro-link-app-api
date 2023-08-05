package api

type (
	Language struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Proficiency string `json:"proficiency"`
	}
)
