package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/neil-vass/advent-of-code-2025/shared/assert"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

func TestSolvePart1(t *testing.T) {
	example := input.Lines(
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	)
	assert.Equal(t, SolvePart1(example), 13)
}

func TestSolvePart2(t *testing.T) {
	example := input.Lines(
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	)
	assert.Equal(t, SolvePart2(example), 43)
}

func TestRollsFromDescription(t *testing.T) {
	lines := input.Lines(
		"@..",
		"@.@",
	)
	got := RollsFromDescription(lines)
	want := Rolls{
		{X: 0, Y: 0}: {{X: 1, Y: 0}: Empty{}},
		{X: 1, Y: 0}: {{X: 0, Y: 0}: Empty{}},
		{X: 1, Y: 2}: {},
	}
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}
