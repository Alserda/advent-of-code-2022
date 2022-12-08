package day06

import "testing"

type TestItem struct {
	buffer string
	marker int
}

type Test struct {
	t         *testing.T
	charCount int
	items     []TestItem
}

func TestPart1(t *testing.T) {
	test := Test{
		t:         t,
		charCount: 4,
		items: []TestItem{
			{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
			{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
			{"nppdvjthqldpwncqszvftbrmjlhg", 6},
			{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
			{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
		},
	}

	runTests(test)
}

func TestPart2(t *testing.T) {
	test := Test{
		t:         t,
		charCount: 14,
		items: []TestItem{
			{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
			{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
			{"nppdvjthqldpwncqszvftbrmjlhg", 23},
			{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
			{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
		},
	}

	runTests(test)
}

func runTests(test Test) {
	for _, ti := range test.items {
		marker := search(ti.buffer, test.charCount)

		test.t.Run(ti.buffer, func(t *testing.T) {
			if marker != ti.marker {
				t.Errorf("expected marker %d, got: %d", ti.marker, marker)
			}
		})
	}
}
