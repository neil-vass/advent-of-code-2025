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
	TilesInside set.Set[Pos]
}

func (poly Polygon) FindAllIntersects(ls LineSegment, checker func(LineSegment, LineSegment) (Pos, bool)) []Pos {
	result := []Pos{}
	for i := range len(poly.Vertices) - 1 {
		edge := LineSegment{poly.Vertices[i], poly.Vertices[i+1]}
		if intersect, ok := checker(ls, edge); ok {
			result = append(result, intersect)
		}
	}
	return result
}

func (poly Polygon) PointInPolygon(point Pos) bool {
	ray := LineSegment{
		Start: Pos{poly.BoundingBox.Min.X, point.Y},
		End:   point,
	}
	edgeCrossings := poly.FindAllIntersects(ray, FindIntersectIncludingTouching)

	if len(edgeCrossings)%2 == 0 {
		touchingEdge := (len(edgeCrossings) > 0 && edgeCrossings[len(edgeCrossings)-1] == ray.End)
		if !touchingEdge {
			return false
		}
	}
	return true
}

func (poly Polygon) LineStaysInsidePolygon(line LineSegment) bool {
	edgeCrossings := poly.FindAllIntersects(line, FindIntersect)

	edgeCrossingsWithLineEndsDiscarded := []Pos{}
	for _, pos := range edgeCrossings {
		if pos != line.Start && pos != line.End {
			edgeCrossingsWithLineEndsDiscarded = append(edgeCrossingsWithLineEndsDiscarded, pos)
		}
	}

	for i := range len(edgeCrossingsWithLineEndsDiscarded) - 1 {
		if !(gap(edgeCrossingsWithLineEndsDiscarded[i], edgeCrossingsWithLineEndsDiscarded[i+1]) == 1) {
			return false
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

		// Are corners inside?
		if !(poly.PointInPolygon(corners[0]) &&
			poly.PointInPolygon(corners[1]) &&
			poly.PointInPolygon(corners[2]) &&
			poly.PointInPolygon(corners[3])) {
			continue
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

func gap(a, b Pos) int {
	xDiff := a.X - b.X
	if xDiff < 0 {
		xDiff = -xDiff
	}
	yDiff := a.Y - b.Y
	if yDiff < 0 {
		yDiff = -yDiff
	}
	return xDiff + yDiff
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

	poly := Polygon{
		Vertices:    append(tiles, tiles[0]),
		BoundingBox: Rect{minBound, maxBound},
		Candidates:  rectAreas,
		TilesInside: set.Set[Pos]{},
	}

	// Add edges
	for i := range len(poly.Vertices) - 1 {
		edge := LineSegment{poly.Vertices[i], poly.Vertices[i+1]}
		if edge.IsHorizontal() {
			x := edge.Start.X
			for y := min(edge.Start.Y, edge.End.Y); y <= max(edge.Start.Y, edge.End.Y); y++ {
				poly.TilesInside.Add(Pos{x, y})
			}
		} else {
			y := edge.Start.Y
			for x := min(edge.Start.X, edge.End.X); x <= max(edge.Start.X, edge.End.X); x++ {
				poly.TilesInside.Add(Pos{x, y})
			}
		}
	}

	// Scan. Misses some things.
	for x := minBound.X; x <= maxBound.X; x++ {
		scanline := LineSegment{Pos{x, minBound.Y}, Pos{x, maxBound.Y}}
		boundaries := poly.FindAllIntersects(scanline, FindIntersectIncludingTouching)
		if len(boundaries) > 0 {
			slices.SortFunc(boundaries, func(a, b Pos) int { return a.Y - b.Y })
			changePos := boundaries[0]
			inside := true
			for _, nextChangePos := range boundaries[1:] {
				if inside {
					for y := changePos.Y; y < nextChangePos.Y; y++ {
						poly.TilesInside.Add(Pos{x, y})
					}
				}
				changePos = nextChangePos
				inside = !inside
			}
		}
	}

	//drawRedTiles(poly)
	drawInterior(poly)
	return poly
}

func drawRedTiles(poly Polygon) {
	fmt.Println("Drawing corners")
	lines := make([][]rune, poly.BoundingBox.Max.Y+1)
	for y := range poly.BoundingBox.Max.Y + 1 {
		chars := make([]rune, poly.BoundingBox.Max.X+1)
		for x := range poly.BoundingBox.Max.X + 1 {
			chars[x] = '.'
		}
		lines[y] = chars
	}
	for _, tile := range poly.Vertices {
		lines[tile.Y][tile.X] = '#'
	}
	for _, ln := range lines {
		fmt.Println(string(ln))
	}
}
func drawInterior(poly Polygon) {
	fmt.Println("Drawing interior")
	for y := range poly.BoundingBox.Max.Y + 1 {
		chars := make([]rune, poly.BoundingBox.Max.X+1)
		for x := range poly.BoundingBox.Max.X + 1 {
			if poly.TilesInside.Has(Pos{x, y}) {
				chars[x] = '#'
			} else {
				chars[x] = '.'
			}
		}
		fmt.Println(string(chars))
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
	return a < b && b < c
}

func FindIntersectIncludingTouching(line1, line2 LineSegment) (Pos, bool) {
	if line1.IsHorizontal() {
		if line2.IsVertical() {
			if betweenInclusive(line2.Start.X, line1.Start.X, line2.End.X) &&
				betweenInclusive(line1.Start.Y, line2.Start.Y, line1.End.Y) {
				return Pos{line1.Start.X, line2.Start.Y}, true
			}
		}
	} else {
		if line2.IsHorizontal() {
			if betweenInclusive(line2.Start.Y, line1.Start.Y, line2.End.Y) &&
				betweenInclusive(line1.Start.X, line2.Start.X, line1.End.X) {
				return Pos{line2.Start.X, line1.Start.Y}, true
			}
		}
	}
	return Pos{}, false
}

func betweenInclusive(a, b, c int) bool {
	a, c = min(a, c), max(a, c)
	return a <= b && b <= c
}
