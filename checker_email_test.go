package validation

import (
	"github.com/goline/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleEmailInputTest1 struct {
	Email int64 `validate:"email"`
}

type sampleEmailInputTest2 struct {
	Email string `validate:"email"`
}

var _ = Describe("EmailChecker", func() {
	It("should return error code ERR_VALIDATOR_NOT_STRING", func() {
		err := New().Validate(sampleEmailInputTest1{5})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_STRING))
	})

	It("should return error code ERR_VALIDATOR_NOT_EMAIL", func() {
		cases := []string{
			"mail.abc.com",
			"###@ab.@@",
		}

		for _, email := range cases {
			err := New().Validate(sampleEmailInputTest2{email})
			Expect(err).NotTo(BeNil())
			Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_EMAIL))
		}
	})

	It("should return nil", func() {
		cases := []string{
			"mail#&*&@abc.com",
			"778@abc.com",
		}

		for _, email := range cases {
			err := New().Validate(sampleEmailInputTest2{email})
			Expect(err).To(BeNil())
		}
	})
})
