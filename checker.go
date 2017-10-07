package validation

import "github.com/goline/errors"

type Checker interface {
	// Name returns checker name
	Name() string

	// Check tests input value with expectation
	Check(v interface{}, expects string) errors.Error
}
