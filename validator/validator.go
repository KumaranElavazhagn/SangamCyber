package validator

import (
	"sangamCyber/dto"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateUserInfoRequest(req *dto.UserInfoRequest) []string {
	validate := validator.New()

	var errors []string

	if err := validate.Struct(req); err != nil {
		errorMappings := map[string]string{
			"'UserName' failed on the 'required' tag":    "Missing Parameter, userName",
			"'Password' failed on the 'required' tag":    "Missing Parameter, password",
			"'FirstName' failed on the 'required' tag":   "Missing Parameter, firstName",
			"'LastName' failed on the 'required' tag":    "Missing Parameter, lastName",
			"'EmailID' failed on the 'required' tag":     "Missing Parameter, emailID",
			"'DateOfBirth' failed on the 'required' tag": "Missing Parameter, dateOfBirth",
		}

		for key, value := range errorMappings {
			if strings.Contains(err.Error(), key) {
				errors = append(errors, value)
			}
		}
	}

	return errors
}
