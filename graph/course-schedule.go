package graph

// https://leetcode.com/problems/course-schedule-ii/description/

// using topological sorting
func FindOrder(numCourses int, prerequisites [][]int) []int {
	adj := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, val := range prerequisites {
		a := val[0]
		b := val[1]

		adj[b] = append(adj[b], a)
		indegree[a]++
	}

	tps := func() []int {
		var result []int
		var que []int
		count := 0

		for i := 0; i < numCourses; i++ {
			if indegree[i] == 0 {
				count++
				que = append(que, i)
				result = append(result, i)
			}
		}

		for len(que) > 0 {
			u := que[0]
			que = que[1:]

			for _, v := range adj[u] {
				indegree[v]--

				if indegree[v] == 0 {
					count++
					que = append(que, v)
					result = append(result, v)
				}
			}
		}

		if count != numCourses {
			return []int{}
		}

		return result
	}

	return tps()
}
