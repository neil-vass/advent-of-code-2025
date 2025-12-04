package main

import (
	_ "embed"
	"fmt"
	"iter"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type Pos struct{ X, Y int }
type Empty struct{}
type Neighbours map[Pos]Empty
type Rolls map[Pos]Neighbours

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines iter.Seq[string]) int {
	rolls := RollsFromDescription(lines)
	accessibleRolls := findAccessibleRolls(rolls)
	return len(accessibleRolls)
}

func SolvePart2(lines iter.Seq[string]) int {
	count := 0
	rolls := RollsFromDescription(lines)
	accessibleRolls := findAccessibleRolls(rolls)
	for len(accessibleRolls) > 0 {
		count += len(accessibleRolls)
		for _, r := range accessibleRolls {
			delete(rolls, r)
		}
		accessibleRolls = findAccessibleRolls(rolls)
	}

	return count
}

func findAccessibleRolls(rolls Rolls) []Pos {
	result := []Pos{}
	for pos, neighbours := range rolls {
		// Remove deleted neighbours
		for n := range neighbours {
			if _, exists := rolls[n]; !exists {
				delete(neighbours, n)
			}
		}
		if len(neighbours) < 4 {
			result = append(result, pos)
		}
	}
	return result
}

func RollsFromDescription(lines iter.Seq[string]) Rolls {
	rolls := Rolls{}
	x := 0
	for ln := range lines {
		for y, val := range ln {
			if val == '@' {
				pos := Pos{x, y}
				rolls[pos] = Neighbours{}
				for _, nPos := range []Pos{{pos.X - 1, pos.Y - 1}, {pos.X - 1, pos.Y}, {pos.X - 1, pos.Y + 1}, {pos.X, pos.Y - 1}} {
					if _, exists := rolls[nPos]; exists {
						rolls[pos][nPos] = Empty{}
						rolls[nPos][pos] = Empty{}
					}
				}
			}
		}
		x++
	}

	return rolls
}
