package validation

import (
	ee "errors"
	"testing"

	"fmt"
	"github.com/goline/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator Suite")
}

type xChecker struct{}

func (c *xChecker) Name() string                              { return "x" }
func (c *xChecker) Check(v interface{}, expects string) error { return nil }

type emptyValidatorInput struct{}
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

type sampleValidatorInput struct {
	Y string `validate:"y=1;z"`
	X int    `not_validate:"true"`
}

type sampleValidatorInput2 struct {
	Y string `validate:""`
}

type sampleValidatorInput3 struct {
	Y string `validate:"#$%%@"`
}

type sampleValidatorInput4 struct {
	Y string `validate:"mailer"`
}

var _ = Describe("Validator", func() {
	It("should return an instance of Validator", func() {
		Expect(New()).NotTo(BeNil())
	})
})
var _ = Describe("FactoryValidator", func() {
	It("[private] should parse tags", func() {
		tag := "email;min=30;max=100;in=40,50,55;regexp=(\\d+)"
		v := &FactoryValidator{}
		m, err := v.parseTags(tag)

		Expect(err).To(BeNil())
		Expect(m["in"]).To(Equal("40,50,55"))
		Expect(m["regexp"]).To(Equal(`(\d+)`))
		Expect(m["email"]).To(BeEmpty())
		Expect(m["min"]).To(Equal("30"))
		Expect(m["max"]).To(Equal("100"))
	})

	It("should return validator's tag", func() {
		v := &FactoryValidator{}
		v.tag = "another_tag"
		Expect(v.Tag()).To(Equal("another_tag"))
	})

	It("should allow to set validator's tag", func() {
		v := &FactoryValidator{tag: "another_tag"}
		v.WithTag("validator")
		Expect(v.tag).To(Equal("validator"))
	})

	It("should return Checker", func() {
		v := &FactoryValidator{checkers: make(map[string]Checker)}
		v.WithChecker(&xChecker{})

		_, ok := v.Checker("y")
		Expect(ok).To(BeFalse())

		_, ok = v.Checker("x")
		Expect(ok).To(BeTrue())
	})

	It("should allow to set Checker", func() {
		v := &FactoryValidator{checkers: make(map[string]Checker)}
		Expect(len(v.checkers)).To(BeZero())

		v.WithChecker(&xChecker{})
		Expect(len(v.checkers)).To(Equal(1))
	})

	It("should return error code ERR_VALIDATOR_INVALID_TYPE", func() {
		err := New().Validate("a string")
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_INVALID_TYPE))
	})

	It("should ignore when no validation rules found", func() {
		err := New().Validate(emptyValidatorInput{})
		Expect(err).To(BeNil())
	})

	It("should return nil", func() {
		v := New().WithChecker(&yChecker{}).WithChecker(&zChecker{})

		err := v.Validate(sampleValidatorInput{})
		Expect(err).NotTo(BeNil())

		err = v.Validate(&sampleValidatorInput{Y: "not_empty"})
		Expect(err).To(BeNil())
	})

	It("should validate empty tag", func() {
		v := New().WithChecker(&yChecker{}).WithChecker(&zChecker{})
		err := v.Validate(sampleValidatorInput2{"hello"})
		Expect(err).To(BeNil())
	})

	It("should return error code ERR_VALIDATOR_INVALID_TAG", func() {
		v := New().WithChecker(&yChecker{}).WithChecker(&zChecker{})
		err := v.Validate(sampleValidatorInput3{"hello"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_INVALID_TAG))
	})

	It("should return continue when checker is not found", func() {
		v := &FactoryValidator{
			tag:      "validate",
			checkers: make(map[string]Checker),
		}
		err := v.Validate(sampleValidatorInput4{"hello"})
		Expect(err).To(BeNil())
	})

	It("should allow to set error's modifier", func() {
		v := &FactoryValidator{
			tag:      "validate",
			checkers: make(map[string]Checker),
		}
		v.WithErrorModifier(func(err errors.Error) {})
		Expect(v.errorModifier).NotTo(BeNil())
		err := v.Validate("string")
		fmt.Println(err)
		Expect(err).NotTo(BeNil())
	})

	It("should return error code ERR_VALIDATOR_UNKNOWN_ERROR if error is unknown", func() {
		v := &FactoryValidator{
			tag:      "validate",
			checkers: make(map[string]Checker),
		}
		err := v.modifyError("name", ee.New("an error"))
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_UNKNOWN_ERROR))
	})

	It("should return error code ERR_VALIDATOR_INVALID_TYPE if value is not supported", func() {
		v := &FactoryValidator{
			tag:      "validate",
			checkers: make(map[string]Checker),
		}
		_, err := v.valueOf("string")
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_INVALID_TYPE))
	})
})
