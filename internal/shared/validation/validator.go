package validation

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)


var validate *validator.Validate

func ValidateCheckFields(ctx context.Context, data interface{}) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	if err:= validate.StructCtx(ctx, data);err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return err
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			var errorMessages []string
			for _, e := range validateErrs {
				errorMessages = append(errorMessages, fmt.Sprintf("field %s: %s %s", strings.ToLower(e.Field()), e.ActualTag(), e.Param()))
			}

			return errors.New(strings.Join(errorMessages, ", "))
		}

	}

	

	return nil
}