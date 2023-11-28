package string

// https://leetcode.com/problems/longest-common-subsequence/description/

func LongestCommonSubsequence(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	var mem [1001][1001]int

	// init mem 2d slice
	for i := 0; i < 1001; i++ {
		for j := 0; j < 1001; j++ {
			mem[i][j] = -1
		}
	}

	var solve func(string, string, int, int) int

	solve = func(s1 string, s2 string, i, j int) int {
		// out of bound check
		if i >= m || j >= n {
			return 0
		}
		// return if already processed
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		// update when both ch are same
		if s1[i] == s2[j] {
			mem[i][j] = 1 + solve(s1, s2, i+1, j+1)
			return mem[i][j]
		}
		// try for both cases
		mem[i][j] = max(solve(s1, s2, i+1, j), solve(s1, s2, i, j+1))
		return mem[i][j]
	}
	return solve(s1, s2, 0, 0)
}
