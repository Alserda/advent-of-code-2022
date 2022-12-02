package day02

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Alserda/advent-of-code-2022/utils"
)

//go:embed input.txt
var input_day string

// shapeID is the identifier of the hand shape (rock, paper or scissor).
type shapeID int

const (
	rock    shapeID = 0
	paper   shapeID = 1
	scissor shapeID = 2
)

// The scores of a round
const (
	loseScore int = 0
	drawScore int = 3
	winScore  int = 6
)

// Shape is a hand shape of rock, paper or scissors.
type Shape struct {
	ID    shapeID
	wins  shapeID
	loses shapeID
}

var shapes = map[shapeID]Shape{
	rock:    {rock, scissor, paper},
	paper:   {paper, rock, scissor},
	scissor: {scissor, paper, rock},
}

// scores is a mapping of points that are given for playing a type of shape.
var scores = map[shapeID]int{
	rock:    1,
	paper:   2,
	scissor: 3,
}

// Round is the strategy that is played for one game.
type Round struct {
	opponent Shape
	ours     Shape
}

func (r Round) score() int {
	return r.selectScore() + r.outcomeScore()
}

func (r Round) selectScore() int {
	return scores[r.ours.ID]
}

func (r Round) outcomeScore() int {
	switch r.opponent.ID {
	case r.ours.wins:
		return winScore
	case r.ours.loses:
		return loseScore
	default:
		return drawScore
	}
}

func Run() {
	fmt.Println(part1(input_day))
	fmt.Println(part2(input_day))
}

func part1(input string) string {
	rounds := makeRounds(input, func(opponent OpponentColumn, ours OurColumn) Round {
		return Round{opponent.shape(), ours.shape()}
	})

	return fmt.Sprintf("total points: %d", sumPoints(rounds))
}

func part2(input string) string {
	rounds := makeRounds(input, func(opponent OpponentColumn, ours OurColumn) Round {
		os := opponent.shape()
		return Round{os, ours.elfStrategy(os)}
	})

	return fmt.Sprintf("total points: %d", sumPoints(rounds))
}

func sumPoints(rnds []Round) int {
	var pts int

	for _, round := range rnds {
		pts = pts + round.score()
	}

	return pts
}

// OpponentColumn is either A, B, C.
type OpponentColumn string

func (c OpponentColumn) shape() Shape {
	switch c {
	case "A":
		return shapes[rock]
	case "B":
		return shapes[paper]
	default:
		return shapes[scissor]
	}
}

// OurColumn is either X, Y, Z.
type OurColumn string

func (c OurColumn) shape() Shape {
	switch c {
	case "X":
		return shapes[rock]
	case "Y":
		return shapes[paper]
	default:
		return shapes[scissor]
	}
}

// shape is the Shape to determine the outcome of.
// X -> we lose
// Y -> draw
// Z -> we win
func (c OurColumn) elfStrategy(opponent Shape) Shape {
	switch c {
	case "X":
		return shapes[opponent.wins]
	case "Y":
		return shapes[opponent.ID]
	default:
		return shapes[opponent.loses]
	}
}

type RoundCreator = func(opponent OpponentColumn, ours OurColumn) Round

// makeRound parses input and calls 'rc' to determine a strategy
// for each round.
func makeRounds(input string, rc RoundCreator) []Round {
	rounds := []Round{}

	for _, v := range utils.SplitRows(input) {
		partials := strings.Split(v, " ")

		opponent := OpponentColumn(partials[0])
		ours := OurColumn(partials[1])

		rounds = append(rounds, rc(opponent, ours))
	}

	return rounds
}
