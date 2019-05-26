package main

import "sort"

// A Vertex is a point with connections to other vertices.
type Vertex struct {
	point       Point
	color       Color
	connections []*Vertex
}

// connect adds a vertex as a connection. The color is not assigned. The connection is directional, that is u points to v, but v does not point to u.
func (v *Vertex) connect(u *Vertex) {
	if !v.connected(u) {
		v.connections = append(v.connections, u)
		v.sortConnections()
	}
}

// connected determines if v points to u.
func (v *Vertex) connected(u *Vertex) bool {
	n := len(v.connections)
	if sort.Search(n, func(i int) bool { return CompareVertices(v.connections[i], u) <= 0 }) < n {
		return true
	}

	return false
}

// assignColor sets the color.
func (v *Vertex) assignColor(color Color) {
	v.color = color
}

// CompareVertices returns -1, 0, or 1 indicating if u precedes, is equal to, or follows v.
func CompareVertices(u, v *Vertex) int {
	if u == v {
		return 0
	}

	return ComparePoints(u.point, v.point)
}

// sortConnections sorts the connections.
func (v *Vertex) sortConnections() {
	sort.SliceStable(v.connections, func(i, j int) bool { return CompareVertices(v.connections[i], v.connections[j]) < 0 })
}

// copyVertex returns a copy of a vertex.
func copyVertex(v *Vertex) *Vertex {
	cpy := &Vertex{
		point:       make(Point, len(v.point)),
		color:       v.color,
		connections: make([]*Vertex, len(v.connections)),
	}

	copy(cpy.point, v.point)
	copy(cpy.connections, v.connections)
	return cpy
}
