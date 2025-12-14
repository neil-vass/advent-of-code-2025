package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

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
