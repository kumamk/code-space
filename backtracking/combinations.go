package backtracking

// https://leetcode.com/problems/combinations/description/

// Approach-1 without loop
func Combine(n int, k int) [][]int {
	var result [][]int

	var solve func(int, int, []int)
	solve = func(start, limit int, ans []int) {
		// base case
		if limit == 0 {
			result = append(result, append([]int{}, ans...))
			return
		}

		if start > n {
			return
		}
		// take and explore
		ans = append(ans, start)
		solve(start+1, limit-1, ans)
		// not take and explore
		ans = ans[:len(ans)-1]
		solve(start+1, limit, ans)
	}
	ans := make([]int, 0)
	solve(1, k, ans)
	return result
}

// Approach-2 with loop
func combineViaLoop(n int, k int) [][]int {
	var result [][]int

	var solve func(int, int, []int)
	solve = func(start, limit int, ans []int) {
		// base case
		if limit == 0 {
			result = append(result, append([]int{}, ans...))
			return
		}

		for i := start; i <= n; i++ {
			ans = append(ans, i)     // take
			solve(i+1, limit-1, ans) // explore
			ans = ans[:len(ans)-1]   // not take
		}
	}
	ans := make([]int, 0)
	solve(1, k, ans)
	return result
}
