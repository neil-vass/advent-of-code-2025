package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/neil-vass/advent-of-code-2025/shared/assert"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

var example = input.Lines(
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 3)
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, SolvePart2(example), 14)
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		ranges []Range
		want   []Range
	}{
		{name: "Non-overlaps return unchanged", ranges: []Range{{1, 3}, {5, 7}}, want: []Range{{1, 3}, {5, 7}}},
		{name: "Overlaps merge", ranges: []Range{{1, 3}, {2, 7}}, want: []Range{{1, 7}}},
		{name: "Touching ranges merge", ranges: []Range{{1, 3}, {3, 7}}, want: []Range{{1, 7}}},
		{name: "Enclosed ranges vanish", ranges: []Range{{1, 10}, {2, 3}}, want: []Range{{1, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.ranges)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("Contents mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
