package fifoqueue

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestFifoQueue(t *testing.T) {
	q := New[string]()
	q.Push("A")
	q.Push("B")
	q.Push("C")

	assert.Equal(t, q.Pull(), "A")
	assert.Equal(t, q.IsEmpty(), false)

	q.Push("D")
	assert.Equal(t, q.Pull(), "B")
	assert.Equal(t, q.Pull(), "C")
	assert.Equal(t, q.Pull(), "D")

	assert.Equal(t, q.IsEmpty(), true)
	assert.Equal(t, q.Pull(), "")
}

func TestFifoQueueWithInitialItems(t *testing.T) {
	q := New(1, 2)
	assert.Equal(t, q.Pull(), 1)
	assert.Equal(t, q.Pull(), 2)
	assert.Equal(t, q.IsEmpty(), true)
}
