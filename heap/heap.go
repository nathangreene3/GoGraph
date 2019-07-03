package heap

import "github.com/nathangreene3/graph"

// MaxHeap ...
type MaxHeap struct {
	heap      []graph.Interface
	size, cap int
}

// New ...
func New(cap int) *MaxHeap {
	return &MaxHeap{
		heap: make([]graph.Interface, cap+1),
		cap:  cap + 1,
	}
}

// Sort ...
func Sort(values []graph.Interface) []graph.Interface {
	h := New(len(values))
	for i := range values {
		h.Push(values[i])
	}
	for range values {
		h.Pop()
	}
	return h.heap
}

// Push ...
func (h *MaxHeap) Push(value graph.Interface) {
	if h.size < h.cap {
		h.size++
		h.heap[h.size] = value
		h.siftUp(h.size)
	} else {
		h.heap = append(h.heap, make([]graph.Interface, h.cap-1)...)
		h.cap = len(h.heap)
		h.Push(value)
	}
}

// Pop ...
func (h *MaxHeap) Pop() graph.Interface {
	value := h.heap[1]
	h.swap(1, h.size)
	h.size--
	h.siftDown(1)
	return value
}

// siftUp ...
func (h *MaxHeap) siftUp(i int) {
	if 1 < i {
		j := parent(i)
		if h.heap[j].CompareTo(h.heap[i]) < 0 {
			temp := h.heap[i]
			h.heap[i] = h.heap[j]
			h.heap[j] = temp
			h.siftUp(j)
		}
	}
}

// siftDown ...
func (h *MaxHeap) siftDown(i int) {
	j := left(i)
	if j <= h.size {
		if h.heap[i].CompareTo(h.heap[j]) < 0 {
			h.swap(i, j)
			h.siftDown(j)
		} else {
			j++
			if j <= h.size && h.heap[i].CompareTo(h.heap[j]) < 0 {
				h.swap(i, j)
				h.siftDown(j)
			}
		}
	}
}

// swap ...
func (h *MaxHeap) swap(i, j int) {
	temp := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = temp
}

// parent ...
func parent(i int) int { return i / 2 }

// left ...
func left(i int) int { return 2 * i }

// right ...
func right(i int) int { return 2*i + 1 }
