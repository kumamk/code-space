package main

import (
	recurmemo "code-space/recur_memo"
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
	// nums := []int{1, 2, 3, 4, 3}

	// list := [][]int{{1, 2}, {3, 2}, {5, 3}, {6, 2}, {7, 2}}
	list := [][]int{{3, 2}, {5, 4}, {6, 3}}
	fmt.Println(recurmemo.MaxGifts(list))
}
