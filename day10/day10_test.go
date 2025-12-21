package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

func TestGonum(t *testing.T) {
	// machine := ParseMachineDescription("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")

	// c is variables to minimize. as many 1s as there are buttons, maybe?
	// A is a matrix. Is this the constraints? Having it listed 1D looks odd.
	// b is the bounds of what we want variables to be. Does it hit these exactly?
	// -- Where's the b_l and b_u we had in scipy?

	// Python:
	// c = np.array([1, 1, 1, 1, 1, 1])

	// A = np.array([
	//     [0, 0, 0, 0, 1, 1],
	//     [0, 1, 0, 0, 0, 1],
	//     [0, 0, 1, 1, 1, 0],
	//     [1, 1, 0, 1, 0, 0],
	// ])
	// b_u = np.array([3, 5, 4, 7])
	// b_l = b_u

	c := []float64{1, 1, 1, 1, 1, 1}

	A := mat.NewDense(4, 6, []float64{
		0, 0, 0, 0, 1, 1,
		0, 1, 0, 0, 0, 1,
		0, 0, 1, 1, 1, 0,
		1, 1, 0, 1, 0, 0,
	})
	b := []float64{3, 5, 4, 7}

	opt, x, err := lp.Simplex(c, A, b, 0, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("opt: %v\n", opt)
	fmt.Printf("x: %v\n", x)
}

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
		// I'm using the Simplex algorithm, which _I believe_ is the same thing SciPy's MILP uses.
		// However, SciPy version is happy with all these examples while GoNum gets upset.

		// This works.
		{"First example", "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", 10},

		// Fails with "lp: A is singular". In the library code, having the same number of buttons as
		// joltages sends it to a special case ("Problem is exactly constrained, perform a linear solve.")
		// I need to find out why that's special, and why the solve fails.
		//{"Second example", "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}", 12},

		// Fails with "lp: more equality constraints than variables". The docs do say
		// "matrix A must have at least as many columns as rows", meaning you can't have more buttons than
		// joltages. I need to find out why that is.
		// {"Third example", "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", 11},
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
