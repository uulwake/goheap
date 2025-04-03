package goheap

type Heap[T any] struct {
	queue []T
	comp  func(T, T) bool
}

func NewHeap[T any](comp func(T, T) bool, initial T) *Heap[T] {
	return &Heap[T]{
		queue: []T{initial},
		comp:  comp,
	}
}

func (h *Heap[T]) Peek() (T, bool) {
	if len(h.queue) <= 1 {
		return h.queue[0], false
	}

	return h.queue[1], true
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.queue) <= 1
}

func (h *Heap[T]) Size() int {
	return len(h.queue) - 1
}

func (h *Heap[T]) shouldSwap(i, j int) bool {
	return !h.comp(h.queue[i], h.queue[j])
}

func (h *Heap[T]) swap(i, j int) {
	h.queue[i], h.queue[j] = h.queue[j], h.queue[i]
}

func (h *Heap[T]) Enqueue(data T) {
	h.queue = append(h.queue, data)

	idx := len(h.queue) - 1
	parentIdx := idx / 2

	for parentIdx > 0 && h.shouldSwap(parentIdx, idx) {
		h.swap(parentIdx, idx)

		idx = parentIdx
		parentIdx = idx / 2
	}
}

func (h *Heap[T]) Dequeue() (T, bool) {
	if len(h.queue) <= 1 {
		return h.queue[0], false
	}

	el := h.queue[1]
	last := h.queue[len(h.queue)-1]
	h.queue = h.queue[:len(h.queue)-1]

	if len(h.queue) == 1 {
		return el, true
	}

	h.queue[1] = last
	parentIdx := 1
	leftIdx := parentIdx * 2
	rightIdx := leftIdx + 1

	for leftIdx < len(h.queue) && h.shouldSwap(parentIdx, leftIdx) || rightIdx < len(h.queue) && h.shouldSwap(parentIdx, rightIdx) {
		if rightIdx >= len(h.queue) {
			h.swap(parentIdx, leftIdx)
			parentIdx = leftIdx
		} else {
			isLeft := h.comp(h.queue[leftIdx], h.queue[rightIdx])

			if isLeft {
				h.swap(parentIdx, leftIdx)
				parentIdx = leftIdx
			} else {
				h.swap(parentIdx, rightIdx)
				parentIdx = rightIdx
			}
		}

		leftIdx = parentIdx * 2
		rightIdx = leftIdx + 1
	}

	return el, true
}
