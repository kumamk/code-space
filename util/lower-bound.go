package util

// given list and target element find lower bound
func LowerBound(list []int, target int) int {
	left := 0
	right := len(list) - 1
	result := len(list)

	for left <= right {
		mid := left + (right-left)/2

		if list[mid] >= target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
