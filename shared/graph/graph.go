package graph

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/neil-vass/advent-of-code-2025/shared/priorityqueue"
)

type NodeCost[TNode any] struct {
	Node TNode
	Cost int
}

// Graph objects passed to A* search need these methods.
type GraphWithCosts[TNode any] interface {
	// Nodes you can get to in one step, along with cost to move there.
	Neighbours(node TNode) []NodeCost[TNode]

	// A* needs a heuristic: what is the estimated cost to get from this node
	// to the goal? If your heuristic:
	// - Underestimates costs, then search will find the correct answer
	//     (the more you underestimate, the less efficient search will be).
	// - Estimates perfectly, search will find the correct answer without needing to
	//     explore anything except the best path.
	// - Overestimates costs, then search might not find the correct answer
	//     (can find one path, and ignore others as they look more expensive).
	//
	// You can use A* search with no heuristic (just return 0) to do Dijkstra's algorithm.
	Heuristic(from TNode) float64

	// Confirm whether the given node meets the goal condition.
	GoalReached(candidate TNode) bool
}

func A_StarSearch[TNode any](g GraphWithCosts[TNode], start TNode) (goalFound bool, cost int) {
	frontier := priorityqueue.New[TNode]()
	frontier.Push(start, 0)

	savedStart := save(start)
	visited := map[string]struct {
		costSoFar int
		cameFrom  string
	}{savedStart: {costSoFar: 0}}

	for !frontier.IsEmpty() {
		current := frontier.Pull()
		savedCurr := save(current)
		if g.GoalReached(current) {

			// Debugging
			path := []struct {
				costSoFar int
				state     string
			}{}

			for step := savedCurr; step != ""; step = visited[step].cameFrom {
				x := struct {
					costSoFar int
					state     string
				}{visited[step].costSoFar, step}
				path = append(path, x)
			}

			slices.Reverse(path)

			for _, step := range path {
				fmt.Println(step)
			}

			// End debugging

			goalFound = true
			cost = visited[savedCurr].costSoFar
			return
		}

		for _, n := range g.Neighbours(current) {
			newCost := visited[savedCurr].costSoFar + n.Cost
			savedNeighbour := save(n.Node)
			old, beenHereBefore := visited[savedNeighbour]

			// If we haven't been here before, _or_ if we've found a cheaper way to get here
			if !beenHereBefore || newCost < old.costSoFar {
				priority := float64(newCost) + g.Heuristic(n.Node)
				frontier.Push(n.Node, priority)
				visited[savedNeighbour] = struct {
					costSoFar int
					cameFrom  string
				}{newCost, savedCurr}
			}
		}
	}

	// The end of all our exploring.
	goalFound = false
	return
}

func save[TNode any](node TNode) string {
	bytesNode, err := json.Marshal(node)
	savedNode := string(bytesNode)
	if err != nil {
		panic("Can't save nodes")
	}
	return savedNode
}
