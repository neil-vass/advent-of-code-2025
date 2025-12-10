package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/neil-vass/advent-of-code-2025/shared/fifoqueue"
	"github.com/neil-vass/advent-of-code-2025/shared/graph"
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
	//fmt.Printf("Part 1: %d\n", Solve(lines, FewestPressesForLights))
	fmt.Printf("Part 2: %d\n", Solve(lines, FewestPressesForJoltage))
}

func Solve(lines []string, CounterFn func(string) int) int {
	total := 0
	for i, ln := range lines {
		fmt.Printf("Line %d of %d\n", i, len(lines))
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
	reached := Set[string]{}
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

// Serialise joltage to string so we can explore
func save(joltage []int) string {
	j, err := json.Marshal(joltage)
	if err != nil {
		panic(err) // I want to stop and look if this happens...
	}
	return string(j)
}

// Deserialise a saved joltage so we can work with it
func load(s string) []int {
	var joltage []int
	err := json.Unmarshal([]byte(s), &joltage)
	if err != nil {
		panic(err) // I want to stop and look if this happens...
	}
	return joltage
}

func FewestPressesForJoltage(machineDescription string) int {
	m := ParseMachineDescription(machineDescription)
	initialJoltage := save(make([]int, len(m.Joltage)))
	goalFound, presses := graph.A_StarSearch(m, initialJoltage)
	if !goalFound {
		panic("Can't make joltage match")
	}

	return presses
}

func (m MachineDescription) Neighbours(node string) []graph.NodeCost[string] {
	neighbours := make([]graph.NodeCost[string], len(m.Buttons))
	for i, btn := range m.Buttons {
		joltage := load(node)
		PressForJoltage(btn, joltage)
		neighbours[i] = graph.NodeCost[string]{Node: save(joltage), Cost: 1}
	}
	return neighbours
}

func (m MachineDescription) Heuristic(from string) float64 {
	currJoltage := load(from)
	worstCase := 0
	for i, jVal := range currJoltage {
		if jVal > m.Joltage[i] {
			return math.Inf(1)
		}
		diff := m.Joltage[i] - jVal
		if diff > worstCase {
			worstCase = diff
		}
	}
	return float64(worstCase)
}

func (m MachineDescription) GoalReached(candidate string) bool {
	return candidate == save(m.Joltage)
}

func PressForJoltage(button []int, joltage []int) {
	for _, pos := range button {
		joltage[pos]++
	}
}
