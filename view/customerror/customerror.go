package customerror

import (
	"fmt"

	"github.com/freitzzz/iped/model/customerror"
)

// ErrorMessageModelView is a modelview that displays an error message to the consumer
type ErrorMessageModelView struct {
	Message string `json:"message"`
}

// UsingFieldErrorToErrorMessageModelView allows the creation of a custom ErrorMessageModelView using a FieldError
func UsingFieldErrorToErrorMessageModelView(Error customerror.FieldError) ErrorMessageModelView {
	message := fmt.Sprintf(Error.Explanation)

	return ErrorMessageModelView{Message: message}
}
