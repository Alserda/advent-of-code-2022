package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitRows(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n")
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", s))
	}

	return i
}
