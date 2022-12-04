package day04

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input_test.txt
var input_test string

func ExamplePart1() {
	fmt.Println(part1(input_test))

	// assignment pairs where one range fully contain the other: 2
}

func ExamplePart2() {
	fmt.Println(part2(input_test))

	// assignment pairs where one overlaps the other: 4
}

func TestPart1(t *testing.T) {
	result := part1(input_test)

	fmt.Println(result)
}

func TestPart2(t *testing.T) {
	result := part2(input_test)

	fmt.Println(result)
}
