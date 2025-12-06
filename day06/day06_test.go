package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

var example = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 4277556)
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, SolvePart2(example), 3263827)
}