package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInvalidIDs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		s    string
		want []int
	}{
		{name: "Ends are invalid", s: "11-22", want: []int{11, 22}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvalidIDs(tt.s)
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("Contents mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
