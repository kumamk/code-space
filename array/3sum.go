package array

import "sort"

//https://leetcode.com/problems/3sum/

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	var result [][]int

	if n < 3 {
		return [][]int{}
	}

	var twoSum func(int, int, int)
	twoSum = func(l, r, target int) {
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				result = append(result, []int{-target, nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	// sort list
	sort.Ints(nums)
	// for each nums[i], find other 2 by 2sum approach
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		t := nums[i]
		twoSum(i+1, n-1, -t)
	}
	return result
}
