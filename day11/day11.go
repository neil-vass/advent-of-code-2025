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
	return CountProblemPaths(graph, "svr", "out")
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

func CountProblemPaths(graph Graph, start, goal string) int {
	type Path struct {
		head                 string
		passedDAC, passedFFT bool
	}

	pathsFound := 0
	frontier := fifoqueue.New(Path{start, false, false})

	// Have we been to this point before? How many problem paths did it find?
	cache := map[Path]int{}

	for !frontier.IsEmpty() {
		curr := frontier.Pull()

		if pathsFoundLastTimeWeSawThis, ok := cache[curr]; ok {
			pathsFound += pathsFoundLastTimeWeSawThis
			continue
		}

		for _, connection := range graph[curr.head] {
			if connection == goal {
				if curr.passedDAC && curr.passedFFT {
					pathsFound++
					cache[curr] = 1
				} else {
					cache[curr] = 0
				}
			} else {
				updatedPath := Path{
					head:      connection,
					passedDAC: curr.passedDAC || connection == "dac",
					passedFFT: curr.passedFFT || connection == "fft",
				}
				frontier.Push(updatedPath)
			}
		}
	}
	return pathsFound
}
