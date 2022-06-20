package heap_pattern

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Insert(v int) {
	h.data = append(h.data, v)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap) Extract() int {
	if len(h.data) == 0 {
		return 0
	}
	v := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return v
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) left(i int) int {
	return i*2 + 1
}

func (h *Heap) right(i int) int {
	return i*2 + 2
}

func (h *Heap) heapifyUp(i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	for h.data[i] > h.data[parent] {
		h.swap(i, parent)
		i = parent
		parent = (i - 1) / 2
		if parent < 0 {
			break
		}
	}
}

func (h *Heap) swap(parent int, i int) {
	h.data[parent], h.data[i] = h.data[i], h.data[parent]
}

func (h *Heap) heapifyDown(i int) {
	left := h.left(i)
	right := h.right(i)
	if left >= len(h.data) {
		return
	}
	if right >= len(h.data) {
		right = left
	}
	if h.data[i] > h.data[left] && h.data[i] > h.data[right] {
		return
	}
	if h.data[left] > h.data[right] {
		h.swap(i, left)
		h.heapifyDown(left)
	} else {
		h.swap(i, right)
		h.heapifyDown(right)
	}
}
