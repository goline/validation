package validation

import (
	"github.com/goline/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getRangeValidator() Validator {
	return New().WithChecker(RangeChecker())
}

type sampleRangeInputTest1 struct {
	Age string `validate:"range=18-60"`
}

type sampleRangeInputTest2 struct {
	Age int64 `validate:"range=18-60a"`
}

type sampleRangeInputTest3 struct {
	Age int64 `validate:"range=18-60"`
}

type sampleRangeInputTest4 struct {
	Age float64 `validate:"range=18-60a"`
}

type sampleRangeInputTest5 struct {
	Age float64 `validate:"range=10.0-10.2"`
}

type sampleRangeInputTest6 struct {
	Age int64 `validate:"range=9-11"`
}

type sampleRangeInputTest7 struct {
	Age float64 `validate:"range=10.0-10.2"`
}

var _ = Describe("MinLengthChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_NUMBER", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest1{"10"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))
	})

	It("should return error code ERR_VALIDATOR_INVALID_FORMAT", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest2{10})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_INVALID_FORMAT))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_RANGE", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest3{10})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_IN_RANGE))
	})

	It("should return error code ERR_VALIDATOR_INVALID_FORMAT (float)", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest4{10.2})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_INVALID_FORMAT))
	})

	It("should return error code ERR_VALIDATOR_NOT_IN_RANGE (float)", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest5{9.9})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_IN_RANGE))
	})

	It("should return nil (int)", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest6{10})
		Expect(err).To(BeNil())
	})

	It("should return nil (float)", func() {
		err := getMinLengthValidator().Validate(sampleRangeInputTest7{10.1})
		Expect(err).To(BeNil())
	})
})
