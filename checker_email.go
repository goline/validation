package validation

import (
	"fmt"
	"github.com/goline/errors"
	"regexp"
)

func EmailChecker() Checker {
	// A more practical implementation of RFC 5322
	// Omit IP addresses, domain-specific addresses,
	// the syntax using double quotes and square brackets
	pattern := `\A[a-z0-9!#$%&'*+/=?^_` + "`" +
		`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" +
		`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\z`
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return &emailChecker{
		regexp: r,
	}
}

type emailChecker struct {
	regexp *regexp.Regexp
}

func (c *emailChecker) Name() string {
	return "email"
}

func (c *emailChecker) Check(v interface{}, _ string) error {
	s, err := IsString(v)
	if err != nil {
		return err
	}

	if c.regexp.MatchString(s) == false {
		return errors.New(ERR_VALIDATOR_NOT_EMAIL, fmt.Sprintf("%s is not an email address", v))
	}

	return nil
}