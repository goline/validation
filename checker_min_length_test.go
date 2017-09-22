package validation

import (
	"fmt"
	"github.com/goline/errors"
	"testing"
)

func getMinLengthValidator() Validator {
	return New().WithChecker(MinLengthChecker())
}

type sampleMinLengthInput1 struct {
	Username int64 `validate:"minLength=3"`
}

func TestMinLengthChecker_ERR_VALIDATOR_NOT_STRING(t *testing.T) {
	err := getMinLengthValidator().Validate(sampleMinLengthInput1{10})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_STRING {
		t.Errorf("Expects ERR_VALIDATOR_NOT_STRING. Got %s", e.Code())
	}
}

type sampleMinLengthInput2 struct {
	Username string `validate:"minLength=aa"`
}

func TestMinLengthChecker_ERR_VALIDATOR_NOT_INT(t *testing.T) {
	err := getMinLengthValidator().Validate(sampleMinLengthInput2{"aa"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_INT {
		t.Errorf("Expects ERR_VALIDATOR_NOT_INT. Got %s", e.Code())
	}
}

type sampleMinLengthInput3 struct {
	Username string `validate:"minLength=3"`
}

func TestMinLengthChecker_ERR_VALIDATOR_NOT_MIN_LENGTH(t *testing.T) {
	err := getMinLengthValidator().Validate(sampleMinLengthInput3{"aa"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_MIN_LENGTH {
		t.Errorf("Expects ERR_VALIDATOR_NOT_MIN_LENGTH. Got %s", e.Code())
		fmt.Println(e.Message())
	}
}

type sampleMinLengthInput4 struct {
	Username string `validate:"minLength=3"`
}

func TestMinLengthChecker_Ok(t *testing.T) {
	err := getMinLengthValidator().Validate(sampleMinLengthInput3{"aaa"})
	if err != nil {
		t.Error("Expects err is nil")
	}
}
