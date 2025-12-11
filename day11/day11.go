package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/fifoqueue"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type Graph map[string][]string

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	graph := ParseGraph(lines)
	return CountPaths(graph, "you", "out")
}

func SolvePart2(lines []string) int {
	graph := ParseGraph(lines)
	return CountProblemPaths(graph)
}

func ParseGraph(lines []string) Graph {
	graph := make(Graph, len(lines))
	for _, ln := range lines {
		devSplit := strings.Split(ln, ": ")
		device, connections := devSplit[0], strings.Split(devSplit[1], " ")
		graph[device] = connections
	}
	return graph
}

func CountPaths(graph Graph, start, goal string) int {
	pathsFound := 0
	frontier := fifoqueue.New(start)

	for !frontier.IsEmpty() {
		curr := frontier.Pull()
		for _, connection := range graph[curr] {

			if connection == goal {
				pathsFound++
			} else {
				frontier.Push(connection)
			}
		}
	}
	return pathsFound
}

type path struct {
	head                 string
	passedDAC, passedFFT bool
}

func CountProblemPaths(graph Graph) int {
	start := path{"svr", false, false}
	cache := map[path]int{}
	return ProblemPathsFrom(start, graph, cache)
}

func ProblemPathsFrom(curr path, graph Graph, cache map[path]int) int {
	if pathCount, ok := cache[curr]; ok {
		return pathCount
	}

	if curr.head == "out" {
		if curr.passedDAC && curr.passedFFT {
			return 1
		} else {
			return 0
		}
	}

	pathsFromHere := 0
	for _, connection := range graph[curr.head] {
		nextStep := path{
			head:      connection,
			passedDAC: curr.passedDAC || connection == "dac",
			passedFFT: curr.passedFFT || connection == "fft",
		}
		pathsFromHere += ProblemPathsFrom(nextStep, graph, cache)
	}
	cache[curr] = pathsFromHere
	return pathsFromHere
}
