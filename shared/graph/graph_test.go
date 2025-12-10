package graph

import (
	"testing"
)

type fakeGraph struct {
	neighbours map[string][]NodeCost[string]
	goal       string
}

func (fg fakeGraph) Neighbours(node string) []NodeCost[string] { return fg.neighbours[node] }
func (fg fakeGraph) Heuristic(from string) float64             { return 0 }
func (fg fakeGraph) GoalReached(candidate string) bool         { return candidate == fg.goal }

func TestA_StarSearch(t *testing.T) {

	t.Run("Gets from A to B", func(t *testing.T) {
		g := fakeGraph{
			neighbours: map[string][]NodeCost[string]{
				"A": {{Node: "B", Cost: 1}},
				"B": {},
			},
			goal: "B",
		}

		goalFound, cost := A_StarSearch(g, "A")
		if !goalFound {
			t.Errorf("No route to goal found")
		}
		if cost != 1 {
			t.Errorf("Incorrect cost for route, want 1 got %v", cost)
		}
	})

	t.Run("Takes cheaper route with more nodes", func(t *testing.T) {
		g := fakeGraph{
			neighbours: map[string][]NodeCost[string]{
				"A": {{Node: "B", Cost: 10}, {Node: "C", Cost: 2}},
				"B": {},
				"C": {{Node: "A", Cost: 2}, {Node: "D", Cost: 2}},
				"D": {{Node: "C", Cost: 2}, {Node: "B", Cost: 2}},
			},
			goal: "B",
		}

		goalFound, cost := A_StarSearch(g, "A")
		if !goalFound {
			t.Errorf("No route to goal found")
		}
		want := 6
		if cost != want {
			t.Errorf("Incorrect cost for route, got %v want %v", cost, want)
		}
	})
}

// TODO: Can we demonstrate heuristic's working?
