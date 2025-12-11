package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

var example = []string{
	"aaa: you hhh",
	"you: bbb ccc",
	"bbb: ddd eee",
	"ccc: ddd eee fff",
	"ddd: ggg",
	"eee: out",
	"fff: out",
	"ggg: out",
	"hhh: ccc fff iii",
	"iii: out",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 5)
}
