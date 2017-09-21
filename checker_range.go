package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/goline/errors"
)

func RangeChecker() Checker {
	return &rangeChecker{}
}

type rangeChecker struct{}

func (c *rangeChecker) Name() string {
	return "range"
}

func (c *rangeChecker) Check(v interface{}, expects string) error {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int64:
		return c.checkInt(reflect.ValueOf(v).Int(), expects)
	case reflect.Float64:
		return c.checkFloat(reflect.ValueOf(v).Float(), expects)
	default:
		return errors.New(ERR_VALIDATOR_NOT_NUMBER, fmt.Sprintf("%v is not a number", v))
	}
}

func (c *rangeChecker) checkInt(i int64, expects string) error {
	p := `^(\d+)-(\d+)$`
	r, err := regexp.Compile(p)
	if err != nil {
		panic(err)
	}

	if r.MatchString(expects) == false {
		return errors.New(ERR_VALIDATOR_INVALID_FORMAT, fmt.Sprintf("Expects format %s. Got %s", p, expects))
	}

	m := r.FindStringSubmatch(expects)
	min, _ := strconv.ParseInt(m[1], 10, 64)
	max, _ := strconv.ParseInt(m[2], 10, 64)
	if i < min || i > max {
		return errors.New(ERR_VALIDATOR_NOT_IN_RANGE, fmt.Sprintf("%d is not in range (%d, %d)", i, min, max))
	}

	return nil
}

func (c *rangeChecker) checkFloat(f float64, expects string) error {
	p := `^([\d.]+)-([\d.]+)$`
	r, err := regexp.Compile(p)
	if err != nil {
		panic(err)
	}

	if r.MatchString(expects) == false {
		return errors.New(ERR_VALIDATOR_INVALID_FORMAT, fmt.Sprintf("Expects format %s. Got %s", p, expects))
	}

	m := r.FindStringSubmatch(expects)
	min, _ := strconv.ParseFloat(m[1], 64)
	max, _ := strconv.ParseFloat(m[2], 64)
	if f < min || f > max {
		return errors.New(ERR_VALIDATOR_NOT_IN_RANGE, fmt.Sprintf("%f is not in range (%f, %f)", f, min, max))
	}

	return nil
}
