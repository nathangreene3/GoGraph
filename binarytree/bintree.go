package binarytree

import "github.com/nathangreene3/graph"

// BinaryTree ...
type BinaryTree struct {
	root *Vertex
}

// Vertex ...
type Vertex struct {
	value               graph.Interface
	parent, left, right *Vertex
}

// New ...
func New(value graph.Interface) *BinaryTree {
	return &BinaryTree{root: NewVertex(value, nil)}
}

// Insert ...
func (b *BinaryTree) Insert(value graph.Interface) {
	b.root.Insert(value)
}

// Remove ...
func (b *BinaryTree) Remove(value graph.Interface) {
	b.root.Remove(value)
}

// Search ...
func (b *BinaryTree) Search(value graph.Interface) *Vertex {
	return b.root.Search(value)
}

// Slice sorts the data in the binary tree into a slice.
func (b *BinaryTree) Slice() []graph.Interface {
	return b.root.Slice()
}

// Slice returns the post-order traversal of a vertex.
func (v *Vertex) Slice() []graph.Interface {
	if v.left != nil {
		if v.right != nil {
			return append(append(v.left.Slice(), v.right.Slice()...), v.value)
		}
		return append(v.left.Slice(), v.value)
	}

	if v.right != nil {
		return append(v.right.Slice(), v.value)
	}

	return []graph.Interface{v.value}
}

// NewVertex ...
func NewVertex(value graph.Interface, parent *Vertex) *Vertex {
	return &Vertex{
		value:  value,
		parent: parent,
	}
}

// Insert ...
func (v *Vertex) Insert(value graph.Interface) {
	if value.CompareTo(v.value) < 0 {
		if v.left == nil {
			v.left = NewVertex(value, v)
		} else {
			v.left.Insert(value)
		}
	} else {
		if v.right == nil {
			v.right = NewVertex(value, v)
		} else {
			v.right.Insert(value)
		}
	}
}

// Remove ...
func (v *Vertex) Remove(value graph.Interface) {
	c := v.value.CompareTo(value)
	switch {
	case c < 0:
		if v.right != nil {
			v.right.Remove(value)
		}
	case 0 < c:
		if v.left != nil {
			v.left.Remove(value)
		}
	default:
		r := v.RightmostDescendent()
		if r.parent.left == r {
			r.parent.left = nil
		} else {
			r.parent.right = nil
		}

		r.parent = v.parent
		r.left = v.left
		r.right = v.right
		if v == v.parent.left {
			v.parent.left = r
		} else {
			v.parent.right = r
		}
	}
}

// RightmostDescendent ...
func (v *Vertex) RightmostDescendent() *Vertex {
	if v.right == nil {
		if v.left == nil {
			return v
		}

		return v.left.RightmostDescendent()
	}

	return v.right.RightmostDescendent()
}

// Search ...
func (v *Vertex) Search(value graph.Interface) *Vertex {
	comparison := v.value.CompareTo(value)
	switch {
	case comparison < 0:
		if v.right != nil {
			return v.right.Search(value)
		}
	case 0 < comparison:
		if v.left != nil {
			return v.left.Search(value)
		}
	default:
		return v
	}

	return nil
}
