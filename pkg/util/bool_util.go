package util

import (
	"fmt"
	"strconv"
)

// CompareBool checks if the value as string is equals to expected value.
func CompareBool(expected bool, value string) (bool, error) {

	testValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("'%s' was not comparable with a boolean value %t", value, expected)
	}

	return testValue == expected, nil
}
