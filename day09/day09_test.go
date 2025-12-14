package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

var square = []string{"0,0", "0,1", "1,1", "1,0"}

var LShape = []string{"0,0", "0,2", "3,2", "3,5", "5,5", "5,0"}

var example = []string{
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example), 50)
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		polygonDescription []string
		want               int
	}{
		{name: "Square", polygonDescription: square, want: 4},
		{name: "L Shape", polygonDescription: LShape, want: 18},
		{name: "Example from description", polygonDescription: example, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolvePart2(tt.polygonDescription)

			if got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}

	assert.Equal(t, SolvePart2(square), 4)
}

func TestPointInPolygon(t *testing.T) {
	poly := ParsePolygon(LShape)
	assert.Equal(t, poly.PointInPolygon(Pos{0, 0}), true)
	assert.Equal(t, poly.PointInPolygon(Pos{0, 5}), false)
}

func TestArea(t *testing.T) {
	a, b := Pos{2, 5}, Pos{9, 7}
	assert.Equal(t, Area(a, b), 24)
}

func TestFindIntersect(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line1    LineSegment
		line2    LineSegment
		wantPos  Pos
		wantBool bool
	}{
		{name: "Crossing lines", line1: LineSegment{Pos{2, 0}, Pos{2, 5}}, line2: LineSegment{Pos{0, 3}, Pos{6, 3}}, wantPos: Pos{2, 3}, wantBool: true},
		{name: "Missing lines", line1: LineSegment{Pos{2, 0}, Pos{2, 5}}, line2: LineSegment{Pos{3, 3}, Pos{9, 3}}, wantPos: Pos{}, wantBool: false},
		{name: "Colinear lines", line1: LineSegment{Pos{2, 0}, Pos{2, 5}}, line2: LineSegment{Pos{2, 0}, Pos{2, 10}}, wantPos: Pos{}, wantBool: false},
		{name: "Touching line ends doesn't count as intersecting", line1: LineSegment{Pos{2, 4}, Pos{6, 4}}, line2: LineSegment{Pos{2, 4}, Pos{2, 10}}, wantPos: Pos{}, wantBool: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPos, gotBool := FindIntersect(tt.line1, tt.line2)

			if gotPos != tt.wantPos {
				t.Errorf("FindIntersect() = %v, want %v", gotPos, tt.wantPos)
			}
			if gotBool != tt.wantBool {
				t.Errorf("FindIntersect() = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}
