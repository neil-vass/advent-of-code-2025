package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type MachineDescription struct {
	Lights string
	Wiring [][]int
}

//go:embed input.txt
var puzzleData string

func main() {
	lines := strings.Split(strings.TrimSpace(puzzleData), "\n")
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	panic("unimplemented")
}

var NumbersRe = regexp.MustCompile(`\d+`)
func ParseMachineDescription(s string) MachineDescription {
	var m MachineDescription
	fields := strings.Fields(s)
	m.Lights = strings.Trim(fields[0], "[]")
	m.Wiring = make([][]int, len(fields)-2)
	for i, schematicDescription := range fields[1:len(fields)-1] {
		numbers := NumbersRe.FindAllString(schematicDescription, -1)
		schematic := make([]int, len(numbers))
		for j, lightPos := range numbers {
			n, _ := strconv.Atoi(lightPos)
			schematic[j] = n
		}
		m.Wiring[i] = schematic
	}
	return m
}

func FewestPresses(machineDescription string) int {
	//var m Machine
	return 0
}
