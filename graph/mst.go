package graph

import "container/heap"

// Prim's Algorithm
func MST(adj map[int][]Node) int {
	visited := make(map[int]bool)
	cost := 0

	pque := &PQ{}
	// push source node
	heap.Push(pque, Node{V: 0, W: 0})

	for pque.Len() > 0 {
		neigh := heap.Pop(pque).(Node)

		w := neigh.W
		node := neigh.V

		// skip if already visited
		if visited[node] {
			continue
		}

		// mark visited once added
		visited[node] = true
		cost += w

		for _, v := range adj[node] {
			// if not visited
			if !visited[v.V] {
				heap.Push(pque, Node{V: v.V, W: v.W})
			}
		}
	}
	return cost
}
