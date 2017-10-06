package validation

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/goline/errors"
)

type Validator interface {
	// Validate checks input value for error
	Validate(v interface{}) error

	ValidatorTagger
	ValidatorChecker
}

type ValidatorTagger interface {
	// Tag returns tag's string
	Tag() string

	// WithTag sets tag
	WithTag(tag string) Validator
}

type ValidatorChecker interface {
	// Checker returns a specific check by name
	Checker(name string) (Checker, bool)

	// WithChecker registers a checker
	WithChecker(checker Checker) Validator
}

type ValidatorErrorModifier interface {
	// WithErrorModifier sets default error modifier
	WithErrorModifier(modifier ErrorModifier) Validator
}

type ErrorModifier func(err errors.Error)

func New() Validator {
	v := &FactoryValidator{
		tag:      "validate",
		checkers: make(map[string]Checker),
	}
	v.WithChecker(MinChecker()).
		WithChecker(MaxChecker()).
		WithChecker(MinLengthChecker()).
		WithChecker(MaxLengthChecker()).
		WithChecker(RangeChecker()).
		WithChecker(EmailChecker()).
		WithChecker(RegExpChecker()).
		WithChecker(InChecker())

	return v
}

type FactoryValidator struct {
	tag           string
	checkers      map[string]Checker
	errorModifier ErrorModifier
}

func (v *FactoryValidator) Tag() string {
	return v.tag
}

func (v *FactoryValidator) WithTag(tag string) Validator {
	v.tag = tag
	return v
}

func (v *FactoryValidator) Checker(name string) (Checker, bool) {
	c, ok := v.checkers[name]
	return c, ok
}

func (v *FactoryValidator) WithChecker(checker Checker) Validator {
	v.checkers[checker.Name()] = checker
	return v
}

func (v *FactoryValidator) WithErrorModifier(modifier ErrorModifier) Validator {
	v.errorModifier = modifier
	return v
}

func (v *FactoryValidator) Validate(input interface{}) error {
	t, err := v.validateType(input)
	if err != nil {
		return v.modifyError("VALIDATOR", err)
	}

	n := t.NumField()
	if n == 0 {
		// No fields => no process
		return nil
	}

	val, _ := v.valueOf(input)
	for i := 0; i < n; i++ {
		vf := val.Field(i)
		if vf.CanInterface() && (vf.Kind() == reflect.Struct || vf.Kind() == reflect.Ptr) {
			if err := v.Validate(vf.Interface()); err != nil {
				return err
			}
		}

		sf := t.Field(i)
		if _, ok := sf.Tag.Lookup(v.tag); ok == false {
			continue
		}

		tag := sf.Tag.Get(v.tag)
		if tag == "" {
			continue
		}

		m, err := v.parseTags(tag)
		if err != nil {
			return v.modifyError("VALIDATOR", err)
		}

		for k, p := range m {
			c, ok := v.Checker(k)
			if ok == false {
				continue
			}
			if err := c.Check(vf.Interface(), p); err != nil {
				return v.modifyError(sf.Name, err)
			}
		}
	}

	return nil
}

func (v *FactoryValidator) validateType(input interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(input)
	switch t.Kind() {
	case reflect.Ptr:
		return t.Elem(), nil
	case reflect.Struct:
		return t, nil
	default:
		return nil, errors.New(ERR_VALIDATOR_INVALID_TYPE, fmt.Sprintf("%s type is not supported", t.Kind().String()))
	}
}

func (v *FactoryValidator) valueOf(input interface{}) (reflect.Value, error) {
	t := reflect.TypeOf(input)
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.ValueOf(input).Elem(), nil
	case reflect.Struct:
		return reflect.ValueOf(input), nil
	default:
		return reflect.Value{}, errors.New(ERR_VALIDATOR_INVALID_TYPE, fmt.Sprintf("%s type is not supported", t.Kind().String()))
	}
}

func (v *FactoryValidator) modifyError(key string, err error) error {
	var r errors.Error
	if e, ok := err.(errors.Error); ok == true {
		e.WithMessage(fmt.Sprintf("%s: %s", key, e.Message()))
		r = e
	} else {
		r = errors.New(ERR_VALIDATOR_UNKNOWN_ERROR, fmt.Sprintf("%s: %s", key, err.Error()))
	}
	if v.errorModifier != nil {
		v.errorModifier(r)
	}

	return r
}

func (v *FactoryValidator) parseTags(tag string) (map[string]string, error) {
	m := make(map[string]string)
	p := `([^\W]+)(=?([^=;]+)?)`
	r, _ := regexp.Compile(p)

	if !r.MatchString(tag) {
		return m, errors.New(ERR_VALIDATOR_INVALID_TAG, fmt.Sprintf("%s is not a valid tag string", tag))
	}

	mm := r.FindAllStringSubmatch(tag, -1)
	for _, sm := range mm {
		m[sm[1]] = sm[3]
	}

	return m, nil
}
