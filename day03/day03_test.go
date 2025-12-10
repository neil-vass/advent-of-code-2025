package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestSolve_Part1(t *testing.T) {
	banks := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	assert.Equal(t, Solve(banks, 2), 357)
}

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		bank    string
		numBats int
		want    int
	}{
		{name: "First 2 chars", bank: "987654321111111", numBats: 2, want: 98},
		{name: "First and last", bank: "811111111111119", numBats: 2, want: 89},
		{name: "Last 2 chars", bank: "234234234234278", numBats: 2, want: 78},
		{name: "2 chars inside", bank: "818181911112111", numBats: 2, want: 92},
		{name: "First 12 chars", bank: "987654321111111", numBats: 12, want: 987654321111},
		{name: "12 chars inside", bank: "234234234234278", numBats: 12, want: 434234234278},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxJoltage(tt.bank, tt.numBats)
			if got != tt.want {
				t.Errorf("MaxJoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}
