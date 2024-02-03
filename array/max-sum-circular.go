package array

//https://leetcode.com/problems/maximum-sum-circular-subarray/

func MaxSubarraySumCircular(nums []int) int {
	n := len(nums)

	total_sum := 0

	for i := 0; i < n; i++ {
		total_sum += nums[i]
	}

	max_sum := kadaneMax(nums)
	min_sum := kadaneMin(nums)

	cir_max := total_sum - min_sum

	if max_sum > 0 {
		return max(max_sum, cir_max)
	}
	return max_sum
}

func kadaneMax(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func kadaneMin(nums []int) int {
	sum := nums[0]
	minSum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = min(sum+nums[i], nums[i])
		minSum = min(minSum, sum)
	}
	return minSum
}
