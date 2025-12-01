package assert

import "testing"

func Equal[V comparable](t *testing.T, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
