package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/fifoqueue"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
	"github.com/neil-vass/advent-of-code-2025/shared/set"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

type MachineDescription struct {
	Lights  string
	Buttons [][]int
	Joltage []int
}

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", Solve(lines, FewestPressesForLights))
	fmt.Printf("Part 2: %d\n", Solve(lines, FewestPressesForJoltage))
}

func Solve(lines []string, CounterFn func(string) int) int {
	total := 0
	start := 0
	for i, ln := range lines[start:] {
		fmt.Printf("Line %d of %d\n", i+start, len(lines))
		total += CounterFn(ln)
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

func FewestPressesForLights(machineDescription string) int {
	type Pair struct {
		lights  string
		presses int
	}
	m := ParseMachineDescription(machineDescription)

	initialLights := strings.Repeat(".", len(m.Lights))
	frontier := fifoqueue.New(Pair{initialLights, 0})
	reached := set.Set[string]{}
	reached.Add(initialLights)

	for !frontier.IsEmpty() {
		curr := frontier.Pull()
		for _, button := range m.Buttons {
			lightsAfterPressing := PressForLights(button, curr.lights)
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

func PressForLights(button []int, currentLights string) string {
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

func FewestPressesForJoltage(machineDescription string) int {
	m := ParseMachineDescription(machineDescription)

	// Aim: minimize total number of button presses.
	// This is vector "c" in the standard form of the linear program.
	variablesToMinimize := make([]float64, len(m.Buttons))
	for i := range len(m.Buttons) {
		variablesToMinimize[i] = 1
	}

	// Constraints: each button press affects some joltages.
	// This is matrix "A" in the standard form of the linear program.
	rows, cols := len(m.Joltage), len(m.Buttons)
	data := make([]float64, rows*cols)
	for btnPos, btn := range m.Buttons {
		for _, jPos := range btn {
			idx := (jPos * cols) + btnPos
			data[idx] = 1
		}
	}
	joltagesAffectedByBtns := mat.NewDense(rows, cols, data)

	// Constraint targets: we need to reach these exact joltages.
	// This is vector "b" in the standard form of the linear program.
	targetJoltageResults := make([]float64, len(m.Joltage))
	for i, jolt := range m.Joltage {
		targetJoltageResults[i] = float64(jolt)
	}

	opt, _, err := lp.Simplex(variablesToMinimize, joltagesAffectedByBtns, targetJoltageResults, 0, nil)
	if err != nil {
		panic(err) // If this wasn't a standalone script, remember: don't panic
	}

	return int(opt)
}
