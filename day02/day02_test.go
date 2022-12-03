package day02

import (
	"fmt"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
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

func TestScores(t *testing.T) {
	assert := assert.New(t)

	// we draw
	round := Round{Shapes[Rock], Shapes[Rock]}
	assert.Equal(1, round.selectScore())
	assert.Equal(3, round.outcomeScore())
	assert.Equal(4, round.score())

	// we win
	round = Round{Shapes[Rock], Shapes[Paper]}
	assert.Equal(2, round.selectScore())
	assert.Equal(6, round.outcomeScore())
	assert.Equal(8, round.score())

	// we lose
	round = Round{Shapes[Rock], Shapes[Scissor]}
	assert.Equal(3, round.selectScore())
	assert.Equal(0, round.outcomeScore())
	assert.Equal(3, round.score())
}

func TestSumPoints(t *testing.T) {
	rounds := []Round{
		{Shapes[Rock], Shapes[Rock]},
		{Shapes[Rock], Shapes[Paper]},
		{Shapes[Rock], Shapes[Scissor]},
	}

	assert.Equal(t, 15, sumPoints(rounds))
}

func TestOpponentShapeMatcher(t *testing.T) {
	tests := []struct {
		column string
		expect Shape
	}{
		{"A", Shapes[Rock]},
		{"B", Shapes[Paper]},
		{"C", Shapes[Scissor]},
	}

	for _, test := range tests {
		t.Run(test.column, func(t *testing.T) {
			oc := OpponentColumn(test.column)
			assert.Equal(t, test.expect, oc.shape())
		})
	}
}

func TestOurShapeMatcher(t *testing.T) {
	tests := []struct {
		column string
		expect Shape
	}{
		{"X", Shapes[Rock]},
		{"Y", Shapes[Paper]},
		{"Z", Shapes[Scissor]},
	}

	for _, test := range tests {
		t.Run(test.column, func(t *testing.T) {
			oc := OurColumn(test.column)
			assert.Equal(t, test.expect, oc.shape())
		})
	}
}

func TestElfStrategy(t *testing.T) {
	tests := []struct {
		opponent Shape
		column   string
		expect   Shape
	}{
		{Shapes[Rock], "X", Shapes[Scissor]},
		{Shapes[Rock], "Y", Shapes[Rock]},
		{Shapes[Rock], "Z", Shapes[Paper]},
	}

	for _, test := range tests {
		t.Run(test.column, func(t *testing.T) {
			oc := OurColumn(test.column)

			assert.Equal(t, test.expect, oc.elfStrategy(test.opponent))
		})
	}
}
