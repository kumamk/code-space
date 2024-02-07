package array

// https://leetcode.com/problems/max-consecutive-ones-iii/

func LongestOnes(nums []int, k int) int {
	n := len(nums)

	if k >= n {
		return n
	}

	ans := 0
	count := 0
	i, j := 0, 0

	for i < n {
		if nums[i] == 0 {
			count++
		}
		// if 0 count > k
		if count > k {
			if nums[j] == 0 {
				count-- // decrease count if prev 0 leaving
			}
			j++
		}
		ans = max(ans, i-j+1) // update ans on each iteration
		i++
	}
	return ans
}
