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

type Safe struct{ dial, stoppedAtZero, passedZero int }

var safeRe = regexp.MustCompile(`^(L|R)(\d+)$`)

const dialSize = 100

func (s *Safe) Follow(instructions iter.Seq[string]) {
	for instr := range instructions {
		var dir string
		var dist int
		err := input.Parse(safeRe, instr, &dir, &dist)
		if err != nil {
			panic(err) // If this changes to not be a one-off script, remember: don't panic
		}
		if dir == "L" {
			dist *= -1
		}
		s.Turn(dist)
	}
}

func (s *Safe) Turn(dist int) {
	completeTurns := dist / dialSize
	remainder := dist % dialSize
	if completeTurns < 0 {
		completeTurns = -completeTurns
	}
	s.passedZero += completeTurns

	posWithoutWrap := s.dial + remainder
	if (s.dial > 0 && posWithoutWrap < 0) || posWithoutWrap > dialSize {
		s.passedZero++
	}

	result := posWithoutWrap % dialSize
	if result < 0 {
		result += dialSize
	}

	s.dial = result
	if s.dial == 0 {
		s.stoppedAtZero++
		s.passedZero++
	}
}

func main() {
	instructions := input.SplitIntoLines(puzzleData)
	s := Safe{dial: 50}
	s.Follow(instructions)
	fmt.Printf("Part 1: %d\n", s.stoppedAtZero)
	fmt.Printf("Part 2: %d\n", s.passedZero)
}
