package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"slices"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
	"github.com/neil-vass/advent-of-code-2025/shared/set"
)

//go:embed input.txt
var puzzleData string

func main() {
	lines := input.SplitIntoLines(puzzleData)
	fmt.Printf("Part 1: %d\n", SolvePart1(lines))
	fmt.Printf("Part 2: %d\n", SolvePart2(lines))
}

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
	Outside     set.Set[Pos]
}

func (poly Polygon) LineStaysInsidePolygon(line LineSegment) bool {
	for x := min(line.Start.X, line.End.X); x <= max(line.Start.X, line.End.X); x++ {
		for y := min(line.Start.Y, line.End.Y); y <= max(line.Start.Y, line.End.Y); y++ {
			if poly.Outside.Has(Pos{x, y}) {
				return false
			}
		}
	}
	return true
}

func (poly Polygon) BestRectangle() RectArea {
	for _, c := range poly.Candidates {
		corners := []Pos{
			c.Bounds.Min,
			{c.Bounds.Min.X, c.Bounds.Max.Y},
			c.Bounds.Max,
			{c.Bounds.Max.X, c.Bounds.Min.Y},
		}

		// Do edges stay inside?
		if !(poly.LineStaysInsidePolygon(LineSegment{corners[0], corners[1]}) &&
			poly.LineStaysInsidePolygon(LineSegment{corners[1], corners[2]}) &&
			poly.LineStaysInsidePolygon(LineSegment{corners[2], corners[3]}) &&
			poly.LineStaysInsidePolygon(LineSegment{corners[3], corners[0]})) {
			continue
		}

		// Passed all the tests!
		return c
	}

	panic("No suitable rectangles at all")
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

	poly := Polygon{
		Vertices:    append(tiles, tiles[0]),
		BoundingBox: Rect{minBound, maxBound},
		Candidates:  rectAreas,
	}

	// Make clockwise.
	signedArea := 0
	for i := range len(poly.Vertices) - 1 {
		from, to := poly.Vertices[i], poly.Vertices[i+1]
		signedArea += (from.X*to.Y - to.X*from.Y)
	}
	if signedArea > 0 {
		slices.Reverse(poly.Vertices)
	}

	// Mark the "you're stepping outside" border.
	poly.Outside = set.Set[Pos]{}
	for i := range len(poly.Vertices) - 1 {
		edge := LineSegment{poly.Vertices[i], poly.Vertices[i+1]}
		if edge.IsHorizontal() {
			if edge.Start.Y < edge.End.Y {
				x := edge.Start.X - 1
				for y := edge.Start.Y; y <= edge.End.Y; y++ {
					poly.Outside.Add(Pos{x, y})
				}
			} else {
				x := edge.Start.X + 1
				for y := edge.End.Y; y <= edge.Start.Y; y++ {
					poly.Outside.Add(Pos{x, y})
				}
			}
		} else {
			if edge.Start.X < edge.End.X {
				y := edge.Start.Y + 1
				for x := edge.Start.X; x <= edge.End.X; x++ {
					poly.Outside.Add(Pos{x, y})
				}
			} else {
				y := edge.Start.Y - 1
				for x := edge.End.X; x <= edge.Start.X; x++ {
					poly.Outside.Add(Pos{x, y})
				}
			}
		}
	}

	// Remove edges that got mistakenly marked as "outside" above.
	for i := range len(poly.Vertices) - 1 {
		edge := LineSegment{poly.Vertices[i], poly.Vertices[i+1]}
		if edge.IsHorizontal() {
			x := edge.Start.X
			for y := min(edge.Start.Y, edge.End.Y); y <= max(edge.Start.Y, edge.End.Y); y++ {
				delete(poly.Outside, Pos{x, y})
			}
		} else {
			y := edge.Start.Y
			for x := min(edge.Start.X, edge.End.X); x <= max(edge.Start.X, edge.End.X); x++ {
				delete(poly.Outside, Pos{x, y})
			}
		}
	}

	return poly
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
