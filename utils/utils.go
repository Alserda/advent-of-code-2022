package utils

import "strings"

func SplitRows(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n")
}
