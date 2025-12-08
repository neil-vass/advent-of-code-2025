package main

import (
	_ "embed"
	"fmt"
	"iter"
	"math"
	"regexp"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
	"github.com/neil-vass/advent-of-code-2025/shared/priorityqueue"
)

type Pos struct{ X, Y, Z int }
type Pair struct{ P1, P2 Pos }

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines iter.Seq[string]) int {
	return 0
}

func PairsByDistance(lines iter.Seq[string]) priorityqueue.PriorityQueue[Pair] {
	positions := []Pos{}
	pairs := priorityqueue.New[Pair]()

	for ln := range lines {
		p2 := ParsePos(ln)
		for _, p1 := range positions {
			pairs.Push(Pair{p1, p2}, Distance(p1, p2))
		}
		positions = append(positions, p2)
	}

	return pairs
}

func Distance(p1, p2 Pos) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) +
		(p1.Y-p2.Y)*(p1.Y-p2.Y) +
		(p1.Z-p2.Z)*(p1.Z-p2.Z)))
}

var posRe = regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)

func ParsePos(s string) Pos {
	var p Pos
	err := input.Parse(posRe, s, &p.X, &p.Y, &p.Z)
	if err != nil {
		panic(err) // Don't panic!
	}
	return p
}
