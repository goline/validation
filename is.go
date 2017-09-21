package validation

import (
	"fmt"
	"github.com/goline/errors"
	"reflect"
)

func IsString(v interface{}) (string, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.String {
		return "", errors.New(ERR_VALIDATOR_NOT_STRING, fmt.Sprintf("%v is not a string", v))
	}

	return reflect.ValueOf(v).String(), nil
}

func IsInt(v interface{}) (int64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Int64 {
		return 0, errors.New(ERR_VALIDATOR_NOT_INT, fmt.Sprintf("%v is not int", v))
	}

	return reflect.ValueOf(v).Int(), nil
}

func IsFloat(v interface{}) (float64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Float64 {
		return 0.0, errors.New(ERR_VALIDATOR_NOT_FLOAT, fmt.Sprintf("%v is not float", v))
	}

	return reflect.ValueOf(v).Float(), nil
}

func IsNumber(v interface{}) error {
	_, isF := IsFloat(v)
	_, isI := IsInt(v)
	if isF != nil && isI != nil {
		return errors.New(ERR_VALIDATOR_NOT_NUMBER, fmt.Sprintf("%v is not a number", v))
	}

	return nil
}

func IsNil(v interface{}) error {
	if v != nil {
		return errors.New(ERR_VALIDATOR_NOT_NIL, fmt.Sprintf("%v is not nil", v))
	}

	return nil
}
