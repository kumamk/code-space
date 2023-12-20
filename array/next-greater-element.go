package array

func NGE(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	stack := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = -1
	}
	stack = append(stack, 0)

	for i := 1; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			result[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
