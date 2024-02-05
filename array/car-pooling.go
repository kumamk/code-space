package array

// https://leetcode.com/problems/car-pooling/

func CarPooling(trips [][]int, capacity int) bool {
	stops := make([]int, 1001)

	for i := 0; i < 1001; i++ {
		stops[i] = 0
	}
	// for each stop add/minus no of passenger
	for _, trip := range trips {
		p := trip[0]
		s := trip[1]
		e := trip[2]

		stops[s] += p
		stops[e] -= p
	}

	total_passenger := 0
	for _, stop := range stops {
		total_passenger += stop
		// if at any point total_passenger is not valid
		if total_passenger > capacity {
			return false
		}
	}
	return true
}
