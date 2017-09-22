package validation

import (
	"fmt"
	"unicode/utf8"

	"github.com/goline/errors"
)

func MaxLengthChecker() Checker {
	return &maxLengthChecker{}
}

type maxLengthChecker struct{}

func (c *maxLengthChecker) Name() string {
	return "maxLength"
}

func (c *maxLengthChecker) Check(v interface{}, expects string) error {
	s, err := IsString(v)
	if err != nil {
		return err
	}

	max, err := IsStringInt(expects)
	if err != nil {
		return err
	}

	l := int64(utf8.RuneCountInString(s))
	if l > max {
		return errors.New(ERR_VALIDATOR_NOT_MAX_LENGTH, fmt.Sprintf("Maximum length is %d. Got %d", max, l))
	}

	return nil
}
