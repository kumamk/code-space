package array

//https://leetcode.com/problems/majority-element-ii/

func MajorityElement(nums []int) []int {
	n := len(nums)

	count1 := 0
	ans1 := -1

	count2 := 0
	ans2 := -1

	for i := 0; i < n; i++ {
		if nums[i] == ans1 {
			count1++
		} else if nums[i] == ans2 {
			count2++
		} else if count1 == 0 {
			count1 = 1
			ans1 = nums[i]
		} else if count2 == 0 {
			count2 = 1
			ans2 = nums[i]
		} else {
			count1--
			count2--
		}
	}
	f1 := 0
	f2 := 0
	var result []int

	for i := 0; i < n; i++ {
		if nums[i] == ans1 {
			f1++
		} else if nums[i] == ans2 {
			f2++
		}
	}
	if f1 > n/3 {
		result = append(result, ans1)
	}
	if f2 > n/3 {
		result = append(result, ans2)
	}

	return result
}
