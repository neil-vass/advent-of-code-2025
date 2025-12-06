package main

import (
	_ "embed"
	"fmt"
	"iter"
	"regexp"
	"strconv"
	"strings"
)

type nextFn func() int

//go:embed input.txt
var puzzleData string

func main() {
	lines := strings.Split(strings.TrimSpace(puzzleData), "\n")
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	last := len(lines) - 1
	symbolsStr := lines[last]
	numberRows := []nextFn{}
	for _, ln := range lines[:last] {
		next, _ := iter.Pull(strings.FieldsSeq(ln))
		numberRows = append(numberRows, func() int {
			s, _ := next()
			n, _ := strconv.Atoi(s)
			return n
		})
	}

	total := 0
	for symbol := range strings.FieldsSeq(symbolsStr) {
		totalForThisSum := 0
		accumulate := func(n int) { totalForThisSum += n }
		if symbol == "*" {
			totalForThisSum = 1
			accumulate = func(n int) { totalForThisSum *= n }
		}

		for _, nextVal := range numberRows {
			accumulate(nextVal())
		}
		total += totalForThisSum
	}
	return total
}

func SolvePart2(lines []string) int {
	total := 0
	last := len(lines) - 1
	numberRows, symbolsStr := lines[:last], lines[last]
	re := regexp.MustCompile(`[+*]\s*`)

	for _, sumCols := range re.FindAllStringIndex(symbolsStr, -1) {
		firstCol, lastCol := sumCols[0], sumCols[1]
		if lastCol < len(symbolsStr)-1 {
			lastCol -= 1 // ignore blank col before the next sum
		}
		operator := symbolsStr[firstCol]
		totalForThisSum := 0
		accumulate := func(n int) { totalForThisSum += n }

		if operator == '*' {
			totalForThisSum = 1
			accumulate = func(n int) { totalForThisSum *= n }
		}

		for col := firstCol; col < lastCol; col++ {
			val := []byte{}
			for _, row := range numberRows {
				if row[col] != ' ' {
					val = append(val, row[col])
				}
			}
			intVal, _ := strconv.Atoi(string(val))
			accumulate(intVal)
		}
		total += totalForThisSum
	}
	return total
}
