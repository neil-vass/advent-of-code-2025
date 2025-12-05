package main

import (
	_ "embed"
	"fmt"
	"iter"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type Range struct{ min, max int }

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", Solve(lines))
}

var rangeRe = regexp.MustCompile(`^(\d+)-(\d+)$`)

func Solve(lines iter.Seq[string]) int {
	ranges := []Range{}
	count := 0
	for ln := range lines {
		if strings.Contains(ln, "-") {
			var r Range
			err := input.Parse(rangeRe, ln, &r.min, &r.max)
			if err != nil {
				panic(err) // If this changes to not be a one-off script, remember: don't panic
			}
			ranges = append(ranges, r)
		} else if ln == "" {
			continue
		} else {
			id, err := strconv.Atoi(ln)
			if err != nil {
				panic(err) // If this changes to not be a one-off script, remember: don't panic
			}

			for _, r := range ranges {
				if id >= r.min && id <= r.max {
					count++
					break
				}
			}
		}
	}
	return count
}
