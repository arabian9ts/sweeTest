package validator

import (
	"fmt"
	v9 "gopkg.in/go-playground/validator.v9"
)

type DefaultValidator struct {
	Validator *v9.Validate
}

func NewDefaultValidator() Validation {
	return &DefaultValidator{
		Validator: v9.New(),
	}
}

func (val *DefaultValidator) Validate(form interface{}) (ok bool, tags map[string]string) {
	err := val.Validator.Struct(form)
	if err == nil {
		ok = true
		return
	}

	ok = false
	tags = make(map[string]string)
	errors := err.(v9.ValidationErrors)
	if len(errors) == 0 {
		return
	}

	for i := range errors {
		field := errors[i].Field()
		switch errors[i].Tag() {
		case "required":
			tags[field] = fmt.Sprintf("%s is required", field)
		case "lte":
			param := errors[i].Param()
			tags[field] = fmt.Sprintf("%s is should be equal or smaller than %s", field, param)
		case "lt":
			param := errors[i].Param()
			tags[field] = fmt.Sprintf("%s is should be smaller than %s", field, param)
		case "gte":
			param := errors[i].Param()
			tags[field] = fmt.Sprintf("%s is should be equal or larger than %s", field, param)
		case "gt":
			param := errors[i].Param()
			tags[field] = fmt.Sprintf("%s is should be larger than %s", field, param)
		case "min":
			param := errors[i].Param()
			tags[field] = fmt.Sprintf("%s is should be equal or larger than %s", field, param)
		case "max":
			param := errors[i].Param()
			tags[field] = fmt.Sprintf("%s is should be equal or smaller than %s", field, param)
		default:
			tags[field] = fmt.Sprintf("%s is invalid", field)
		}
	}
	return
}
