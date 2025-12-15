package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
}

func SolvePart1(lines []string) int {
	model := ParseInput(lines)
	yes, _, maybe := Buckets(model)
	if len(maybe) > 0 {
		panic("We have some real work to do")
	}
	return len(yes)
}

type Present struct {
	Size int
}

type Tree struct {
	Length, Width int
	Needs         []int
}

type Model struct {
	Presents []Present
	Trees    []Tree
}

var numColonRe = regexp.MustCompile(`^\d:$`)
var treeDimensionsRe = regexp.MustCompile(`^(\d+)x(\d+):$`)

func ParseInput(lines []string) Model {
	model := Model{[]Present{}, []Tree{}}
	i := 0
	for i < len(lines) {
		if numColonRe.MatchString(lines[i]) {
			present := Present{}
			for lines[i] != "" {
				i++
				present.Size += strings.Count(lines[i], "#")
			}
			model.Presents = append(model.Presents, present)
			i++
		} else {
			tree := Tree{}
			fields := strings.Fields(lines[i])
			input.Parse(treeDimensionsRe, fields[0], &tree.Length, &tree.Width)
			tree.Needs = make([]int, len(fields)-1)
			for i, str := range fields[1:] {
				tree.Needs[i], _ = strconv.Atoi(str)
			}
			model.Trees = append(model.Trees, tree)
			i++
		}
	}
	return model
}

func Buckets(model Model) ([]Tree, []Tree, []Tree) {
	yes, no, maybe := []Tree{}, []Tree{}, []Tree{}

	for _, tree := range model.Trees {
		availableArea := tree.Length * tree.Width

		presentCount := 0
		totalPresentSize := 0
		for i, n := range tree.Needs {
			presentCount += n
			totalPresentSize += n * model.Presents[i].Size
		}

		if availableArea < totalPresentSize {
			no = append(no, tree)
			continue
		}

		areaNeededWithoutPacking := 9 * presentCount
		if availableArea >= areaNeededWithoutPacking {
			yes = append(yes, tree)
			continue
		}

		maybe = append(maybe, tree)
	}
	return yes, no, maybe
}
