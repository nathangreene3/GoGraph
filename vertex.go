package graph

import (
	"fmt"
	"sort"
	"strings"
)

// Interface gives functionality to vertices.
type Interface interface {
	Comparable
	Duplicatable
}

// Comparable defines the context for equality.
type Comparable interface {
	CompareTo(x Comparable) int
	Equals(x Comparable) bool
}

// Duplicatable defines the context for copying.
type Duplicatable interface {
	Copy() Duplicatable
}

// A Vertex is a value with connections to other vertices.
type Vertex struct {
	value               Interface
	in, out             []*Vertex
	inDegree, outDegree int
}

// NewVertex returns a vertex defined by a value.
func NewVertex(value Interface) *Vertex {
	return &Vertex{
		value: value,
		in:    []*Vertex{},
		out:   []*Vertex{},
	}
}

// String returns the string representation of a vertex.
func (v *Vertex) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("Value: %v\n", v.value))
	for j, u := range v.out {
		b.WriteString(fmt.Sprintf("Connection %d: %v\n", j, u.value))
	}
	return b.String()
}

// Connect adds a vertex as a connection.
func (v *Vertex) Connect(u *Vertex) {
	if _, ok := v.PointsTo(u); !ok {
		v.out = append(v.out, u)
		u.in = append(u.in, v)
		v.Sort()
		u.Sort()
	}
}

// PointsTo determines if v points to u.
func (v *Vertex) PointsTo(u *Vertex) (int, bool) {
	switch v.outDegree {
	case 0:
		return 0, false
	default:
		index := v.SearchOut(u.value)
		return index, index < v.outDegree
	}
}

// ComesFrom determines if v comes from u.
func (v *Vertex) ComesFrom(u *Vertex) (int, bool) {
	switch v.inDegree {
	case 0:
		return 0, false
	default:
		index := v.SearchIn(u.value)
		return index, index < v.inDegree
	}
}

// Disconnect removes a connected vertex and returns it.
func (v *Vertex) Disconnect(i int) Interface {
	value := v.out[i].value
	if i+1 < v.outDegree {
		v.out = append(v.out[:i], v.out[i+1:]...)
	} else {
		v.out = v.out[:i]
	}
	v.outDegree--
	return value
}

// CompareTo returns -1, 0, or 1 indicating if v precedes, is equal to, or follows u.
func (v *Vertex) CompareTo(u *Vertex) int {
	if v == u {
		return 0
	}
	return v.value.CompareTo(u.value)
}

// Equals returns true if v and u have equal values.
func (v *Vertex) Equals(u *Vertex) bool {
	return v.CompareTo(u) == 0
}

// SearchIn ...
func (v *Vertex) SearchIn(x Interface) int {
	return sort.Search(v.inDegree, func(i int) bool { return 0 < v.in[i].value.CompareTo(x) })
}

// SearchOut a vertex's out connections for a value. Returns an index on the range [0,len) if found and len if not found.
func (v *Vertex) SearchOut(x Interface) int {
	return sort.Search(v.outDegree, func(i int) bool { return 0 < v.out[i].value.CompareTo(x) })
}

// Sort sorts the connections.
func (v *Vertex) Sort() {
	sort.SliceStable(v.in, func(i, j int) bool { return v.in[i].CompareTo(v.in[j]) < 0 })
	sort.SliceStable(v.out, func(i, j int) bool { return v.out[i].CompareTo(v.out[j]) < 0 })
}

// Copy returns a copy of a vertex.
func (v *Vertex) Copy() *Vertex {
	u := &Vertex{
		value:     v.value.Copy().(Interface),
		in:        make([]*Vertex, v.inDegree),
		out:       make([]*Vertex, v.outDegree),
		inDegree:  v.inDegree,
		outDegree: v.outDegree,
	}
	copy(u.in, v.in)
	copy(u.out, v.out)
	return u
}
