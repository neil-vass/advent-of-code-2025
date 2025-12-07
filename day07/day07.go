package main

import (
	_ "embed"
	"fmt"
	"iter"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines iter.Seq[string]) int {
	// Set of columns that currently have beams in.
	// We update this as we move down the lines.
	beams := map[int]bool{}
	splitCount := 0

	for ln := range lines {
		for col, ch := range ln {
			switch ch {
			case 'S':
				beams[col] = true
			case '^':
				if _, ok := beams[col]; ok {
					delete(beams, col)
					beams[col-1] = true
					beams[col+1] = true
					splitCount++
				}
			}
		}
	}
	return splitCount
}

func SolvePart2(lines iter.Seq[string]) int {
	// Set of columns that currently have beams in,
	// with a count of how many beams across the worlds are there.
	// We update this as we move down the lines.
	beams := map[int]int{}
	numWorlds := 1

	for ln := range lines {
		for col, ch := range ln {
			switch ch {
			case 'S':
				beams[col] = 1
			case '^':
				if numHits, ok := beams[col]; ok {
					delete(beams, col)
					beams[col-1] += numHits
					beams[col+1] += numHits
					numWorlds += numHits
				}
			}
		}
	}
	return numWorlds
}
