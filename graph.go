package graph

import (
	"fmt"
	"sort"
	"strings"
)

// A Graph is a set of vertices.
type Graph struct {
	vertices []*Vertex
	size     int
}

// NewGraph returns a new graph.
func NewGraph() *Graph {
	return &Graph{vertices: []*Vertex{}}
}

// String returns the string representation of a graph.
func (g *Graph) String() string {
	b := strings.Builder{}
	for i, v := range g.vertices {
		b.WriteString(fmt.Sprintf("Vertex %d\n%s\n", i, v))
	}
	return b.String()
}

// Insert a vertex into a graph.
func (g *Graph) Insert(v *Vertex, bidirected bool) {
	switch {
	case g.size == 0:
		g.vertices = append(g.vertices, v)
		g.size++
	case g.Search(v) == g.size:
		g.vertices = append(g.vertices, v)
		g.size++
		g.Sort()
	}
}

// Remove and return a vertex from a graph.
func (g *Graph) Remove(i int) *Vertex {
	u := g.vertices[i].Copy()
	if i+1 < g.size {
		g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
	} else {
		g.vertices = g.vertices[:i]
	}
	g.size--
	return u
}

// Sort ...
func (g *Graph) Sort() {
	sort.SliceStable(g.vertices, func(i, j int) bool { return g.vertices[i].CompareTo(g.vertices[j]) < 0 })
}

// Search ...
func (g *Graph) Search(v *Vertex) int {
	return sort.Search(g.size, func(i int) bool { return 0 < g.vertices[i].CompareTo(v) })
}
