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
	fmt.Printf("Part1: %d\n", SolvePart1(lines))
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

func SolvePart1(lines iter.Seq[string]) int {
	rolls := RollsFromDescription(lines)
	accessibleRolls := 0
	for _, neighbours := range rolls {
		if len(neighbours) < 4 {
			accessibleRolls++
		}
	}
	return accessibleRolls
}
