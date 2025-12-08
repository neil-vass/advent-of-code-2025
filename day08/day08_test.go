package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

var example = input.Lines(
	"162,817,812",
	"57,618,57",
	"906,360,560",
	"592,479,940",
	"352,342,300",
	"466,668,158",
	"542,29,236",
	"431,825,988",
	"739,650,466",
	"52,470,668",
	"216,146,977",
	"819,987,18",
	"117,168,530",
	"805,96,715",
	"346,949,466",
	"970,615,88",
	"941,993,340",
	"862,61,35",
	"984,92,344",
	"425,690,689",
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1(example, 10), 40)
}

func TestPairsByDistance(t *testing.T) {
	pairs := PairsByDistance(example)
	assert.Equal(t, pairs.Pull(), Pair{Pos{162, 817, 812}, Pos{425, 690, 689}})
	assert.Equal(t, pairs.Pull(), Pair{Pos{162, 817, 812}, Pos{431, 825, 988}})
	assert.Equal(t, pairs.Pull(), Pair{Pos{906, 360, 560}, Pos{805, 96, 715}})
}

func TestConnect(t *testing.T) {
	circuits := Connect(example, 10)
	assert.Equal(t, len(circuits), 4)
}
