package priorityqueue

import (
	"testing"

	"github.com/neil-vass/advent-of-code-2025/shared/assert"
)

func TestMinPriorityQueue(t *testing.T) {
	q := New[string]()
	q.Push("A", 5)
	q.Push("B", 10)
	q.Push("C", 2)

	assert.Equal(t, q.Pull(), "C")
	assert.Equal(t, q.IsEmpty(), false)

	q.Push("D", 6)
	q.Push("E", 12)

	assert.Equal(t, q.Pull(), "A")
	assert.Equal(t, q.Pull(), "D")
	assert.Equal(t, q.Pull(), "B")
	assert.Equal(t, q.Pull(), "E")

	assert.Equal(t, q.IsEmpty(), true)
	assert.Equal(t, q.Pull(), "")
}
