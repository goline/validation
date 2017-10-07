package validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
		err := New().Validate(&sampleMaxInputTest1{Age: 10})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))
	})

	It("should return error code ERR_VALIDATOR_NOT_NUMBER", func() {
		err := New().Validate(&sampleMaxInputTest1{Age: 10})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))

		err = New().Validate(&sampleMaxInputTest2{Age: "10"})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_NUMBER))
	})

	It("should return error code ERR_VALIDATOR_NOT_MAX (int)", func() {
		err := New().Validate(&sampleMaxInputTest3{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_MAX))
	})

	It("should return nil (float)", func() {
		err := New().Validate(&sampleMaxInputTest4{Age: 10.1})
		Expect(err).To(BeNil())
	})

	It("should return nil (int)", func() {
		err := New().Validate(&sampleMaxInputTest5{Age: 9})
		Expect(err).To(BeNil())
	})

	It("should return error code ERR_VALIDATOR_NOT_INT", func() {
		err := New().Validate(&sampleMaxInputTest6{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_INT))
	})

	It("should return error code ERR_VALIDATOR_NOT_FLOAT", func() {
		err := New().Validate(&sampleMaxInputTest7{Age: 11})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_FLOAT))
	})

	It("should return error code ERR_VALIDATOR_NOT_MAX (float)", func() {
		err := New().Validate(&sampleMaxInputTest8{Age: 10.2})
		Expect(err).NotTo(BeNil())
		Expect(err.Code()).To(Equal(ERR_VALIDATOR_NOT_MAX))
	})
})
