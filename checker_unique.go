package validation

import (
	"fmt"

	"github.com/goline/errors"
	"strings"
)

func UniqueChecker(fetcher DatabaseFetcher) Checker {
	return &uniqueChecker{
		fetcher: fetcher,
	}
}

type uniqueChecker struct {
	fetcher DatabaseFetcher
}

func (c *uniqueChecker) Name() string {
	return "unique"
}

func (c *uniqueChecker) Check(v interface{}, expects string) error {
	var table, column string
	a := strings.Split(expects, ",")
	switch len(a) {
	case 2:
		table = a[0]
		column = a[1]
	default:
		return errors.New(ERR_VALIDATOR_INVALID_ARGUMENT, "Invalid expectation argument")
	}

	conditions := make(map[string]interface{})
	conditions[column] = v
	row, _ := c.fetcher.FetchOne(table, conditions)
	if row != nil {
		return errors.New(ERR_VALIDATOR_NOT_UNIQUE, fmt.Sprintf("%v already exists. It must be unique", v))
	}

	return nil
}
