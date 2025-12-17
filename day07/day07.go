package main

import (
	_ "embed"
	"fmt"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
	"github.com/neil-vass/advent-of-code-2025/shared/set"
)

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	// Set of columns that currently have beams in.
	// We update this as we move down the lines.
	beams := set.Set[int]{}
	splitCount := 0

	for _, ln := range lines {
		for col, ch := range ln {
			switch ch {
			case 'S':
				beams.Add(col)
			case '^':
				if beams.Has(col) {
					delete(beams, col)
					beams.Add(col - 1)
					beams.Add(col + 1)
					splitCount++
				}
			}
		}
	}
	return splitCount
}

func SolvePart2(lines []string) int {
	// Set of columns that currently have beams in, with a count
	// of how many beams across all the worlds are in that col.
	// We update this as we move down the lines.
	beams := map[int]int{}
	numWorlds := 1

	for _, ln := range lines {
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
