package validation

import (
	"github.com/goline/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getMaxValidator() Validator {
	return New().WithChecker(MaxChecker())
}

type sampleMaxInputTest1 struct {
	Age int64 `validate:"max=aa"`
}

type sampleMaxInputTest2 struct {
	Age string `validate:"max=10"`
}

type sampleMaxInputTest3 struct {
	Age int64 `validate:"max=10"`
}

type sampleMaxInputTest4 struct {
	Age float64 `validate:"max=10.2"`
}

type sampleMaxInputTest5 struct {
	Age int64 `validate:"max=10"`
}

type sampleMaxInputTest6 struct {
	Age float64 `validate:"max=10"`
}

type sampleMaxInputTest7 struct {
	Age int64 `validate:"max=10.2"`
}

type sampleMaxInputTest8 struct {
	Age float64 `validate:"max=10.1"`
}

var _ = Describe("MaxChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_NUMBER", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest1{Age: 10})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))
	})

	It("should return error code ERR_VALIDATOR_NOT_NUMBER", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest1{Age: 10})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))

		err = getMaxValidator().Validate(&sampleMaxInputTest2{Age: "10"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))
	})

	It("should return error code ERR_VALIDATOR_NOT_MAX (int)", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest3{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_MAX))
	})

	It("should return nil (float)", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest4{Age: 10.1})
		Expect(err).To(BeNil())
	})

	It("should return nil (int)", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest5{Age: 9})
		Expect(err).To(BeNil())
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest6{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_INT))
	})

	It("should return error code ERR_VALIDATOR_NOT_FLOAT", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest7{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_FLOAT))
	})

	It("should return error code ERR_VALIDATOR_NOT_MAX (float)", func() {
		err := getMaxValidator().Validate(&sampleMaxInputTest8{Age: 10.2})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_MAX))
	})
})
