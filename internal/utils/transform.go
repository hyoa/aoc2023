package utils

import (
	"strconv"
)

func ExpectedToInt(expected string) int {
	if expected == "" {
		return 0
	}

	v, err := strconv.Atoi(expected)
	if err != nil {
		return 0
	}

	return v
}
