package backtracking

func Permute(nums []int) [][]int {
	n := len(nums)
	var result [][]int

	var solve func(int, []int)
	solve = func(idx int, ans []int) {
		if idx == n {
			result = append(result, append([]int{}, ans...))
			return
		}

		if idx > n {
			return
		}

		for i := idx; i < n; i++ {
			ans[i], ans[idx] = ans[idx], ans[i] // swap
			solve(idx+1, ans)                   // explore
			ans[i], ans[idx] = ans[idx], ans[i] // undo swap
		}
	}
	solve(0, nums)
	return result
}
