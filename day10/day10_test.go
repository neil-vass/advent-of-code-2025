package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseMachineDescription(t *testing.T) {
	got := ParseMachineDescription("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
	want := MachineDescription{
		Lights:  ".##.",
		Buttons: [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
		Joltage: []int{3, 5, 4, 7},
	}
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestPressForLights(t *testing.T) {
	tests := []struct {
		name   string
		button []int
		lights string
		want   string
	}{
		{name: "Toggle on", button: []int{1}, lights: "...", want: ".#."},
		{name: "Toggle several", button: []int{0, 2}, lights: "##.", want: ".##"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PressForLights(tt.button, tt.lights)
			if got != tt.want {
				t.Errorf("Press() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFewestPressesForLights(t *testing.T) {
	tests := []struct {
		name               string // description of this test case
		machineDescription string
		want               int
	}{
		{"First example", "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", 2},
		{"Second example", "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}", 3},
		{"Third example", "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FewestPressesForLights(tt.machineDescription)
			if got != tt.want {
				t.Errorf("FewestPressesForLights() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Fails, please use day10_python instead.
func TestFewestPressesForJoltage(t *testing.T) {
	tests := []struct {
		name               string // description of this test case
		machineDescription string
		want               int
	}{
		{"First example", "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", 10},
		{"Second example", "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}", 12},
		{"Third example", "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FewestPressesForJoltage(tt.machineDescription)
			if got != tt.want {
				t.Errorf("FewestPressesForJoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}
