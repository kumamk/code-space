package heap

import "container/heap"

type mheap []int

func (h mheap) Len() int {
	return len(h)
}

func (h mheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h mheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *mheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *mheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKthLargest(nums []int, k int) int {
	h := &mheap{}
	heap.Init(h)
	for _, val := range nums {
		heap.Push(h, val)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
