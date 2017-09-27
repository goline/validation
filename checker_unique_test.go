package validation

import (
	"github.com/goline/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleDatabaseFetcher struct{}

func (f *sampleDatabaseFetcher) FetchOne(table string, conditions map[string]interface{}) (interface{}, error) {
	if table == "sample_1" {
		return true, nil
	}

	return nil, nil
}

func getUniqueValidator() Validator {
	return New().WithChecker(UniqueChecker(new(sampleDatabaseFetcher)))
}

type sampleUniqueInput1 struct {
	Email string `validate:"unique"`
}

type sampleUniqueInput2 struct {
	Email string `validate:"unique=sample_1,email"`
}

type sampleUniqueInput3 struct {
	Email string `validate:"unique=sample_2,email"`
}

var _ = Describe("UniqueChecker", func() {
	It("should return error code ERR_VALIDATOR_INVALID_ARGUMENT", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput1{"e@mail.com"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_INVALID_ARGUMENT))
	})

	It("should return error code ERR_VALIDATOR_NOT_UNIQUE", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput2{"e@mail.com"})
		Expect(err).NotTo(BeNil())
		Expect(err.(errors.Error).Code()).To(Equal(ERR_VALIDATOR_NOT_UNIQUE))
	})

	It("should return nil", func() {
		err := getUniqueValidator().Validate(sampleUniqueInput3{"e@mail.com"})
		Expect(err).To(BeNil())
	})
})
