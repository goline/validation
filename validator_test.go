package validation

import (
	"testing"

	"github.com/goline/errors"
)

func TestNewValidator(t *testing.T) {
	v := New()
	if v == nil {
		t.Errorf("Expects v is not nil")
	}
}

func TestFactoryValidator_ParseTags(t *testing.T) {
	tag := "email;min=30;max=100;in=40,50,55;regexp=(\\d+)"
	v := &FactoryValidator{}
	m, err := v.parseTags(tag)
	if err != nil {
		t.Errorf("Expect err is nil. Got %v", err)
	} else if in, ok := m["in"]; ok == false || in != "40,50,55" {
		t.Errorf("Expects in is correct")
	} else if re, ok := m["regexp"]; ok == false || re != `(\d+)` {
		t.Errorf("Expects regexp is correct")
	} else if email, ok := m["email"]; ok == false || email != "" {
		t.Errorf("Expects email is empty")
	} else if min, ok := m["min"]; ok == false || min != "30" {
		t.Errorf("Expects min is correct")
	} else if max, ok := m["max"]; ok == false || max != "100" {
		t.Errorf("Expects max is correct")
	}
}

func TestFactoryValidator_Tag(t *testing.T) {
	v := &FactoryValidator{}
	v.tag = "another_tag"
	if v.Tag() != "another_tag" {
		t.Errorf("Expects %s. Got %s", "another_tag", v.Tag())
	}
}

func TestFactoryValidator_WithTag(t *testing.T) {
	v := &FactoryValidator{tag: "another_tag"}
	v.WithTag("validator")
	if v.tag != "validator" {
		t.Errorf("Expects %s. Got %s", "validator", v.tag)
	}
}

type xChecker struct{}

func (c *xChecker) Name() string                              { return "x" }
func (c *xChecker) Check(v interface{}, expects string) error { return nil }

func TestFactoryValidator_Checker(t *testing.T) {
	v := &FactoryValidator{checkers: make(map[string]Checker)}
	v.WithChecker(&xChecker{})
	_, ok := v.Checker("y")
	if ok == true {
		t.Error("Expects y does not exist")
	}
	_, ok = v.Checker("x")
	if ok == false {
		t.Error("Expects x exists")
	}
}

func TestFactoryValidator_WithChecker(t *testing.T) {
	v := &FactoryValidator{checkers: make(map[string]Checker)}
	if len(v.checkers) != 0 {
		t.Error("Expects v has 0 checker")
	}
	v.WithChecker(&xChecker{})
	if len(v.checkers) != 1 {
		t.Error("Expects v has 1 checker")
	}
}

type emptyInput struct{}
type yChecker struct{}

func (c *yChecker) Name() string                              { return "y" }
func (c *yChecker) Check(v interface{}, expects string) error { return nil }

type zChecker struct{}

func (c *zChecker) Name() string { return "z" }
func (c *zChecker) Check(v interface{}, expects string) error {
	if vv, ok := v.(string); ok == true && vv != "" {
		return nil
	}
	return errors.New("code", "msg")
}

func TestFactoryValidator_Validate_InvalidType(t *testing.T) {
	v := New()
	err := v.Validate("a string")
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_INVALID_TYPE {
		t.Errorf("Expects ERR_VALIDATOR_INVALID_TYPE. Got %v", err)
	}
}

func TestFactoryValidator_Validate_EmptyInputNotError(t *testing.T) {
	i := emptyInput{}
	v := New()
	err := v.Validate(i)
	if err != nil {
		t.Errorf("Expects err is nil. Got %v", err)
	}
}

type sampleInput struct {
	Y string `validate:"y=1;z"`
	X int    `not_validate:"true"`
}

func TestFactoryValidator_Validate_Ok(t *testing.T) {
	v := New()
	v.WithChecker(&yChecker{})
	v.WithChecker(&zChecker{})

	i := sampleInput{}
	err := v.Validate(i)
	if err == nil {
		t.Error("Expects err is not nil")
	}

	ii := &sampleInput{Y: "not_empty"}
	err = v.Validate(ii)
	if err != nil {
		t.Errorf("Expects err is nil. Got %v", err)
	}
}

type sampleInput2 struct {
	Y string `validate:""`
}

func TestFactoryValidator_Validate_EmptyTag(t *testing.T) {
	v := New()
	v.WithChecker(&yChecker{})
	v.WithChecker(&zChecker{})
	err := v.Validate(sampleInput2{"hello"})
	if err != nil {
		t.Errorf("Expects err is nil. Got %v", err)
	}
}

type sampleInput3 struct {
	Y string `validate:"#$%%@"`
}

func TestFactoryValidator_Validate_InvalidTag(t *testing.T) {
	v := New()
	v.WithChecker(&yChecker{})
	v.WithChecker(&zChecker{})
	err := v.Validate(sampleInput3{"hello"})
	if err == nil {
		t.Error("Expects err is not nil")
	} else if e, ok := err.(errors.Error); ok == false || e.Code() != ERR_VALIDATOR_INVALID_TAG {
		t.Errorf("Expects ERR_VALIDATOR_INVALID_TAG. Got %s", e.Code())
	}
}
