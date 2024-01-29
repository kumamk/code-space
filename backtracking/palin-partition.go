package backtracking

// https://leetcode.com/problems/palindrome-partitioning/description/

// O(n*2 to power n)
func Partition(s string) [][]string {
	n := len(s)
	var result [][]string

	var solve func(int, []string)
	solve = func(idx int, ans []string) {
		if idx >= n {
			result = append(result, append([]string{}, ans...))
			return
		}

		for i := idx; i < n; i++ {
			if isPalin(s, idx, i) {
				ans = append(ans, s[idx:i+1]) // do
				solve(i+1, ans)               // explore
				ans = ans[:len(ans)-1]        // undo
			}
		}
	}
	ans := make([]string, 0)
	solve(0, ans)
	return result
}

func isPalin(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
