package validation

import (
	"github.com/goline/errors"
	"testing"
)

func getMinValidator() Validator {
	return New().WithChecker(MinChecker())
}

type sampleMinInputTest1 struct {
	Age int64 `validate:"min=aa"`
}

func TestMinChecker_ERR_VALIDATOR_NOT_NUMBER(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest1{Age: 10})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_NUMBER {
		t.Errorf("Expects ERR_VALIDATOR_NOT_NUMBER. Got %s", e.Code())
	}
}

type sampleMinInputTest2 struct {
	Age string `validate:"min=10"`
}

func TestMinChecker_ERR_VALIDATOR_NOT_NUMBER_PROP(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest2{Age: "10"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_NUMBER {
		t.Errorf("Expects ERR_VALIDATOR_NOT_NUMBER. Got %s", e.Code())
	}
}

type sampleMinInputTest3 struct {
	Age int64 `validate:"min=10"`
}

func TestMinChecker_ERR_VALIDATOR_NOT_MIN(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest3{Age: 9})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_MIN {
		t.Errorf("Expects ERR_VALIDATOR_NOT_MIN. Got %s", e.Code())
	}
}

type sampleMinInputTest4 struct {
	Age float64 `validate:"min=10.2"`
}

func TestMinChecker_Ok_Float(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest4{Age: 10.3})
	if err != nil {
		t.Error("Expects err is nil")
	}
}

type sampleMinInputTest5 struct {
	Age int64 `validate:"min=10"`
}

func TestMinChecker_Ok_Int(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest5{Age: 11})
	if err != nil {
		t.Error("Expects err is nil")
	}
}

type sampleMinInputTest6 struct {
	Age float64 `validate:"min=10"`
}

func TestMinChecker_ERR_VALIDATOR_NOT_INT(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest6{Age: 11})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_INT {
		t.Errorf("Expects ERR_VALIDATOR_NOT_INT. Got %s", e.Code())
	}
}

type sampleMinInputTest7 struct {
	Age int64 `validate:"min=10.2"`
}

func TestMinChecker_ERR_VALIDATOR_NOT_FLOAT(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest7{Age: 11})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_FLOAT {
		t.Errorf("Expects ERR_VALIDATOR_NOT_FLOAT. Got %s", e.Code())
	}
}

type sampleMinInputTest8 struct {
	Age float64 `validate:"min=10.1"`
}

func TestMaxChecker_ERR_VALIDATOR_NOT_MIN_Float(t *testing.T) {
	v := getMinValidator()
	err := v.Validate(&sampleMinInputTest8{Age: 10.0})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_MIN {
		t.Errorf("Expects ERR_VALIDATOR_NOT_MIN. Got %s", e.Code())
	}
}
