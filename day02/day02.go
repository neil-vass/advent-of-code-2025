package main

import (
	_ "embed"
	"fmt"
	"math"
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
	fmt.Printf("Part1: %d\n", Solve(puzzleData))
}

func Solve(s string) int {
	total := 0
	for _, id := range InvalidIDs(s) {
		total += id
	}
	return total
}

func InvalidIDs(s string) []int {
	result := []int{}
	for rangeStr := range strings.SplitSeq(s, ",") {
		min, max := ParseRange(rangeStr)
		if IsInvalidID(min) {
			result = append(result, min)
		}
		curr := min
		for curr < max {
			// curr = NextInvalidID(curr)
			// if curr <= max {
			// 	result = append(result, curr)
			// }
			curr++
			if IsInvalidID(curr) {
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

func IsInvalidID(n int) bool {
	digits := strconv.Itoa(n)
	if len(digits)%2 != 0 {
		return false
	}
	half := len(digits) / 2
	return digits[:half] == digits[half:]
}

func NextInvalidID(n int) int {
	digits := strconv.Itoa(n)
	if len(digits)%2 != 0 {
		return NextInvalidID(int(math.Pow10(len(digits))))
	}
	half := len(digits) / 2
	left, _ := strconv.Atoi(digits[:half])
	right, _ := strconv.Atoi(digits[half:])

	if left > right {
		nextID, _ := strconv.Atoi(strconv.Itoa(left) + strconv.Itoa(left))
		return nextID
	} else if (left+1)%10 != 0 {
		nextID, _ := strconv.Atoi(strconv.Itoa(left+1) + strconv.Itoa(left+1))
		return nextID
	} else {
		return NextInvalidID(int(math.Pow10(len(digits) + 1)))
	}
}
