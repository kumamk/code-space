package recurmemo

func MaxGifts(list [][]int) int {
	n := len(list)

	var solve func(int, int) int
	solve = func(idx, budget int) int {
		if idx >= n {
			return 0
		}

		ccost := list[idx][0]
		gift_cost := list[idx][1]

		if idx > 0 {
			ccost = ccost - list[idx-1][0] // get actual current budget
		}
		budget += ccost

		gift := 0
		if budget >= gift_cost {
			gift += 1
			budget = budget - gift_cost
		}
		take := gift + solve(idx+1, budget)
		not_take := solve(idx+1, budget)

		return max(take, not_take)
	}
	return solve(0, 0)
}
