package day04

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/Alserda/advent-of-code-2022/utils"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) string {
	total := 0

	for _, row := range utils.SplitRows(input) {
		pairs := strings.Split(row, ",")

		startA, endA := assignmentRange(pairs[0])
		startB, endB := assignmentRange(pairs[1])

		if startA <= startB && endA >= endB {
			total++
			continue
		}

		if startB <= startA && endB >= endA {
			total++
			continue
		}
	}

	return fmt.Sprintf("assignment pairs where one range fully contain the other: %d", total)
}

func part2(input string) string {
	total := 0

	for _, row := range utils.SplitRows(input) {
		pairs := strings.Split(row, ",")

		startA, endA := assignmentRange(pairs[0])
		startB, endB := assignmentRange(pairs[1])

		if startA <= endB && endA >= startB {
			total++
			continue
		}
	}

	return fmt.Sprintf("assignment pairs where overlaps the other: %d", total)
}

func assignmentRange(assignment string) (int, int) {
	task := strings.Split(assignment, "-")

	start, err := strconv.Atoi(task[0])
	if err != nil {
		panic("cannot convert string to int")
	}

	end, err := strconv.Atoi(task[1])
	if err != nil {
		panic("cannot convert string to int")
	}

	return start, end
}
