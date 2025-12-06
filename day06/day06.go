package main

import (
	_ "embed"
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type nextFn func() int

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines iter.Seq[string]) int {
	rows := []nextFn{}

	var lastStr string
	for ln := range lines {
		if lastStr != "" {
			next, _ := iter.Pull(strings.FieldsSeq(lastStr))
			rows = append(rows, func() int {
				s, _ := next()
				n, _ := strconv.Atoi(s)
				return n
			})
		}
		lastStr = ln
	}

	total := 0
	for symbol := range strings.FieldsSeq(lastStr) {
		if symbol == "+" {
			acc := 0
			for _, nextVal := range rows {
				acc += nextVal()
			}
			total += acc
		} else {
			acc := 1
			for _, nextVal := range rows {
				acc *= nextVal()
			}
			total += acc
		}
	}
	return total
}
