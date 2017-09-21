package validation

import (
	"github.com/goline/errors"
	"testing"
)

func getEmailValidator() Validator {
	return New().WithChecker(EmailChecker())
}

type sampleEmailInputTest1 struct {
	Email int64 `validate:"email"`
}

func TestEmailChecker_ERR_VALIDATOR_NOT_STRING(t *testing.T) {
	err := getEmailValidator().Validate(sampleEmailInputTest1{5})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_STRING {
		t.Errorf("Expects ERR_VALIDATOR_NOT_STRING. Got %s", e.Code())
	}
}

type sampleEmailInputTest2 struct {
	Email string `validate:"email"`
}

func TestEmailChecker_ERR_VALIDATOR_NOT_EMAIL(t *testing.T) {
	cases := []string{
		"mail.abc.com",
		"###@ab.@@",
	}

	for _, email := range cases {
		err := getEmailValidator().Validate(sampleEmailInputTest2{email})
		if err == nil {
			t.Errorf("Expects err is not nil for %s", email)
		} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_NOT_EMAIL {
			t.Errorf("Expects ERR_VALIDATOR_NOT_EMAIL. Got %s for %s", e.Code(), email)
		}
	}
}

func TestEmailChecker_Ok(t *testing.T) {
	cases := []string{
		"mail#&*&@abc.com",
		"778@abc.com",
	}

	for _, email := range cases {
		err := getEmailValidator().Validate(sampleEmailInputTest2{email})
		if err != nil {
			t.Errorf("Expects err is nil for %s", email)
		}
	}
}
