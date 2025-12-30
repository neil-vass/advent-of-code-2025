package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/fifoqueue"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
	"github.com/neil-vass/advent-of-code-2025/shared/set"
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
	for _, ln := range lines[start:] {
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
	programDescription := CreateLpProgram(machineDescription)
	result, err := RunSolver(programDescription)
	if err != nil {
		panic(err) // Remember: Don't panic
	}
	return result
}

func CreateLpProgram(machineDescription string) []string {
	m := ParseMachineDescription(machineDescription)
	programDescription := []string{}

	// We'll use variables named b0, b1, ... bn
	// Where "b0" means "number of times the button at index 0 was pressed".
	variables := make([]string, len(m.Buttons))
	for i := range m.Buttons {
		variables[i] = "b" + strconv.Itoa(i)
	}

	// Objective: minimize total number of button presses: b0 + b1 + ...
	programDescription = append(programDescription, "Minimize")
	programDescription = append(programDescription, strings.Join(variables, " + "))
	programDescription = append(programDescription, "")

	// Constraints: each button press affects some joltages.
	// If, for example, only buttons 0 and 1 affect a joltage that needs to
	// get to 3, we know that b0 + b1 = 3.
	programDescription = append(programDescription, "Subject To")

	// First, note all the buttons that affect each jPos (giving strings like "b0 + b1")
	joltageEffects := make([]string, len(m.Joltage))
	for btnPos, btn := range m.Buttons {
		for _, jPos := range btn {
			if len(joltageEffects[jPos]) == 0 {
				joltageEffects[jPos] = variables[btnPos]
			} else {
				joltageEffects[jPos] += " + " + variables[btnPos]
			}
		}
	}

	// Next, finish each string with its target (giving "b0 + b1 = 3")
	// and add it as a constraint.
	for jPos, constraint := range joltageEffects {
		constraint += " = " + strconv.Itoa(m.Joltage[jPos])
		programDescription = append(programDescription, constraint)
	}
	programDescription = append(programDescription, "")

	// Finally, specify that all variables are general integers.
	// By default they can be any int >= 0, which is fine for us.	
	programDescription = append(programDescription, "General")
	programDescription = append(programDescription, strings.Join(variables, " "))
	programDescription = append(programDescription, "")
	programDescription = append(programDescription, "End")

	return programDescription
}

var statusRe = regexp.MustCompile(`^\s*Status\s+(\w+)\s*$`)
var resultRe = regexp.MustCompile(`^\s*Primal bound\s+(\d+)\s*$`)

func RunSolver(programDescription []string) (int, error) {
	content := []byte(strings.Join(programDescription, "\n"))
	err := os.WriteFile("temp.lp", content, 0644)
	if err != nil {
		return 0, err
	}

	cmd := exec.Command("highs", "--options_file=HiGHS.options", "temp.lp")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	buf := bufio.NewScanner(out)
	cmd.Start()
	defer cmd.Wait()
	for buf.Scan() {

		var status string
		if err := input.Parse(statusRe, buf.Text(), &status); err == nil {
			if status != "Optimal" {
				return 0, errors.New("No optimal solution found")
			}
		}

		var result int
		if err := input.Parse(resultRe, buf.Text(), &result); err == nil {
			return result, nil
		}
	}
	return 0, errors.New("HiGHS output wasn't in expected format")
}
