package api

type (
	ProfileRequest struct {
		Data *ProfileResponse `json:"username"`
	}
	ProfileResponse struct {
		Id          int           `json:"id"`
		AccId       int           `json:"acc_id"`
		FirstName   string        `json:"first_name" binding:"required"`
		LastName    string        `json:"last_name" binding:"required"`
		About       *string       `json:"about"`
		Website     []*Website    `json:"website_list"`
		PhoneNumber *string       `json:"phone_number"`
		PhoneType   *string       `json:"phone_type"`
		Address     *string       `json:"address"`
		BirthMonth  *string       `json:"birth_month"`
		BirthDay    *string       `json:"birth_day"`
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
