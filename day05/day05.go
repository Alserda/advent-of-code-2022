package day05

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
	puzzle := parse(input)
	puzzle.run()

	return puzzle.uppers()
}

func part2(input string) string {
	puzzle := parse(input)
	puzzle.run2()

	return puzzle.uppers()
}

type Instruction struct {
	Amount, From, To int
}

type Stack map[int][]string

type Puzzle struct {
	Stack        Stack
	Instructions []Instruction
}

func (p *Puzzle) uppers() string {
	// result := map[int]string{}

	var msg string
	for i := 1; i <= len(p.Stack); i++ {
		s := p.Stack[i]
		// result[i] = s[len(s)-1]
		msg += s[len(s)-1]
		// msg += fmt.Sprintf("stack %d: %s \n", i, s[len(s)-1])
	}

	return msg
}

func (p *Puzzle) run() {
	s := p.Stack

	for _, ins := range p.Instructions {
		from, to := s[ins.From], s[ins.To]

		for i := 1; i <= ins.Amount; i++ {
			crate := from[len(from)-1]

			from = from[:len(from)-1]
			to = append(to, crate)
		}

		s[ins.From] = from
		s[ins.To] = to
	}

	p.Stack = s
}

func (p *Puzzle) run2() {
	s := p.Stack

	for _, ins := range p.Instructions {
		from, to := s[ins.From], s[ins.To]
		crates := from[len(from)-ins.Amount:]

		to = append(to, crates...)
		from = from[:len(from)-ins.Amount]

		s[ins.From] = from
		s[ins.To] = to
	}

	p.Stack = s
}

func parse(input string) Puzzle {
	puzzle := Puzzle{
		Stack:        map[int][]string{},
		Instructions: []Instruction{},
	}

	for _, row := range utils.SplitRows(input) {
		if row == "" {
			continue
		}

		if strings.Contains(row, "[") {
			crateID := 1

			i := 0
			for i < len(row) {
				crate := row[i+1 : i+2]

				if crate != " " {
					puzzle.Stack[crateID] = append(puzzle.Stack[crateID], crate)
				}

				i += 4
				crateID++
			}
		}

		if strings.Contains(row, "move") {
			parts := strings.Split(row, " ")
			puzzle.Instructions = append(puzzle.Instructions, Instruction{
				Amount: parseInt(parts[1]),
				From:   parseInt(parts[3]),
				To:     parseInt(parts[5]),
			})
		}
	}

	for i, stack := range puzzle.Stack {
		puzzle.Stack[i] = reverse(stack)
	}

	return puzzle
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", s))
	}

	return i
}

func reverse[T any](a []T) []T {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}

	return a
}
