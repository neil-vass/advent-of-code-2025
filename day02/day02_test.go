package main

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestSolvesExample(t *testing.T) {
	s := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	assert.Equal(t, Solve(s, IsInvalidID_Part1), 1227775554)
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
			got := InvalidIDs(tt.s, IsInvalidID_Part1)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("Contents mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsInvalidID_Part1(t *testing.T) {
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
			got := IsInvalidID_Part1(tt.n)
			if got != tt.want {
				t.Errorf("IsInvalidID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInvalidID_Part2(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 12341234, want: true},
		{n: 123123123, want: true},
		{n: 1212121212, want: true},
		{n: 1111111, want: true},
		{n: 1111112, want: false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := IsInvalidID_Part2(tt.n)
			if got != tt.want {
				t.Errorf("IsInvalidID() = %v, want %v", got, tt.want)
			}
		})
	}
}
