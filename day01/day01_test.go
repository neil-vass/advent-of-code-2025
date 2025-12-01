package main

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/input"
)

func TestSafeSetup(t *testing.T) {
	safe := Safe{dial: 50}
	if safe.dial != 50 || safe.stoppedAtZero != 0 {
		t.Errorf("Safe setup failed")
	}
}

func TestSafe_Turn(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		initial int
		dist    int
		want    int
	}{
		{name: "Turns right", initial: 11, dist: 8, want: 19},
		{name: "Turns left", initial: 19, dist: -19, want: 0},
		{name: "Wraps left", initial: 0, dist: -1, want: 99},
		{name: "Wraps right", initial: 99, dist: 1, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Safe{dial: tt.initial}
			s.Turn(tt.dist)
			if s.dial != tt.want {
				t.Errorf("Dial position=%d, want %d", s.dial, tt.want)
			}
		})
	}
}

func TestSafe_FollowInstructions(t *testing.T) {
	s := Safe{dial: 5}
	instructions := input.Lines("L10", "R5")
	s.Follow(instructions)

	want := 0
	if s.dial != want {
		t.Errorf("got %d, want %d", s.dial, want)
	}
}

func TestSafe_CountsZeros(t *testing.T) {
	s := Safe{dial: 50}
	instructions := input.Lines(
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82")
	s.Follow(instructions)

	want := Safe{
		dial:          32,
		stoppedAtZero: 3,
		passedZero:    6,
	}
	if s != want {
		t.Errorf("Safe: got %#v, want %#v", s, want)
	}
}

func TestSafe_CountsCompleteRotations(t *testing.T) {
	s := Safe{dial: 50}
	s.Turn(1000)

	want := Safe{
		dial:          50,
		stoppedAtZero: 0,
		passedZero:    10,
	}
	if s != want {
		t.Errorf("Safe: got %#v, want %#v", s, want)
	}
}

