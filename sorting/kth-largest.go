package sorting

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/

// O(n) on an avg, but can go On*n for worst case, quick select sort
func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	l := 0
	r := n - 1

	var getPivot func(int, int) int
	getPivot = func(l, r int) int {
		p := nums[l]
		i := l + 1
		j := r

		for i <= j {
			if nums[i] < p && nums[j] > p {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}

			if nums[i] >= p {
				i++
			}

			if nums[j] <= p {
				j--
			}
		}
		nums[l], nums[j] = nums[j], nums[l]
		return j
	}

	pivot_i := 0
	for {
		pivot_i = getPivot(l, r)

		if pivot_i == k-1 {
			break
		} else if pivot_i > k-1 {
			r = pivot_i - 1
		} else {
			l = pivot_i + 1
		}
	}

	return nums[pivot_i]
}
