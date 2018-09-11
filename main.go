package main

import (
	"fmt"
	"math"
	"sort"
)

// A graph is a set of vertices with their edges.
type graph struct {
	vertices []vertex
	edges    []edge
}

// A vertex is an identifier with an n-tuple value
type vertex struct {
	id    int
	value []float64
}

// An edge is a pair of vertex ids.
type edge struct {
	left, right int
}

func main() {
	G := graph{
		vertices: []vertex{
			vertex{
				id:    0,
				value: []float64{0, 0},
			},
			vertex{
				id:    1,
				value: []float64{2, 2},
			},
			vertex{
				id:    2,
				value: []float64{3, 1},
			},
			vertex{
				id:    3,
				value: []float64{4, 2},
			},
		},
		edges: []edge{
			edge{0, 1},
			edge{0, 2},
			edge{0, 3},
			edge{1, 2},
			edge{1, 3},
			edge{2, 3},
		},
	}
	// A := distMatrix(G)
	// fmt.Println(A)
	Vs := []vertex{
		vertex{
			id:    0,
			value: []float64{0, 0},
		},
		vertex{
			id:    1,
			value: []float64{2, 2},
		},
		vertex{
			id:    2,
			value: []float64{3, 1},
		},
		vertex{
			id:    3,
			value: []float64{4, 2},
		},
	}
	H := makeCompleteGraph(Vs)
	fmt.Println(G)
	fmt.Println(H)
	sort.Sort(H)
	fmt.Println(H)
}

// func getMinSpanTree(G graph) (MST graph){
// 	MST=graph{
// 		vertices:copyVertices(G.vertices)
// 	}
// 	return MST
// }

// func copyVertices(G graph) (Vs []vertex) {
// 	n := len(G.vertices)
// 	Vs = make([]vertex, n)
// 	for i := 0; i < n; i++ {
// 		Vs[i] = copyVertex(G.vertices[i])
// 	}
// }

func (G graph) Less(i, j int) (less bool) {
	if dist(G, G.edges[i]) < dist(G, G.edges[j]) {
		less = true
	}
	return less
}

func (G graph) Swap(i, j int) {
	G.edges[i], G.edges[j] = G.edges[j], G.edges[i]
}

func (G graph) Len() int {
	return len(G.edges)
}

// makeCompleteGraph returns a complete graph consisting of the set of vertices.
func makeCompleteGraph(Vs []vertex) (G graph) {
	n := len(Vs)
	G.vertices = make([]vertex, n)
	G.edges = make([]edge, n*(n-1)/2)
	for i := 0; i < n; i++ {
		G.vertices[i] = vertex{
			id:    Vs[i].id,
			value: Vs[i].value,
		}
	}
	k := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			G.edges[k] = edge{G.vertices[i].id, G.vertices[j].id}
			k++
		}
	}
	return G
}

// dist returns the Euclidean distance from one vertex to another in a graph given their edge.
func dist(G graph, E edge) (d float64) {
	for i := 0; i < len(G.vertices[0].value); i++ {
		d += math.Pow(G.vertices[E.left].value[i]-G.vertices[E.right].value[i], 2)
	}
	d = math.Sqrt(d)
	return d
}

// distMatrix returns a matrix of distances from each vertext to each other vertex in a graph.
func distMatrix(G graph) (A [][]float64) {
	n := len(G.vertices)
	A = make([][]float64, n)
	for i := 0; i < len(G.vertices); i++ {
		A[i] = make([]float64, n)
	}
	for i := 0; i < len(G.edges); i++ {
		v := dist(G, G.edges[i])
		A[G.edges[i].left][G.edges[i].right] = v
		A[G.edges[i].right][G.edges[i].left] = v
	}
	return A
}
