package validate

import "github.com/go-playground/validator/v10"

var valid = validator.New()

func Var(field interface{}, tag string) error {
	return valid.Var(field, tag)
}

func Struct(s interface{}) error {
	return valid.Struct(s)
}
