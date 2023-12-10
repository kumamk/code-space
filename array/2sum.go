package array

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func TwoSum(numbers []int, target int) []int {
	n := len(numbers)

	l := 0
	r := n - 1

	for l < r {
		if numbers[r]+numbers[l] > target {
			r--
		} else if numbers[r]+numbers[l] < target {
			l++
		} else {
			break
		}
	}
	var result []int
	result = append(result, l+1, r+1)
	return result
}
