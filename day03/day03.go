package main

import (
	_ "embed"
	"fmt"
	"iter"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

//go:embed input.txt
var puzzleData string

func main() {
	banks := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part1: %d\n", Solve(banks, 2))
	fmt.Printf("Part1: %d\n", Solve(banks, 12))
}

func Solve(banks iter.Seq[string], numBats int) int {
	total := 0
	for bank := range banks {
		total += MaxJoltage(bank, numBats)
	}
	return total
}

func MaxJoltage(bank string, numBats int) int {
	digits := make([]int, numBats)

	for i, b := range bank {
		val := int(b - '0')
		batsRemaining := len(bank) - i

		for digitIdx := range digits {
			batsNeeded := len(digits) - digitIdx
			if batsRemaining >= batsNeeded && digits[digitIdx] < val {
				digits[digitIdx] = val
				for otherIdx := digitIdx + 1; otherIdx < len(digits); otherIdx++ {
					digits[otherIdx] = 0
				}
				break
			}
		}
	}

	result := 0
	for _, digit := range digits {
		result *= 10
		result += digit
	}
	return result
}
