package validation

import (
	"fmt"
	"strconv"

	"github.com/goline/errors"
)

func MinChecker() Checker {
	return &minChecker{}
}

type minChecker struct{}

func (c *minChecker) Name() string {
	return "min"
}

// Check tests input value with expectation
func (c *minChecker) Check(v interface{}, expects string) error {
	if err := IsNumber(v); err != nil {
		return err
	}

	ei, err := strconv.ParseInt(expects, 10, 64)
	if err == nil {
		return c.checkInt(v, ei)
	}

	ef, err := strconv.ParseFloat(expects, 64)
	if err == nil {
		return c.checkFloat(v, ef)
	}

	return errors.New(ERR_VALIDATOR_NOT_NUMBER, fmt.Sprintf("%v is not a number", v))
}

func (c *minChecker) checkInt(v interface{}, expects int64) error {
	i, err := IsInt(v)
	if err != nil {
		return err
	}

	if i < expects {
		return errors.New(ERR_VALIDATOR_NOT_MIN, fmt.Sprintf("%d is lower than %d", i, expects))
	}

	return nil
}

func (c *minChecker) checkFloat(v interface{}, expects float64) error {
	f, err := IsFloat(v)
	if err != nil {
		return err
	}

	if f < expects {
		return errors.New(ERR_VALIDATOR_NOT_MIN, fmt.Sprintf("%f is lower than %f", f, expects))
	}

	return nil
}
