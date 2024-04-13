package heap

import "errors"

type Heap[V any] struct {
	Arr     []V
	size    int
	compare func(a, b V) int
}

func NewHeap[V any](
	compare func(a, b V) int,
) *Heap[V] {
	return &Heap[V]{
		compare: compare,
	}
}

func (h *Heap[V]) GetArr() []V {
	return h.Arr
}

func (h *Heap[V]) isHigherRank(a, b V) bool {
	return h.compare(a, b) > 0
}

func (h *Heap[V]) Size() int {
	return h.size
}

func (h *Heap[V]) Insert(v V) {
	if len(h.Arr) > h.size {
		h.Arr[h.size] = v
	} else {
		h.Arr = append(h.Arr, v)
	}

	h.size = h.size + 1
	currIndex := h.size - 1

	for currIndex != 0 && h.isHigherRank(h.Arr[currIndex], h.Arr[h.indexParent(currIndex)]) {
		h.Swap(currIndex, h.indexParent(currIndex))
		currIndex = h.indexParent(currIndex)
	}
}

func (h *Heap[V]) Swap(a, b int) {
	h.Arr[a], h.Arr[b] = h.Arr[b], h.Arr[a]
}

func (h *Heap[V]) indexParent(i int) int {
	return (i - 1) / 2
}

func (h *Heap[V]) Float(i int) {
	h.Swap(i, h.indexParent(i))
}

func (h * Heap[V]) isValidIndex(i int) bool {
	return i <= h.LastIndex()
}

func (h *Heap[V]) LastIndex() int {
	return h.size -1
}

func (h *Heap[V]) Pop() (V, error) {
	if h.size == 0 {
		var v V
		return v, EmptyHeapError 
	}
	popped := h.Arr[0]
	h.Arr[0] = h.Arr[h.LastIndex()]
	var v V
	h.Arr[h.LastIndex()] = v
	h.size = h.size - 1

	currIndex := 0

	for !h.isLeafNode(currIndex) {
		leftIndex := h.indexLeft(currIndex)
		rightIndex := h.indexRight(currIndex)
		largerChildIndex := rightIndex
		if !h.isValidIndex(rightIndex) || h.isHigherRank(h.Arr[leftIndex], h.Arr[rightIndex]) {
			largerChildIndex = leftIndex
		}
		if h.isHigherRank(h.Arr[largerChildIndex], h.Arr[currIndex]) {
			h.Swap(currIndex, largerChildIndex)
			currIndex = largerChildIndex
		} else {
			break
		}
	}

	return popped, nil
}

func (h *Heap[V]) isLeafNode(i int) bool {
	return h.indexLeft(i) >= h.size
}

func (h *Heap[V]) indexLeft(i int) int {
	return i*2 + 1
}

func (h *Heap[V]) indexRight(i int) int {
	return i*2 + 2
}

var EmptyHeapError = errors.New("Heap is empty")

func (h *Heap[V]) Peek() (V, error) {
	if h.size == 0 {
		var v V
		return v, EmptyHeapError
	}
	return h.Arr[0], nil
}
