package input

import (
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestSplitIntoLines(t *testing.T) {
	got := SplitIntoLines("one\ntwo\n")
	want := []string{"one", "two"}
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}

func TestParseSingleInt(t *testing.T) {
	re := regexp.MustCompile(`^(\d+)$`)
	line := "5"
	var x int
	Parse(re, line, &x)
	assert.Equal(t, x, 5)
}

func TestParseMultipleTypes(t *testing.T) {

	type Result struct {
		x int
		y string
		z float64
	}

	re := regexp.MustCompile(`^__(-?\d+)__(\w+)__(-?\d*\.\d+)__$`)
	line := "__-4__test__3.14__"
	want := Result{x: -4, y: "test", z: 3.14}

	var got Result
	err := Parse(re, line, &got.x, &got.y, &got.z)

	assert.Equal(t, err, nil)
	assert.Equal(t, got, want)
}

func TestParseWithTooManyValuesGivesError(t *testing.T) {

	re := regexp.MustCompile(`^one: (\d+), two: (\d+)$`)
	line := "one: 1, two: 2"
	var x, y, z int
	err := Parse(re, line, &x, &y, &z)

	if err == nil || !strings.Contains(err.Error(), "wrong number") {
		t.Errorf("Wanted error about wrong number of values")
	}
}
