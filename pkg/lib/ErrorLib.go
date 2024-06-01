package lib

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be minimum " + fe.Param()
	case "min":
		return "Should be greater " + fe.Param()
	case "max":
		return "Should not be greater than " + fe.Param()
	case "oneof":
		return "Invalid planet type: " + fe.Value().(string)
	default:
		return "Validation error:" + fe.Error()
	}

}
func ValidationError(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	var OutPut []ErrorMsg
	if errors.As(err, &ve) {

		for _, fe := range ve {
			custome_message := ErrorMsg{fe.Field(), getErrorMsg(fe)}
			OutPut = append(OutPut, custome_message)
		}
	} else {
		custome_message := ErrorMsg{"Validation Error", err.Error()}
		OutPut = append(OutPut, custome_message)
	}
	return OutPut
}
