package day03

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input_test.txt
var input_test string

func ExamplePart1() {
	fmt.Println(part1(input_test))

	// output: the sum of the most common item type is: 157
}

func ExamplePart2() {
	fmt.Println(part2(input_test))

	// output: the sum of the most common item type is: 70
}

func TestSplit(t *testing.T) {
	tests := []struct{ content, first, second string }{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrVvPwwTWBwg", "PmmdzqPrV", "vPwwTWBwg"},
		{"code", "co", "de"},
	}

	for _, test := range tests {
		a, b := split(test.content)

		if a != test.first {
			t.Errorf("expect first compartment to be %s, got: %s", test.first, a)
		}

		if b != test.second {
			t.Errorf("expect first compartment to be %s, got: %s", test.first, a)
		}
	}
}

func TestPriorities(t *testing.T) {
	p := priorities()

	tests := []struct {
		char     string
		expected int
	}{
		{"a", 1},
		{"p", 16},
		{"z", 26},
		{"A", 27},
		{"Z", 52},
	}

	for _, test := range tests {
		c := []rune(test.char)

		if len(c) != 1 {
			t.Fatalf("bad test. %s has more than 1 rune", test.char)
		}

		actual := p[c[0]]

		if actual != test.expected {
			t.Errorf("expected %d, got: %d", test.expected, actual)
		}
	}
}
