package validation

import (
	"github.com/goline/errors"
	"testing"
)

func getMaxValidator() Validator {
	return New().WithChecker(MaxChecker())
}

type sampleMaxInputTest1 struct {
	Age int64 `validate:"max=aa"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_NUMBER(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest1{Age: 10})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_NUMBER {
		t.Errorf("Expects ERR_VALIDATOR_NOT_NUMBER. Got %s", e.Code())
	}
}

type sampleMaxInputTest2 struct {
	Age string `validate:"max=10"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_NUMBER_PROP(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest2{Age: "10"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_NUMBER {
		t.Errorf("Expects ERR_VALIDATOR_NOT_NUMBER. Got %s", e.Code())
	}
}

type sampleMaxInputTest3 struct {
	Age int64 `validate:"max=10"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_MAX(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest3{Age: 11})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_MAX {
		t.Errorf("Expects ERR_VALIDATOR_NOT_MAX. Got %s", e.Code())
	}
}

type sampleMaxInputTest4 struct {
	Age float64 `validate:"max=10.2"`
}

func TestMaxChecker_Ok_Float(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest4{Age: 10.1})
	if err != nil {
		t.Error("Expects err is nil")
	}
}

type sampleMaxInputTest5 struct {
	Age int64 `validate:"max=10"`
}

func TestMaxChecker_Ok_Int(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest5{Age: 9})
	if err != nil {
		t.Error("Expects err is nil")
	}
}

type sampleMaxInputTest6 struct {
	Age float64 `validate:"max=10"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_INT(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest6{Age: 11})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_INT {
		t.Errorf("Expects ERR_VALIDATOR_NOT_INT. Got %s", e.Code())
	}
}

type sampleMaxInputTest7 struct {
	Age int64 `validate:"max=10.2"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_FLOAT(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest7{Age: 11})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_FLOAT {
		t.Errorf("Expects ERR_VALIDATOR_NOT_FLOAT. Got %s", e.Code())
	}
}

type sampleMaxInputTest8 struct {
	Age float64 `validate:"max=10.1"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_MAX_Float(t *testing.T) {
	v := getMaxValidator()
	err := v.Validate(&sampleMaxInputTest8{Age: 10.2})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_MAX {
		t.Errorf("Expects ERR_VALIDATOR_NOT_MAX. Got %s", e.Code())
	}
}
