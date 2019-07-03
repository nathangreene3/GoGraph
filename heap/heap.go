package heap

import "github.com/nathangreene3/graph"

// MaxHeap holds comparable data for sorting and prioritizing.
type MaxHeap struct {
	heap      []graph.Comparable
	size, cap int
}

// New returns a maximum heap for prioritizing data.
func New(cap int) *MaxHeap {
	return &MaxHeap{
		heap: make([]graph.Comparable, cap+1),
		cap:  cap,
	}
}

// Sort comparable data.
func Sort(values []graph.Comparable) []graph.Comparable {
	h := New(len(values))
	for i := range values {
		h.Push(values[i])
	}

	return h.sort()
}

// sort returns the data.
func (h *MaxHeap) sort() []graph.Comparable {
	if h.size < 1 {
		return nil
	}

	for 0 < h.size {
		h.Pop()
	}

	return h.heap[1:]
}

// grow the heap by a given capacity.
func (h *MaxHeap) grow(cap int) {
	h.heap = append(h.heap, make([]graph.Comparable, cap)...)
	h.cap += cap
}

// Push a value onto the heap.
func (h *MaxHeap) Push(value graph.Comparable) {
	if h.size < h.cap {
		h.size++
		h.heap[h.size] = value
		h.siftUp(h.size)
	} else {
		if h.cap < 256 {
			h.grow(h.cap)
		} else {
			h.grow(256)
		}

		h.Push(value)
	}
}

// Pop returns the top of the heap.
func (h *MaxHeap) Pop() graph.Comparable {
	if h.size < 1 {
		return nil
	}

	h.swap(1, h.size)
	h.size--
	h.siftDown(1)
	return h.heap[h.size+1]
}

// siftUp corrects the heap from i up.
func (h *MaxHeap) siftUp(i int) {
	if 1 < i {
		j := i / 2 // Parent index
		if h.heap[j].CompareTo(h.heap[i]) < 0 {
			h.swap(i, j)
			h.siftUp(j)
		}
	}
}

// siftDown corrects the heap from i down.
func (h *MaxHeap) siftDown(i int) {
	if i < 1 {
		panic("index out of range")
	}

	j := 2 * i // Left child index
	k := j + 1 // Right child index
	if j <= h.size {
		if k <= h.size && h.heap[j].CompareTo(h.heap[k]) < 0 {
			j = k
		}

		if h.heap[i].CompareTo(h.heap[j]) < 0 {
			h.swap(i, j) // Swap with largest child
			h.siftDown(j)
		}
	}
}

// swap two values on the heap.
func (h *MaxHeap) swap(i, j int) {
	temp := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = temp
}
