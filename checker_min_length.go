package validation

import (
	"fmt"
	"unicode/utf8"

	"github.com/goline/errors"
)

func MinLengthChecker() Checker {
	return &minLengthChecker{}
}

type minLengthChecker struct{}

func (c *minLengthChecker) Name() string {
	return "minLength"
}

func (c *minLengthChecker) Check(v interface{}, expects string) error {
	s, err := IsString(v)
	if err != nil {
		return err
	}

	min, err := IsStringInt(expects)
	if err != nil {
		return err
	}

	l := int64(utf8.RuneCountInString(s))
	if l < min {
		return errors.New(ERR_VALIDATOR_NOT_MIN_LENGTH, fmt.Sprintf("Minimum length is %d. Got %d", min, l))
	}

	return nil
}
