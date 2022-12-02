package day02

import (
	"fmt"

	_ "embed"
)

//go:embed input_test.txt
var input_test string

func ExamplePart1() {
	fmt.Println(part1(input_test))

	// output: total points: 15
}

func ExamplePart2() {
	fmt.Println(part2(input_test))

	// output: total points: 12
}
