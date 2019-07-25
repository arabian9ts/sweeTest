package validator

import v9 "gopkg.in/go-playground/validator.v9"

type DefaultValidator struct {
	Validator *v9.Validate
}

func NewDefaultValidator() Validation {
	return &DefaultValidator{
		Validator: v9.New(),
	}
}

func (val *DefaultValidator) Validate(form interface{}) error {
	return val.Validator.Struct(form)
}
