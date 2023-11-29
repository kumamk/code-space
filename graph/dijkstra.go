package graph

import (
	"container/heap"
	"math"
)

/*
https://leetcode.com/problems/network-delay-time/description/
Dijkstra's algo using priority queue (min heap)
*/
type Node struct {
	V int
	W int
}

// priority queue
type PQ []Node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].W < pq[j].W
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PQ) Pop() interface{} {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return val
}

func NetworkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][]Node)
	result := make([]int, n+1)

	// fill adj list
	for _, v := range times {
		u, v, w := v[0], v[1], v[2]
		adj[u] = append(adj[u], Node{V: v, W: w})
	}
	// set all node path to max
	for i := 0; i <= n; i++ {
		result[i] = math.MaxInt
	}

	pque := &PQ{}
	// source node
	heap.Push(pque, Node{V: k, W: 0})
	result[k] = 0

	for pque.Len() > 0 {
		node := heap.Pop(pque).(Node)

		v := node.V
		w := node.W

		if result[v] < w {
			continue // skip already set
		}

		for _, ne := range adj[v] {
			nv := ne.V
			nw := ne.W

			if nw+w < result[nv] {
				result[nv] = nw + w
				heap.Push(pque, Node{V: ne.V, W: nw + w})
			}
		}
	}

	time := math.MinInt
	for i := 1; i <= n; i++ {
		time = max(time, result[i])
	}

	if time == math.MaxInt {
		return -1
	}
	return time
}
