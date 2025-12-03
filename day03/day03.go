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
}

func Solve(banks iter.Seq[string], numBats int) int {
	total := 0
	for bank := range banks {
		total += MaxJoltage(bank, numBats)
	}
	return total
}

func MaxJoltage(bank string, numBats int) int {
	var tensVal, unitsVal int

	for i, b := range bank {
		val := int(b - '0')
		if i < len(bank)-1 && val > tensVal {
			tensVal = val
			unitsVal = 0
		} else if val > unitsVal {
			unitsVal = val
		}
	}
	return tensVal*10 + unitsVal
}
