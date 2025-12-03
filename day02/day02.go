package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

var rangeRe = regexp.MustCompile(`^(\d+)-(\d+)$`)

//go:embed input.txt
var puzzleData string

func main() {
	puzzleData = strings.TrimSpace(puzzleData)
	fmt.Printf("Part1: %d\n", Solve(puzzleData, IsInvalidID_Part1))
	fmt.Printf("Part2: %d\n", Solve(puzzleData, IsInvalidID_Part2))
}

func Solve(s string, checkFn func(int) bool) int {
	total := 0
	for _, id := range InvalidIDs(s, checkFn) {
		total += id
	}
	return total
}

func InvalidIDs(s string, checkFn func(int) bool) []int {
	result := []int{}
	for rangeStr := range strings.SplitSeq(s, ",") {
		min, max := ParseRange(rangeStr)
		for curr := min; curr <= max; curr++ {
			if checkFn(curr) {
				result = append(result, curr)
			}
		}
	}
	return result
}

func ParseRange(rangeStr string) (int, int) {
	var min, max int
	err := input.Parse(rangeRe, rangeStr, &min, &max)
	if err != nil {
		panic(err) // If this changes to not be a one-off script, remember: don't panic
	}
	return min, max
}

func IsInvalidID_Part1(n int) bool {
	digits := strconv.Itoa(n)
	if len(digits)%2 != 0 {
		return false
	}
	half := len(digits) / 2
	return digits[:half] == digits[half:]
}

func IsInvalidID_Part2(n int) bool {
	digits := strconv.Itoa(n)
	for chunks := 2; chunks <= len(digits); chunks++ {
		if len(digits)%chunks != 0 {
			continue
		}
		size := len(digits) / chunks
		first := digits[:size]
		foundMismatch := false
		for step := 1; step < chunks; step++ {
			other := digits[(step * size) : (step+1)*size]
			if first != other {
				foundMismatch = true
				continue
			}
		}
		if !foundMismatch {
			return true
		}

		if chunks >= len(digits)/2 {
			chunks = len(digits)
		}
	}
	return false
}
