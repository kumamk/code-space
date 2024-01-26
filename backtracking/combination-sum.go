package backtracking

// https://leetcode.com/problems/combination-sum/

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	n := len(candidates)

	var solve func(int, int, []int)
	solve = func(idx, sum int, ans []int) {
		// base condition
		if sum == 0 {
			result = append(result, append([]int{}, ans...))
			return
		}

		if sum < 0 {
			return
		}
		// take, explore and not take flow
		for i := idx; i < n; i++ {
			ans = append(ans, candidates[i])
			solve(i, sum-candidates[i], ans)
			ans = ans[:len(ans)-1]
		}
	}
	ans := make([]int, 0)
	solve(0, target, ans)
	return result
}
