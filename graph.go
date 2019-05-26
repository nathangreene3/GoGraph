package main

import "sort"

// A Graph is a set of vertices.
type Graph struct {
	vertices []*Vertex
	size     int
}

// newGraph returns a new graph.
func newGraph(dims int) *Graph {
	return &Graph{vertices: []*Vertex{}}
}

// insert a vertex into a graph.
func (g *Graph) insert(v *Vertex) {
	if g.indexOf(v) < g.size {
		g.vertices = append(g.vertices, v)
		g.sortVertices()
		g.size++
	}
}

// remove and return a vertex from a graph.
func (g *Graph) remove(i int) *Vertex {
	u := g.copyVertex(i)
	if i+1 == g.size {
		g.vertices = g.vertices[:i]
	} else {
		g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
	}

	return u
}

func (g *Graph) sortVertices() {
	sort.Slice(g.vertices, func(i, j int) bool { return CompareVertices(g.vertices[i], g.vertices[j]) < 0 })
}

func (g *Graph) indexOf(v *Vertex) int {
	return sort.Search(g.size, func(i int) bool { return CompareVertices(g.vertices[i], v) <= 0 })
}

func (g *Graph) copyVertex(i int) *Vertex {
	return copyVertex(g.vertices[i])
}
