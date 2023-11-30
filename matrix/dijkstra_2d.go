package graph

// https://leetcode.com/problems/path-with-minimum-effort/description/

import (
	"container/heap"
	"math"
)

type Cell struct {
	R int
	C int
}

type Node struct {
	V Cell
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

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func MinimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])

	result := make([][]int, m)
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// set all node path to max
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = math.MaxInt
		}
	}

	pque := &PQ{}
	// source node
	heap.Push(pque, Node{V: Cell{R: 0, C: 0}, W: 0})
	result[0][0] = 0

	// out of bound check
	isSafe := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	for pque.Len() > 0 {
		node := heap.Pop(pque).(Node)

		v := node.V
		w := node.W

		if result[v.R][v.C] < w {
			continue // skip already set
		}

		// if popped cell is last return w as min heap keeps the lowest
		if v.R == m-1 && v.C == n-1 {
			return w
		}

		for _, d := range dir {
			nx := v.R + d[0]
			ny := v.C + d[1]

			if isSafe(nx, ny) {
				diff := absDiffInt(heights[v.R][v.C], heights[nx][ny])
				maxDiff := max(diff, w)

				if result[nx][ny] > maxDiff {
					result[nx][ny] = maxDiff
					heap.Push(pque, Node{V: Cell{R: nx, C: ny}, W: maxDiff})
				}
			}
		}
	}
	return result[m-1][n-1]
}
