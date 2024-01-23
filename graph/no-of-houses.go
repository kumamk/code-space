package graph

// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/

// using floyd-warshall
func CountOfPairs(n int, x int, y int) []int {
	grid := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		grid[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			grid[i][j] = 100000
		}
	}

	for i := 1; i < n; i++ {
		grid[i][i+1] = 1
		grid[i+1][i] = 1
	}

	grid[x][y] = 1
	grid[y][x] = 1

	for via := 1; via <= n; via++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if i == j {
					grid[i][j] = 0
				}
				grid[i][j] = min(grid[i][j], grid[i][via]+grid[via][j])
			}
		}
	}

	result := make([]int, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				continue
			}

			v := grid[i][j]
			result[v-1]++
		}
	}
	return result
}
