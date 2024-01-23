package matrix

// https://leetcode.com/problems/maximal-square/

func MaximalSquare(matrix [][]byte) int {
	r := len(matrix)
	if r == 0 {
		return 0
	}
	c := len(matrix[0])
	dp := make([][]int, r)

	for i := 0; i < r; i++ {
		dp[i] = make([]int, c)
	}

	side := 0
	if matrix[0][0] == '1' {
		side = 1
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else {
				if matrix[i][j] == '1' {
					dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				}
			}
			side = max(side, dp[i][j])
		}
	}
	return side * side
}
