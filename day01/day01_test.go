package main

import "testing"

func TestGo(t *testing.T) {
	if 1 != 2 {
		t.Errorf("Our first test failure")
	}
}
