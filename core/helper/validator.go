package helper

import (
	"Fibers-Example/core/interfaces"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidatorHelper[T interfaces.RequestBody[T]] struct {
	Model T
}

func (v ValidatorHelper[T]) Validate() []string {
	validate := validator.New()
	err := validate.Struct(v.Model)
	if err != nil {
		return Map(err.(validator.ValidationErrors),
			func(e validator.FieldError) string {
				return fmt.Sprintf("Field %s induces validation error, because of conditions not met: %s", e.Field(), e.Tag())
			},
		)

	}
	return nil
}
