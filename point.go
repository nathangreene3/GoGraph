package graph

import (
	"fmt"
	"math"
)

// Point is an n-tuple.
type Point []float64

// Origin returns the zero-valued point of a given number of dimensions.
func Origin(dims int) Point {
	return make(Point, dims)
}

// String returns the string representation of a point.
func (x Point) String() string {
	return fmt.Sprintf("%0.2f", x)
}

// Add ...
func (x Point) Add(y Point) {
	for i := range x {
		x[i] += y[i]
	}
}

// Subtract ...
func (x Point) Subtract(y Point) {
	for i := range x {
		x[i] -= y[i]
	}
}

// Multiply ...
func (x Point) Multiply(a float64) {
	for i := range x {
		x[i] *= a
	}
}

// Divide ...
func (x Point) Divide(a float64) {
	x.Multiply(1.0 / a)
}

// Distance ...
func (x Point) Distance(y Point) float64 {
	var d, dist float64
	for i := range x {
		d = x[i] - y[i]
		dist += d * d
	}
	return math.Sqrt(dist)
}

// CompareTo returns -1, 0, or 1 indicating if x precedes, is equal to, or follows y.
func (x Point) CompareTo(y Comparable) int {
	if _, ok := y.(Point); !ok {
		panic("not a point")
	}
	m, n := len(x), len(y.(Point))
	for i := 0; i < min(m, n); i++ {
		switch {
		case x[i] < y.(Point)[i]:
			return -1
		case y.(Point)[i] < x[i]:
			return 1
		}
	}

	switch {
	case m < n:
		return -1
	case n < m:
		return 1
	default:
		return 0
	}
}

// Equals returns true if two points are equal in dimension and in value.
func (x Point) Equals(y Comparable) bool {
	if _, ok := y.(Point); !ok {
		panic("not a point")
	}
	return x.CompareTo(y) == 0
}

// Copy returns a copy of a point.
func (x Point) Copy() Duplicatable {
	y := make(Point, len(x))
	copy(y, x)
	return y
}
