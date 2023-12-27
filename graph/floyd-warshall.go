package graph

func MinimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	grid := make([][]int, 26)
	const MAX = 100000000

	// init and prefill with max value
	for i := 0; i < 26; i++ {
		grid[i] = make([]int, 26)
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j {
				grid[i][j] = 0
			} else {
				grid[i][j] = MAX
			}
		}
	}

	// update cost as per given list
	n := len(original)
	for i := 0; i < n; i++ {
		x := original[i] - 'a'
		y := changed[i] - 'a'
		grid[x][y] = min(cost[i], grid[x][y])
	}

	//Floyd Warshall to update min cost from a -> b
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				grid[i][j] = min(grid[i][j], grid[i][k]+grid[k][j])
			}
		}
	}

	n = len(source)
	var result int64
	// update cost using updated grid
	for i := 0; i < n; i++ {
		if source[i] == target[i] {
			continue
		}
		x := source[i] - 'a'
		y := target[i] - 'a'

		if grid[x][y] == MAX {
			return -1
		}

		result += int64(grid[x][y])
	}

	return result
}
