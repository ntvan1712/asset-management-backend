package validator_app

import (
	"strings"

	"asset_management_backend/common/error_app"
	"github.com/go-playground/validator/v10"
)

var baseValidator *validator.Validate

func init() {
	baseValidator = validator.New()
}

func ValidateStruct(data interface{}) *error_app.ErrorResponse {
	if err := baseValidator.Struct(data); err != nil {
		var errMsg strings.Builder
		errMsg.WriteString("Invalid input:")
		for _, err := range err.(validator.ValidationErrors) {
			errMsg.WriteString(" " + err.Field() + ": " + err.Tag())
		}
		return error_app.BadRequestErrorResponse(errMsg.String())
	}
	return nil
}

// func IsValidUnixTimeInSecs(unixTimeInSecs int64) bool {
// 	return unixTimeInSecs > minUnixTimeInSecs
// }

// func IsValidUnixTimeInMillis(unixTimeInMillis int64) bool {
// 	return unixTimeInMillis > minUnixTimeInMillis
// }
