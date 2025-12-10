package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type Range struct{ Min, Max int }

var rangeRe = regexp.MustCompile(`^(\d+)-(\d+)$`)

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	ranges, idList := ParseDescription(lines)
	count := 0

	for _, id := range idList {
		for _, r := range ranges {
			if id >= r.Min && id <= r.Max {
				count++
				break
			}
		}
	}
	return count
}

func SolvePart2(lines []string) int {
	ranges, _ := ParseDescription(lines)
	mergedRanges := Merge(ranges)
	count := 0

	for _, r := range mergedRanges {
		count += r.Max - r.Min + 1
	}
	return count
}

func ParseDescription(lines []string) ([]Range, []int) {
	ranges := []Range{}
	idList := []int{}

	for _, ln := range lines {
		if strings.Contains(ln, "-") {
			var r Range
			err := input.Parse(rangeRe, ln, &r.Min, &r.Max)
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
			idList = append(idList, id)
		}
	}
	return ranges, idList
}

func Merge(ranges []Range) []Range {
	if len(ranges) == 0 {
		return []Range{}
	}

	// Sort by range mins
	slices.SortFunc(ranges, func(a, b Range) int { return a.Min - b.Min })
	head, rest := ranges[0], ranges[1:]
	merged := []Range{head}

	for len(rest) > 0 {
		head, rest = rest[0], rest[1:]
		highestSoFar := &merged[len(merged)-1]

		if head.Min <= highestSoFar.Max {
			if head.Max > highestSoFar.Max {
				highestSoFar.Max = head.Max
			}
		} else {
			merged = append(merged, head)
		}
	}
	return merged
}
