package main

import (
	_ "embed"
	"fmt"
	"iter"
	"math"
	"regexp"
	"slices"

	"github.com/neil-vass/advent-of-code-2025/shared/fifoqueue"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
	"github.com/neil-vass/advent-of-code-2025/shared/priorityqueue"
)

type Pos struct{ X, Y, Z int }
type Pair struct{ P1, P2 Pos }

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines, 1000))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines iter.Seq[string], cables int) int {
	circuits := Connect(lines, cables)
	slices.SortFunc(circuits, func(a, b Set[Pos]) int {
		return len(b) - len(a)
	})
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
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

func Connect(lines iter.Seq[string], cables int) []Set[Pos] {
	pairs := PairsByDistance(lines)

	graph := buildConnectionGraph(cables, pairs)

	getAny := func() (Pos, bool) {
		for k := range graph {
			return k, true
		}
		return *new(Pos), false
	}

	circuits := []Set[Pos]{}

	key, ok := getAny()
	for ok {
		frontier := fifoqueue.New(key)
		reached := Set[Pos]{}
		reached.Add(key)

		for !frontier.IsEmpty() {
			curr := frontier.Pull()
			for n := range graph[curr] {
				if !reached.Has(n) {
					frontier.Push(n)
					reached.Add(n)
				}
			}
		}

		circuits = append(circuits, reached)
		for p := range reached {
			delete(graph, p)
		}
		key, ok = getAny()
	}

	return circuits
}

func buildConnectionGraph(cables int, pairs priorityqueue.PriorityQueue[Pair]) map[Pos]Set[Pos] {
	graph := map[Pos]Set[Pos]{}
	for range cables {
		pair := pairs.Pull()
		neighbours, ok := graph[pair.P1]
		if !ok {
			neighbours = Set[Pos]{}
			graph[pair.P1] = neighbours
		}
		neighbours.Add(pair.P2)
		neighbours, ok = graph[pair.P2]
		if !ok {
			neighbours = Set[Pos]{}
			graph[pair.P2] = neighbours
		}
		neighbours.Add(pair.P1)
	}
	return graph
}
