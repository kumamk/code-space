package heap

import (
	"container/heap"
	"math"
)

//https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

type triplet struct {
	val  int
	lidx int
	nidx int
}

type minheap []triplet

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(triplet))
}

func (h *minheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func SmallestRange(nums [][]int) []int {
	h := &minheap{}
	heap.Init(h)

	mmax := math.MinInt

	for i := range nums {
		heap.Push(h, triplet{val: nums[i][0], lidx: i, nidx: 0})
		mmax = max(mmax, nums[i][0])
	}

	ans := []int{(*h)[0].val, mmax}

	for {
		cnum := heap.Pop(h).(triplet)

		if cnum.nidx == len(nums[cnum.lidx])-1 {
			break
		}

		nnum := nums[cnum.lidx][cnum.nidx+1]
		mmax = max(mmax, nnum)
		heap.Push(h, triplet{val: nnum, lidx: cnum.lidx, nidx: cnum.nidx + 1})

		if mmax-(*h)[0].val < ans[1]-ans[0] {
			ans[0] = (*h)[0].val
			ans[1] = mmax
		}
	}
	return ans
}
