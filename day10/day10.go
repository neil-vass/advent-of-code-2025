package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/fifoqueue"
)

type MachineDescription struct {
	Lights  string
	Buttons [][]int
	Joltage []int
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

//go:embed input.txt
var puzzleData string

func main() {
	lines := strings.Split(strings.TrimSpace(puzzleData), "\n")
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	total := 0
	for _, ln := range lines {
		total += FewestPresses(ln)
	}
	return total
}

var NumbersRe = regexp.MustCompile(`\d+`)

func ParseMachineDescription(s string) MachineDescription {
	var m MachineDescription
	fields := strings.Fields(s)
	m.Lights = strings.Trim(fields[0], "[]")

	m.Buttons = make([][]int, len(fields)-2)
	for i, schematic := range fields[1 : len(fields)-1] {
		numbers := NumbersRe.FindAllString(schematic, -1)
		button := make([]int, len(numbers))
		for j, lightPos := range numbers {
			n, _ := strconv.Atoi(lightPos)
			button[j] = n
		}
		m.Buttons[i] = button
	}

	joltageVals := NumbersRe.FindAllString(fields[len(fields)-1], -1)
	m.Joltage = make([]int, len(joltageVals))
	for i, val := range joltageVals {
		n, _ := strconv.Atoi(val)
		m.Joltage[i] = n
	}
	return m
}

func FewestPresses(machineDescription string) int {
	type Pair struct {
		lights  string
		presses int
	}
	m := ParseMachineDescription(machineDescription)

	initialLights := strings.Repeat(".", len(m.Lights))
	frontier := fifoqueue.New(Pair{initialLights, 0})
	reached := Set[string]{}
	reached.Add(initialLights)

	for !frontier.IsEmpty() {
		curr := frontier.Pull()
		for _, button := range m.Buttons {
			lightsAfterPressing := Press(button, curr.lights)
			presses := curr.presses + 1
			if lightsAfterPressing == m.Lights {
				return presses
			}
			if !reached.Has(lightsAfterPressing) {
				frontier.Push(Pair{lightsAfterPressing, presses})
				reached.Add(lightsAfterPressing)
			}
		}
	}

	panic("Can't make lights match")
}

func Press(button []int, currentLights string) string {
	lightsAfterPressing := []byte(currentLights)
	for _, pos := range button {
		if lightsAfterPressing[pos] == '.' {
			lightsAfterPressing[pos] = '#'
		} else {
			lightsAfterPressing[pos] = '.'
		}
	}
	return string(lightsAfterPressing)
}
