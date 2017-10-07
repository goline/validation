package validation

import (
	"fmt"
	"regexp"

	"github.com/goline/errors"
)

func RegExpChecker() Checker {
	return &regExpChecker{}
}

type regExpChecker struct{}

func (c *regExpChecker) Name() string {
	return "regexp"
}

func (c *regExpChecker) Check(v interface{}, expects string) errors.Error {
	r, e := regexp.Compile(expects)
	if e != nil {
		return errors.New(ERR_VALIDATOR_REGEXP_WRONG_PATTERN, fmt.Sprintf("RegExp failed to compile. Got %s", e.Error()))
	}

	s, err := IsString(v)
	if err != nil {
		return err
	}

	if r.MatchString(s) == false {
		return errors.New(ERR_VALIDATOR_REGEXP_NOT_MATCH, fmt.Sprintf("%s does not match %s", s, expects))
	}

	return nil
}
