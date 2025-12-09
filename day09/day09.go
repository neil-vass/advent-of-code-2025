package main

import (
	_ "embed"
	"fmt"
	"iter"
	"regexp"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type Pos struct{ X, Y int }

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines iter.Seq[string]) int {
	corners := ParseTiles(lines)
	maxArea := 0
	for i := 0; i < len(corners)-1; i++ {
		for j := i + 1; j < len(corners); j++ {
			area := Area(corners[i], corners[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func SolvePart2(lines iter.Seq[string]) int {
	return 0
}

var tileRe = regexp.MustCompile(`^(\d+),(\d+)$`)

func ParseTiles(lines iter.Seq[string]) []Pos {
	tiles := []Pos{}
	for ln := range lines {
		var p Pos
		err := input.Parse(tileRe, ln, &p.X, &p.Y)
		if err != nil {
			panic(err) // Don't panic!
		}
		tiles = append(tiles, p)
	}
	return tiles
}

func Area(a, b Pos) int {
	length := a.X - b.X
	if length < 0 {
		length = -length
	}
	length += 1
	height := a.Y - b.Y
	if height < 0 {
		height = -height
	}
	height += 1
	return length * height
}
