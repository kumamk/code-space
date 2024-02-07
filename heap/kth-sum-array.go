package heap

//https://leetcode.com/problems/find-the-k-sum-of-an-array/

import (
	"container/heap"
	"sort"
)

type Pair struct {
	sum int
	idx int
}

type sumheap []Pair

func (h sumheap) Len() int {
	return len(h)
}

func (h sumheap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h sumheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *sumheap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *sumheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KSum(nums []int, k int) int64 {
	maxSum := 0
	n := len(nums)

	// update max sum and -ve nums to +ve
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
		} else {
			maxSum += nums[i]
		}
	}

	// sort nums in asc order
	sort.Ints(nums)
	// topk list to add smallest element
	topk := make([]int, 0)
	// add 0 init
	topk = append(topk, 0)

	h := &sumheap{}
	heap.Init(h)
	heap.Push(h, Pair{sum: nums[0], idx: 1})

	// until topk has k element
	for len(topk) < k {
		pair := heap.Pop(h).(Pair)

		topk = append(topk, pair.sum)
		// add to heap for both cases
		// 1. by taking current ele
		// 2. by taking current and removing prev ele
		if pair.idx < n {
			heap.Push(h, Pair{sum: pair.sum + nums[pair.idx], idx: pair.idx + 1})
			heap.Push(h, Pair{sum: pair.sum - nums[pair.idx-1] + nums[pair.idx], idx: pair.idx + 1})
		}
	}
	return int64(maxSum - topk[k-1])
}
