package exceptions

import (
	"fmt"
	"net/http"
	"pro-link-api/api"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HttpError struct {
	Status  int
	Code    string
	Message string
}

func HttpErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {

			lastError := c.Errors.Last()
			msg := "An internal server error occurred"
			statusCode := http.StatusInternalServerError
			validateError := make([]*api.FieldError, 0)
			if lastError != nil {
				if err, ok := lastError.Err.(validator.ValidationErrors); ok {
					fmt.Println("Is validateor Error :", ok)
					msg = "Validation errors"
					statusCode = http.StatusBadRequest
					validateError = getValidatorResponseError(err)
				} else if err, ok := lastError.Err.(*HttpError); ok {
					fmt.Println("Is http error :", ok)
					msg = err.Message
					statusCode = err.Status
				} else {
					fmt.Println("Common Error :", lastError.Err.Error())
					msg = lastError.Err.Error()
				}
			}

			c.JSON(statusCode, &api.ErrorResponse{
				Code:     statusCode,
				Message:  msg,
				Validate: validateError,
			})
			c.Abort()
		}
	}
}

func NewWithStatus(status int, code, message string) *HttpError {
	return &HttpError{Status: status, Code: code, Message: message}
}

func (he *HttpError) Error() string {
	return he.Code
}

func getValidatorResponseError(errorList validator.ValidationErrors) []*api.FieldError {
	result := make([]*api.FieldError, 0)

	for _, err := range errorList {
		result = append(result, &api.FieldError{

			FieldName: strings.ToLower(err.Field()),
			Tag:       err.Tag(),
		})
	}

	return result
}
