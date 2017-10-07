package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/goline/errors"
)

func IsString(v interface{}) (string, errors.Error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.String {
		return "", errors.New(ERR_VALIDATOR_NOT_STRING, fmt.Sprintf("%v is not a string", v))
	}

	return reflect.ValueOf(v).String(), nil
}

func IsInt(v interface{}) (int64, errors.Error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Int64 {
		return 0, errors.New(ERR_VALIDATOR_NOT_INT, fmt.Sprintf("%v is not an integer", v))
	}

	return reflect.ValueOf(v).Int(), nil
}

func IsStringInt(s string) (int64, errors.Error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New(ERR_VALIDATOR_NOT_INT, fmt.Sprintf("%s is not an integer", s))
	}

	return i, nil
}

func IsFloat(v interface{}) (float64, errors.Error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Float64 {
		return 0.0, errors.New(ERR_VALIDATOR_NOT_FLOAT, fmt.Sprintf("%v is not a float number", v))
	}

	return reflect.ValueOf(v).Float(), nil
}

func IsStringFloat(s string) (float64, errors.Error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, errors.New(ERR_VALIDATOR_NOT_FLOAT, fmt.Sprintf("%s is not a float number", s))
	}

	return f, nil
}

func IsNumber(v interface{}) errors.Error {
	_, isF := IsFloat(v)
	_, isI := IsInt(v)
	if isF != nil && isI != nil {
		return errors.New(ERR_VALIDATOR_NOT_NUMBER, fmt.Sprintf("%v is not a number", v))
	}

	return nil
}

func IsNil(v interface{}) errors.Error {
	if v != nil {
		return errors.New(ERR_VALIDATOR_NOT_NIL, fmt.Sprintf("%v is not nil", v))
	}

	return nil
}
