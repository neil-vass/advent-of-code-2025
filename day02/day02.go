package main

import (
	"regexp"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

var rangeRe = regexp.MustCompile(`^(\d+)-(\d+)$`)

func ParseRange(rangeStr string) (int, int) {
	var min, max int
	err := input.Parse(rangeRe, rangeStr, &min, &max)
	if err != nil {
		panic(err) // If this changes to not be a one-off script, remember: don't panic
	}
	return min, max
}

func InvalidIDs(s string) []int {
	result := []int{}
	for _, rangeStr := range strings.Split(s, ",") {
		min, max := ParseRange(rangeStr)
		if isInvalidID(min) {
			result = append(result, min)
		}
		curr := min
		for curr < max {
			curr = nextInvalidID(curr)
			if curr <= max {
				result = append(result, curr)
			}
		}
	}
	return result
}

func isInvalidID(min int) bool {
	return true
}

func nextInvalidID(curr int) int {
	return 22
}
