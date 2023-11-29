package dsu

func DetectCycle(V int, adj map[int][]int) bool {
	parent := make([]int, V)
	rank := make([]int, V)

	for i := 0; i < V; i++ {
		parent[i] = i
		rank[i] = 1
	}

	var find func(int) int
	var union func(int, int)

	// find func
	find = func(x int) int {
		if x == parent[x] {
			return x
		}
		parent[x] = find(parent[x])
		return parent[x]
	}

	// union func
	union = func(x, y int) {
		p_x := find(x)
		p_y := find(y)

		// if belong to same set
		if p_x == p_y {
			return
		}

		if rank[p_x] > rank[p_y] {
			parent[p_y] = p_x
		} else if rank[p_y] > rank[p_x] {
			parent[p_x] = p_y
		} else {
			parent[p_x] = p_y
			rank[p_y]++
		}
	}

	for u := 0; u < V; u++ {
		for _, v := range adj[u] {
			// 0 ->1 only
			if u < v {
				p_u := find(u)
				p_v := find(v)

				if p_u == p_v {
					return true
				}
				union(u, v)
			}
		}
	}
	return false
}
