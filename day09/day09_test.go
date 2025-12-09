package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

var example = input.Lines(
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 50)
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, SolvePart2(example), 24)
}

func TestArea(t *testing.T) {
	a, b := Pos{2, 5}, Pos{9, 7}
	assert.Equal(t, Area(a, b), 24)
}
