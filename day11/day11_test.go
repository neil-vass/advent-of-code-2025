package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestSolvePart1(t *testing.T) {
	example := []string{
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
	assert.Equal(t, SolvePart1(example), 5)
}

func TestSolvePart2(t *testing.T) {
	example := []string{
		"svr: aaa bbb",
		"aaa: fft",
		"fft: ccc",
		"bbb: tty",
		"tty: ccc",
		"ccc: ddd eee",
		"ddd: hub",
		"hub: fff",
		"eee: dac",
		"dac: fff",
		"fff: ggg hhh",
		"ggg: out",
		"hhh: out",
	}
	assert.Equal(t, SolvePart2(example), 2)
}
