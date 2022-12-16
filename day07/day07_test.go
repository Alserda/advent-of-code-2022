package day07

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input_test string

func ExamplePart1() {
	fmt.Println(part1(input_test))

	// output: Sum of directory sizes: 95437
}

func ExamplePart2() {
	fmt.Println(part2(input_test))

	// output: Size of directory to delete: 24933642
}

func TestParser(t *testing.T) {
	dir := parse(input_test)
	assert.Equalf(t, "/", dir.Name, "expected outermost to be /, got: %s", dir.Name)
}
