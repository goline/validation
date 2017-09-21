package validation

import (
	"fmt"
	"strconv"

	"github.com/goline/errors"
)

func MaxChecker() Checker {
	return &maxChecker{}
}

type maxChecker struct{}

func (c *maxChecker) Name() string {
	return "max"
}

// Check tests input value with expectation
func (c *maxChecker) Check(v interface{}, expects string) error {
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

func (c *maxChecker) checkInt(v interface{}, expects int64) error {
	i, err := IsInt(v)
	if err != nil {
		return err
	}

	if i > expects {
		return errors.New(ERR_VALIDATOR_NOT_MAX, fmt.Sprintf("%d is greater than %d", i, expects))
	}

	return nil
}

func (c *maxChecker) checkFloat(v interface{}, expects float64) error {
	f, err := IsFloat(v)
	if err != nil {
		return err
	}

	if f > expects {
		return errors.New(ERR_VALIDATOR_NOT_MAX, fmt.Sprintf("%f is greater than %f", f, expects))
	}

	return nil
}
