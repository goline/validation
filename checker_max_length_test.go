package validation

import (
	"github.com/goline/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getMaxLengthValidator() Validator {
	return New().WithChecker(MaxLengthChecker())
}

type sampleMaxLengthInput1 struct {
	Username int64 `validate:"maxLength=3"`
}

type sampleMaxLengthInput2 struct {
	Username string `validate:"maxLength=aa"`
}

type sampleMaxLengthInput3 struct {
	Username string `validate:"maxLength=3"`
}

type sampleMaxLengthInput4 struct {
	Username string `validate:"maxLength=3"`
}

var _ = Describe("MaxLengthChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_STRING", func() {
		err := getMaxLengthValidator().Validate(sampleMaxLengthInput1{10})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_STRING))
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := getMaxLengthValidator().Validate(sampleMaxLengthInput2{"aa"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_INT))
	})

	It("should return error code ERR_VALIDATOR_NOT_MAX_LENGTH", func() {
		err := getMaxLengthValidator().Validate(sampleMaxLengthInput3{"aaaa"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_MAX_LENGTH))
	})

	It("should return nil", func() {
		err := getMaxLengthValidator().Validate(sampleMaxLengthInput3{"aaa"})
		Expect(err).To(BeNil())
	})
})
