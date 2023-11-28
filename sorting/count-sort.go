package sorting

import "math"

// count sort
func CountSort(nums []int) []int {
	var result []int
	freq := make(map[int]int)
	n := len(nums)

	// fill the freq map
	for i := 0; i < n; i++ {
		freq[nums[i]]++
	}

	minE := math.MaxInt
	maxE := math.MinInt

	// find min and max from given list
	for _, v := range nums {
		if v < minE {
			minE = v
		}

		if v > maxE {
			maxE = v
		}
	}
	// count sort for sorted list
	k := 0
	for i := minE; i <= maxE; i++ {
		if freq[i] > 0 {
			for freq[i] > 0 {
				result = append(result, i)
				k++
				freq[i]--
			}
		}
	}
	return result
}
