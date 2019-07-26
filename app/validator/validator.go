package validator

type Validation interface {
	Validate(form interface{}) error
}
