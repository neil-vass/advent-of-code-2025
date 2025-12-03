package main

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestSolvesExample(t *testing.T) {
	s := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	assert.Equal(t, Solve(s), 1227775554)
}

func TestInvalidIDs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		s    string
		want []int
	}{
		{name: "Ends are invalid", s: "11-22", want: []int{11, 22}},
		{name: "Invalid inside range", s: "998-1012", want: []int{1010}},
		{name: "Searches multiple", s: "1188511880-1188511890,222220-222224", want: []int{1188511885, 222222}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvalidIDs(tt.s)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("Contents mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsInvalidID(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 11, want: true},
		{n: 101, want: false},
		{n: 222220, want: false},
		{n: 1188511885, want: true},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := IsInvalidID(tt.n)
			if got != tt.want {
				t.Errorf("IsInvalidID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextInvalidID(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 11, want: 22},
		{n: 998, want: 1010},
		{n: 9999, want: 100100},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := NextInvalidID(tt.n)
			if got != tt.want {
				t.Errorf("NextInvalidID() = %v, want %v", got, tt.want)
			}
		})
	}
}
