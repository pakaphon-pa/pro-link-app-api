package api

type (
	ErrorResponse struct {
		Code     int           `json:"code"`
		Message  string        `json:"message"`
		Validate []*FieldError `json:"validate,omitempty"`
	}

	FieldError struct {
		FieldName string `json:"field_name"`
		Tag       string `json:"tag"`
	}
)
