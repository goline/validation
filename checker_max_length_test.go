package validation

import (
	"fmt"
	"github.com/goline/errors"
	"testing"
)

func getMaxLengthValidator() Validator {
	return New().WithChecker(MaxLengthChecker())
}

type sampleMaxLengthInput1 struct {
	Username int64 `validate:"maxLength=3"`
}

func TestMaxLengthChecker_ERR_VALIDATOR_NOT_STRING(t *testing.T) {
	err := getMaxLengthValidator().Validate(sampleMaxLengthInput1{10})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_STRING {
		t.Errorf("Expects ERR_VALIDATOR_NOT_STRING. Got %s", e.Code())
	}
}

type sampleMaxLengthInput2 struct {
	Username string `validate:"maxLength=aa"`
}

func TestMaxLengthChecker_ERR_VALIDATOR_NOT_INT(t *testing.T) {
	err := getMaxLengthValidator().Validate(sampleMaxLengthInput2{"aa"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_INT {
		t.Errorf("Expects ERR_VALIDATOR_NOT_INT. Got %s", e.Code())
	}
}

type sampleMaxLengthInput3 struct {
	Username string `validate:"maxLength=3"`
}

func TestMaxLengthChecker_ERR_VALIDATOR_NOT_MAX_LENGTH(t *testing.T) {
	err := getMaxLengthValidator().Validate(sampleMaxLengthInput3{"aaaa"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_MAX_LENGTH {
		t.Errorf("Expects ERR_VALIDATOR_NOT_MAX_LENGTH. Got %s", e.Code())
		fmt.Println(e.Message())
	}
}

type sampleMaxLengthInput4 struct {
	Username string `validate:"maxLength=3"`
}

func TestMaxLengthChecker_Ok(t *testing.T) {
	err := getMaxLengthValidator().Validate(sampleMaxLengthInput3{"aaa"})
	if err != nil {
		t.Error("Expects err is nil")
	}
}
