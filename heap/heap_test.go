package heap

import (
	"testing"

	"github.com/nathangreene3/graph"
)

// hint is an int that can be put into a heap.
type hint int

func sorted(values []graph.Comparable) bool {
	n := len(values)
	for i := 0; i+1 < n; i++ {
		if 0 < values[i].CompareTo(values[i+1]) {
			return false
		}
	}

	return true
}

func TestHeap(t *testing.T) {
	result := Sort([]graph.Comparable{hint(9), hint(1), hint(8), hint(2), hint(7), hint(3), hint(6), hint(4), hint(5)})
	if !sorted(result) {
		t.Fatalf("Failed to sort: %v", result)
	}
}

func (x hint) CompareTo(y graph.Comparable) int {
	switch {
	case x < y.(hint):
		return -1
	case y.(hint) < x:
		return 1
	default:
		return 0
	}
}

func (x hint) Equals(y graph.Equatable) bool {
	return x.CompareTo(y.(graph.Comparable)) == 0
}

func (x hint) Copy() graph.Duplicatable {
	return x
}
