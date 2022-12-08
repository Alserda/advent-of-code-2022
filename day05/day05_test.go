package day05

import (
	_ "embed"
	"fmt"

	"testing"
)

//go:embed input_test.txt
var input_test string

func TestPart1(t *testing.T) {
	result := part1(input_test)
	fmt.Println(result)
}

func TestPart2(t *testing.T) {
	result := part2(input_test)
	fmt.Println(result)
}
