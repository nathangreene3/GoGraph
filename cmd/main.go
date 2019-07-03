package cmd

import (
	"fmt"

	"github.com/nathangreene3/graph"
	"github.com/nathangreene3/graph/heap"
)

func main() {
	testHeap()
}

type hint struct {
	x int
}

func testHeap() {
	vals := []*hint{
		&hint{9},
		&hint{1},
		&hint{8},
		&hint{2},
		&hint{7},
		&hint{3},
		&hint{6},
		&hint{4},
		&hint{5},
	}
	fmt.Println(heap.Sort(vals))
}

func (x *hint) CompareTo(y graph.Comparable) int {
	switch {
	case x.x < y.(*hint).x:
		return -1
	case x.x > y.(*hint).x:
		return -1
	default:
		return 0
	}
}

func (x *hint) Equals(y graph.Comparable) bool { return x.CompareTo(y) == 0 }

func (x *hint) Copy() graph.Duplicatable { return &hint{x.x} }
