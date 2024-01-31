package backtracking

//https://leetcode.com/problems/fair-distribution-of-cookies/description/

import "math"

func DistributeCookies(cookies []int, k int) int {
	n := len(cookies)
	child := make([]int, k)
	result := math.MaxInt

	var solve func(int)
	solve = func(idx int) {
		// base case
		if idx >= n {
			ans := 0
			for _, val := range child {
				ans = max(ans, val)
			}
			result = min(result, ans) // update result on each end
			return
		}

		cc := cookies[idx]
		for i := 0; i < k; i++ {
			child[i] += cc // give cookie to ith child
			solve(idx + 1) // exlore
			child[i] -= cc // revert back cookie from ith child
		}
	}
	solve(0)
	return result
}
