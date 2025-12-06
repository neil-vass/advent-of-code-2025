package main

import (
	_ "embed"
	"fmt"
	"iter"
	"strconv"
	"strings"
)

type nextFn func() int

//go:embed input.txt
var puzzleData string

func main() {
	splitOnNewline := func(r rune) bool { return r == '\n' }
	lines := strings.FieldsFunc(puzzleData, splitOnNewline)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	numberRows := []nextFn{}
	var lastStr string
	for _, ln := range lines {
		if lastStr != "" {
			next, _ := iter.Pull(strings.FieldsSeq(lastStr))
			numberRows = append(numberRows, func() int {
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
			for _, nextVal := range numberRows {
				acc += nextVal()
			}
			total += acc
		} else {
			acc := 1
			for _, nextVal := range numberRows {
				acc *= nextVal()
			}
			total += acc
		}
	}
	return total
}

func SolvePart2(lines []string) int {
	return 9
}
