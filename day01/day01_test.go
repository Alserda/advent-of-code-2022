package day01

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input_test.txt
var input_test string

func ExampleDay01() {
	fmt.Println(part1(input_test))

	// output: elf 4 carries the most calories: 24000
}

func ExampleDay02() {
	fmt.Println(part2(input_test))

	// output: the top 3 elfs carry 45000 calories in total
}

func TestAggregate(t *testing.T) {
	inv := makeInventory(input_test)

	tests := []struct{ elf, expect int }{
		{1, 6000},
		{2, 4000},
		{3, 11000},
		{4, 24000},
		{5, 10000},
	}

	for _, test := range tests {
		name := fmt.Sprintf("elf_%d", test.elf)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expect, inv[test.elf])
		})
	}
}

func TestHighest(t *testing.T) {
	inv := makeInventory(input_test)
	elf, calories, err := inv.highest()

	assert.Nil(t, err, "got unexpected error")
	assert.Equal(t, 4, elf)
	assert.Equal(t, 24000, calories)
}

func TestOrder(t *testing.T) {
	inv := makeInventory(input_test)

	sum := 0
	for _, calories := range inv.order()[0:3] {
		sum = sum + calories
	}

	assert.Equal(t, 45000, sum)
}
