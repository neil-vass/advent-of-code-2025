package main

import (
	_ "embed"
	"fmt"
	"iter"
	"regexp"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

//go:embed input.txt
var puzzleData string

type Safe struct{ dial, zeroes int }

var safeRe = regexp.MustCompile(`^(L|R)(\d+)$`)

func (s *Safe) Follow(instructions iter.Seq[string]) {
	for instr := range instructions {
		var dir string
		var dist int
		err := input.Parse(safeRe, instr, &dir, &dist)
		if err != nil {
			panic(err)
		}
		if dir == "L" {
			dist *= -1
		}
		s.Turn(dist)
	}
}

func (s *Safe) Turn(dist int) {
	result := (s.dial + dist) % 100
	if result < 0 {
		result += 100
	}
	s.dial = result
	if s.dial == 0 {
		s.zeroes++
	}
}

func main() {
	instructions := input.SplitIntoLines(puzzleData)
	s := Safe{dial: 50}
	s.Follow(instructions)
	fmt.Printf("Part 1: %d\n", s.zeroes)
}
