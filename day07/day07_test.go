package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

var example = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 21)
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, SolvePart2(example), 40)
}
