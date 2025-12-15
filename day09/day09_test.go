package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

var square = []string{"0,0", "0,1", "1,1", "1,0"}

var LShape = []string{"0,0", "0,2", "3,2", "3,5", "5,5", "5,0"}

var example = []string{
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 50)
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		polygonDescription []string
		want               int
	}{
		{name: "Square", polygonDescription: square, want: 4},
		{name: "L Shape", polygonDescription: LShape, want: 18},
		{name: "Example from description", polygonDescription: example, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolvePart2(tt.polygonDescription)

			if got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}

	assert.Equal(t, SolvePart2(square), 4)
}

func TestArea(t *testing.T) {
	a, b := Pos{2, 5}, Pos{9, 7}
	assert.Equal(t, Area(a, b), 24)
}
