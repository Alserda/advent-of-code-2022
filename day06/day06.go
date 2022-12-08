package day06

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("start-of-packet marker: %d\n", search(input, 4))
	fmt.Printf("start-of-message marker: %d\n", search(input, 14))
}

func search(input string, charCount int) int {
	var received []rune

	for i, c := range input {
		for j, r := range received {
			if c == r {
				received = received[j:]
			}
		}

		if len(received) < charCount {
			received = append(received, c)
			continue
		}

		return i + 1
	}

	return -1
}
