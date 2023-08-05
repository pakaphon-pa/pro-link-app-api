package api

type (
	Profile struct {
		Id          int           `json:"id"`
		FirstName   string        `json:"first_name"`
		LastName    string        `json:"last_name"`
		About       *string       `json:"about"`
		Website     []*Website    `json:"website_list"`
		PhoneNumber *string       `json:"phone_number"`
		PhoneType   *string       `json:"phone_type"`
		Address     *string       `json:"address"`
		BirthMonth  *int          `json:"birth_month"`
		BirthDay    *int          `json:"birth_day"`
		Skill       []*Skill      `json:"skill"`
		Experience  []*Experience `json:"experience"`
		Education   []*Education  `json:"education"`
		Language    []*Language   `json:"language"`
	}

	Website struct {
		Id          int    `json:"id"`
		Website     string `json:"website"`
		WebsiteType string `json:"website_type"`
	}
)
