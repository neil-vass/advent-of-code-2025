package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

func TestSolve_Part1(t *testing.T) {
	banks := input.Lines(
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	)
	assert.Equal(t, Solve(banks, 2), 357)
}

func TestMaxJoltage_Part1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		bank string
		want int
	}{
		{name: "First 2 chars", bank: "987654321111111", want: 98},
		{name: "First and last", bank: "811111111111119", want: 89},
		{name: "Last 2 chars", bank: "234234234234278", want: 78},
		{name: "Chars inside", bank: "818181911112111", want: 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxJoltage(tt.bank, 2)
			if got != tt.want {
				t.Errorf("MaxJoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}
