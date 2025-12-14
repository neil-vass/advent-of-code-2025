package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

type Pos struct{ X, Y int }
type Rect struct{ Min, Max Pos }
type RectArea struct {
	Bounds Rect
	Area   int
}

type Polygon struct {
	Vertices    []Pos
	BoundingBox Rect
	Candidates  []RectArea
}

func (poly Polygon) BestRectangle() RectArea {
	for _, c := range poly.Candidates {

		// For each corner of the rectangle:
		// Use the winding number algorithm, to check it's inside the polygon.
		// If any aren't, we're done!
		// https://stackoverflow.com/questions/217578/how-can-i-determine-whether-a-2d-point-is-within-a-polygon

		// Then, for each edge of the rectangle:
		// Use the instersection test with each edge of the polygon.
		// If they meet, we're OK iff another polygon edge comes right after the intersection.
		// If any aren't OK, we're done!
		//
		// Note: this might be very simple since edges are always horizontal or vertical.
		// We could sort edges into horiz, vert, ordered by the relevant axis value.

		// Not continued yet? We have a winner!
		return c
	}

	panic("No suitable rectangles at all")
}

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	//fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	corners := ParseTiles(lines)
	maxArea := 0
	for i := 0; i < len(corners)-1; i++ {
		for j := i + 1; j < len(corners); j++ {
			area := Area(corners[i], corners[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func SolvePart2(lines []string) int {
	poly := ParsePolygon(lines)
	winner := poly.BestRectangle()
	return winner.Area

}

var tileRe = regexp.MustCompile(`^(\d+),(\d+)$`)

func ParseTiles(lines []string) []Pos {
	tiles := []Pos{}
	for _, ln := range lines {
		var p Pos
		err := input.Parse(tileRe, ln, &p.X, &p.Y)
		if err != nil {
			panic(err) // Don't panic!
		}
		tiles = append(tiles, p)
	}
	return tiles
}

func ParsePolygon(lines []string) Polygon {
	tiles := []Pos{}
	minBound := Pos{math.MaxInt, math.MaxInt}
	maxBound := Pos{math.MinInt, math.MinInt}
	rectAreas := []RectArea{}

	for _, ln := range lines {
		var tile Pos
		err := input.Parse(tileRe, ln, &tile.X, &tile.Y)
		if err != nil {
			panic(err) // Don't panic!
		}

		// Build bounding box of polygon
		minBound.X = min(minBound.X, tile.X)
		minBound.Y = min(minBound.Y, tile.Y)
		maxBound.X = max(maxBound.X, tile.X)
		maxBound.Y = max(maxBound.Y, tile.Y)

		// Build list of rectangles we can choose from
		for _, otherTile := range tiles {
			minPos := Pos{
				X: min(tile.X, otherTile.X),
				Y: min(tile.Y, otherTile.Y),
			}
			maxPos := Pos{
				X: max(tile.X, otherTile.X),
				Y: max(tile.Y, otherTile.Y),
			}
			ra := RectArea{
				Bounds: Rect{minPos, maxPos},
				Area:   Area(tile, otherTile),
			}
			rectAreas = append(rectAreas, ra)
		}
		tiles = append(tiles, tile)
	}

	// Sort by area, greatest first
	slices.SortFunc(rectAreas, func(a, b RectArea) int { return b.Area - a.Area })

	return Polygon{
		Vertices:    append(tiles, tiles[0]),
		BoundingBox: Rect{minBound, maxBound},
		Candidates:  rectAreas,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Area(a, b Pos) int {
	length := a.X - b.X
	if length < 0 {
		length = -length
	}
	length += 1
	height := a.Y - b.Y
	if height < 0 {
		height = -height
	}
	height += 1
	return length * height
}

type LineSegment struct {
	Start, End Pos
}

func (ls LineSegment) IsVertical() bool {
	return ls.Start.Y == ls.End.Y
}

func (ls LineSegment) IsHorizontal() bool {
	return ls.Start.X == ls.End.X
}

func FindIntersect(line1, line2 LineSegment) (Pos, bool) {
	if line1.IsHorizontal() {
		if line2.IsVertical() {
			if between(line2.Start.X, line1.Start.X, line2.End.X) &&
				between(line1.Start.Y, line2.Start.Y, line1.End.Y) {
				return Pos{line1.Start.X, line2.Start.Y}, true
			}
		}
	} else {
		if line2.IsHorizontal() {
			if between(line2.Start.Y, line1.Start.Y, line2.End.Y) &&
				between(line1.Start.X, line2.Start.X, line1.End.X) {
				return Pos{line2.Start.X, line1.Start.Y}, true
			}
		}
	}
	return Pos{}, false
}

func between(a, b, c int) bool {
	a, c = min(a, c), max(a, c)
	return a <= b && b <= c
}
