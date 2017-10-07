package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/goline/errors"
)

func InChecker() Checker {
	return new(inChecker)
}

type inChecker struct{}

func (c *inChecker) Name() string {
	return "in"
}

func (c *inChecker) Check(v interface{}, expects string) errors.Error {
	if expects == "" {
		return errors.New(ERR_VALIDATOR_IN_EMPTY_LIST, fmt.Sprintf("%s is not expected to appear in an empty list", v))
	}
	in := strings.Split(expects, ",")

	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int64:
		return c.checkInt(reflect.ValueOf(v).Int(), in)
	case reflect.Float64:
		return c.checkFloat(reflect.ValueOf(v).Float(), in)
	case reflect.String:
		return c.checkString(reflect.ValueOf(v).String(), in)
	default:
		return errors.New(ERR_VALIDATOR_INVALID_ARGUMENT, fmt.Sprintf("%s is not supported by inChecker", t.Kind().String()))
	}
}

func (c *inChecker) checkInt(v int64, in []string) errors.Error {
	for _, s := range in {
		i, err := IsStringInt(s)
		if err != nil {
			return err
		}
		if v == i {
			return nil
		}
	}

	return errors.New(ERR_VALIDATOR_NOT_IN_LIST, fmt.Sprintf("%d is not appeared in %v", v, in))
}

func (c *inChecker) checkFloat(v float64, in []string) errors.Error {
	for _, s := range in {
		i, err := IsStringFloat(s)
		if err != nil {
			return err
		}
		if v == i {
			return nil
		}
	}

	return errors.New(ERR_VALIDATOR_NOT_IN_LIST, fmt.Sprintf("%f is not appeared in %v", v, in))
}

func (c *inChecker) checkString(v string, in []string) errors.Error {
	for _, s := range in {
		if strings.Compare(v, s) == 0 {
			return nil
		}
	}

	return errors.New(ERR_VALIDATOR_NOT_IN_LIST, fmt.Sprintf("%s is not appeared in %v", v, in))
}
