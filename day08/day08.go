package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

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
	lines := strings.Split(strings.TrimSpace(puzzleData), "\n")
	fmt.Printf("Part 1: %d\n", SolvePart1(lines, 1000))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines, 1000))
}

func SolvePart1(lines []string, cables int) int {
	pairs := PairsByDistance(lines)
	circuits := Connect(pairs, cables)
	slices.SortFunc(circuits, func(a, b Set[Pos]) int {
		return len(b) - len(a)
	})
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func SolvePart2(lines []string, cables int) int {
	pairs := PairsByDistance(lines)
	circuits := Connect(pairs, cables)

	for {
		if pairs.IsEmpty() {
			panic("No pairs left, no solution found!")
		}
		pair := pairs.Pull()

		// add this pair, merging circuits as needed.
		p1Circuit := -1
		p2Circuit := -1
		for i, c := range circuits {
			if c.Has(pair.P1) {
				p1Circuit = i
			}
			if c.Has(pair.P2) {
				p2Circuit = i
			}
			if p1Circuit != -1 && p2Circuit != -1 {
				break
			}
		}

		if p1Circuit == -1 && p2Circuit == -1 {
			newCircuit := Set[Pos]{}
			newCircuit.Add(pair.P1)
			newCircuit.Add(pair.P2)
			circuits = append(circuits, newCircuit)
		} else if p1Circuit == -1 {
			circuits[p2Circuit].Add(pair.P1)
		} else if p2Circuit == -1 {
			circuits[p1Circuit].Add(pair.P2)
		} else if p1Circuit != p2Circuit {
			for item := range circuits[p1Circuit] {
				circuits[p2Circuit].Add(item)
			}
			last := len(circuits) - 1
			circuits[p1Circuit] = circuits[last]
			circuits[last] = nil
			circuits = circuits[:last]
		}

		if len(circuits) == 1 && len(circuits[0]) == len(lines) {
			return pair.P1.X * pair.P2.X
		}
	}
}

func PairsByDistance(lines []string) priorityqueue.PriorityQueue[Pair] {
	positions := []Pos{}
	pairs := priorityqueue.New[Pair]()

	for _, ln := range lines {
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

func Connect(pairs priorityqueue.PriorityQueue[Pair], cables int) []Set[Pos] {
	graph := buildConnectionGraph(cables, pairs)

	circuits := []Set[Pos]{}

	key, ok := getAnyKey(graph)
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
		key, ok = getAnyKey(graph)
	}

	return circuits
}

func getAnyKey[TKey comparable, TValue any](m map[TKey]TValue) (TKey, bool) {
	for k := range m {
		return k, true
	}
	return *new(TKey), false
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
