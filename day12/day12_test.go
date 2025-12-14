package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var example = []string{
	"0:",
	"###",
	"##.",
	"##.",
	"",
	"1:",
	"###",
	"##.",
	".##",
	"",
	"2:",
	".##",
	"###",
	"##.",
	"",
	"3:",
	"##.",
	"###",
	"##.",
	"",
	"4:",
	"###",
	"#..",
	"###",
	"",
	"5:",
	"###",
	".#.",
	"###",
	"",
	"4x4: 0 0 0 0 2 0",
	"12x5: 1 0 1 0 2 2",
	"12x5: 1 0 1 0 3 2",
}

func TestParseInput(t *testing.T) {
	got := ParseInput(example)
	want := Model{
		Presents: []Present{{7}, {7}, {7}, {7}, {7}, {7}},
		Trees: []Tree{
			{4, 4, []int{0, 0, 0, 0, 2, 0}},
			{12, 5, []int{1, 0, 1, 0, 2, 2}},
			{12, 5, []int{1, 0, 1, 0, 3, 2}},
		},
	}
	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf("Contents mismatch (-want +got):\n%s", diff)
	}
}
