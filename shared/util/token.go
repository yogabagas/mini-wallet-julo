package util

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

func TokenValidation(token string) bool {
	err := validation.Validate(token,
		validation.Required,
		validation.Match(regexp.MustCompile(`^(s|bearer|Bearer).([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`)))
	return err != nil
}
