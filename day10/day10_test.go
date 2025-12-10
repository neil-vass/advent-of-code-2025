package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseMachineDescription(t *testing.T) {
	got := ParseMachineDescription("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
	want := MachineDescription{
		Lights: ".##.",
		Wiring: [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
	}
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestFewestPresses(t *testing.T) {
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
			got := FewestPresses(tt.machineDescription)
			if got != tt.want {
				t.Errorf("FewestPresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
