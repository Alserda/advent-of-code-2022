package day03

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Alserda/advent-of-code-2022/utils"
)

//go:embed input.txt
var input string

type Rucksack = string

type Priorities map[rune]int

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) string {
	rucksacks := utils.SplitRows(input)
	prios := priorities()

	sum := 0
	for _, r := range rucksacks {
		a, b := split(r)

		for _, char := range a {
			if strings.ContainsRune(b, char) {
				sum = sum + prios[char]
				break
			}
		}
	}

	return fmt.Sprintf("the sum of the most common item type is: %d", sum)
}

func part2(input string) string {
	rucksacks := utils.SplitRows(input)
	prios := priorities()
	total := 0

	i := 0
	for i < len(rucksacks) {
	out:
		for _, char1 := range rucksacks[i] {
			for _, char2 := range rucksacks[i+1] {
				if char1 != char2 {
					continue
				}

				for _, char3 := range rucksacks[i+2] {
					if char1 == char3 {
						total = total + prios[char1]
						break out
					}
				}
			}
		}

		i += 3
	}

	return fmt.Sprintf("the sum of the most common item type is: %d", total)
}

// split splits a string in half.
func split(s string) (a, b string) {
	half := len(s) / 2

	a = s[:half]
	b = s[half:]

	return
}

// priorities creates a mapping of item type values to their value.
func priorities() Priorities {
	prio := make(Priorities)

	i := 1
	for char := 97; char <= 122; char++ {
		prio[rune(char)] = i
		i++
	}

	for char := 65; char <= 90; char++ {
		prio[rune(char)] = i
		i++
	}

	return prio
}
