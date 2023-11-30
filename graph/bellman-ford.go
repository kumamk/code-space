package graph

import "math"

func Bellmanfor(V int, edges [][]int, source int) []int {
	dist := make([]int, V)

	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt
	}
	dist[source] = 0

	// relax all edges v-1 times
	for i := 1; i <= V-1; i++ {
		for _, edge := range edges {
			u, v, d := edge[0], edge[1], edge[2]

			if dist[u] != math.MaxInt && dist[u]+d < dist[v] {
				dist[v] = dist[u] + d
			}
		}
	}

	cycle := false

	// detect -ve cylce if present
	for _, edge := range edges {
		u, v, d := edge[0], edge[1], edge[2]
		if dist[u] != math.MaxInt && dist[u]+d < dist[v] {
			cycle = true
			break
		}

	}

	if cycle {
		return []int{-1}
	}
	return dist
}
