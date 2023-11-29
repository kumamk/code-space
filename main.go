package main

import (
	"code-space/graph"
	"fmt"
)

func main() {
	//nums := []int{5, 3, 2, 1}
	// s1 := "abc"
	// s2 := "def"
	// fmt.Println(string.LongestCommonSubsequence(s1, s2))
	// V := 5
	// graph := [][]int{{2, 3, 4}, {3}, {0, 4}, {0, 1}, {0, 2}}
	// adj := make(map[int][]int)

	// for u := 0; u < V; u++ {
	// 	adj[u] = append(adj[u], graph[u]...)
	// }
	time := [][]int{{1, 2, 1}}
	fmt.Println(graph.NetworkDelayTime(time, 2, 1))
}
