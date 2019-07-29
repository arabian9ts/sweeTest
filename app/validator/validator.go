package validator

type Validation interface {
	Validate(form interface{}) (bool, map[string]string)
}
