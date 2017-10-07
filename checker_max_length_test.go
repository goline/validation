package validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
		err := New().Validate(sampleMaxLengthInput1{10})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_STRING))
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := New().Validate(sampleMaxLengthInput2{"aa"})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_INT))
	})

	It("should return error code ERR_VALIDATOR_NOT_MAX_LENGTH", func() {
		err := New().Validate(sampleMaxLengthInput3{"aaaa"})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_MAX_LENGTH))
	})

	It("should return nil", func() {
		err := New().Validate(sampleMaxLengthInput3{"aaa"})
		Expect(err).To(BeNil())
	})
})
