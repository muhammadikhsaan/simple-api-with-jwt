package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(errs validator.ValidationErrors) []ErrorModel {

	validationErrors := make([]ErrorModel, 0, len(errs))

	for _, validationErr := range errs {
		validationErrors = append(validationErrors, ErrorModel{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}

	return validationErrors
}
