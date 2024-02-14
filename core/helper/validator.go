package helper

import (
	"Fibers-Example/core/interfaces"
	"github.com/go-playground/validator/v10"
)

type ValidatorHelper[T interfaces.RequestBody[T]] struct {
	Model T
}

func (v ValidatorHelper[T]) Validate() error {
	validate := validator.New()
	err := validate.Struct(v.Model)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
