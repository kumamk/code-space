package heap

import "container/heap"

//https://leetcode.com/problems/task-scheduler/

// using max heap
func LeastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	h := &mheap{}
	heap.Init(h)

	// count freq of tasks
	for _, ch := range tasks {
		freq[ch-'A']++
	}

	// push non empty task to heap
	for _, val := range freq {
		if val > 0 {
			heap.Push(h, val)
		}
	}

	time := 0

	for h.Len() > 0 {
		var temp []int
		// add n+1 tasks
		for i := 1; i <= n+1; i++ {
			if h.Len() > 0 {
				t := heap.Pop(h).(int)
				t--
				temp = append(temp, t)
			}
		}
		// readd to heap if freq > 0
		for _, v := range temp {
			if v > 0 {
				heap.Push(h, v)
			}
		}
		// update time after n+1 iteration
		if h.Len() == 0 {
			time += len(temp)
		} else {
			time = time + (n + 1)
		}
	}
	return time
}
