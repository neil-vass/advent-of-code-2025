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
	//fmt.Printf("Part 2: %d\n", Solve(lines, FewestPressesForJoltage))
}

func SolvePart1(lines []string) int {
	graph := ParseGraph(lines)
	return CountPaths(graph, "you", "out")
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
