package validation

import (
	"github.com/goline/errors"
	"testing"
)

func getRangeValidator() Validator {
	return New().WithChecker(RangeChecker())
}

type sampleRangeInputTest1 struct {
	Age string `validate:"range=18-60"`
}

func TestRangeChecker_ERR_VALIDATOR_NOT_NUMBER(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest1{"10"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_NUMBER {
		t.Errorf("Expects ERR_VALIDATOR_NOT_NUMBER. Got %s", e.Code())
	}
}

type sampleRangeInputTest2 struct {
	Age int64 `validate:"range=18-60a"`
}

func TestRangeChecker_ERR_VALIDATOR_INVALID_FORMAT(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest2{10})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_INVALID_FORMAT {
		t.Errorf("Expects ERR_VALIDATOR_INVALID_FORMAT. Got %s", e.Code())
	}
}

type sampleRangeInputTest3 struct {
	Age int64 `validate:"range=18-60"`
}

func TestRangeChecker_ERR_VALIDATOR_NOT_IN_RANGE(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest3{10})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_IN_RANGE {
		t.Errorf("Expects ERR_VALIDATOR_NOT_IN_RANGE. Got %s", e.Code())
	}
}

type sampleRangeInputTest4 struct {
	Age float64 `validate:"range=18-60a"`
}

func TestRangeChecker_ERR_VALIDATOR_INVALID_FORMAT_float(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest4{10.2})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_INVALID_FORMAT {
		t.Errorf("Expects ERR_VALIDATOR_INVALID_FORMAT. Got %s", e.Code())
	}
}

type sampleRangeInputTest5 struct {
	Age float64 `validate:"range=10.0-10.2"`
}

func TestRangeChecker_ERR_VALIDATOR_NOT_IN_RANGE_float(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest5{9.9})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_IN_RANGE {
		t.Errorf("Expects ERR_VALIDATOR_NOT_IN_RANGE. Got %s", e.Code())
	}
}

type sampleRangeInputTest6 struct {
	Age int64 `validate:"range=9-11"`
}

func TestRangeChecker_Ok_Int(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest6{10})
	if err != nil {
		t.Error("Expects err is nil")
	}
}

type sampleRangeInputTest7 struct {
	Age float64 `validate:"range=10.0-10.2"`
}

func TestRangeChecker_Ok_Float(t *testing.T) {
	err := getRangeValidator().Validate(sampleRangeInputTest7{10.1})
	if err != nil {
		t.Error("Expects err is nil")
	}
}
