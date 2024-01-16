package bitmanipulation

import (
	"math"
)

//https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/description/

func FindMaximumNumber(k int64, x int) int64 {
	var bitCount []int64
	left := int64(0)
	right := int64(1e15)
	ans := int64(0)

	var getBitCount func(int64)
	getBitCount = func(num int64) {
		if num == 0 {
			return
		}

		if num == 1 {
			bitCount[0]++
			return
		}

		if num == 2 {
			bitCount[0]++
			bitCount[1]++
			return
		}

		bitLen := int64(math.Log2(float64(num)))
		closePowerTwo := int64(1 << bitLen)
		bitCount[bitLen] += (num - closePowerTwo + 1)

		for i := int64(0); i < bitLen; i++ {
			bitCount[int(i)] += (closePowerTwo / 2)
		}
		num = num - closePowerTwo
		getBitCount(num)
	}

	for left <= right {
		mid := left + (right-left)/2

		bitCount = make([]int64, 65)
		getBitCount(mid)
		totalCount := int64(0)

		for i := int64(0); i < int64(math.Log2(float64(mid)))+1; i++ {
			if int(i+1)%x == 0 {
				totalCount += bitCount[i]
			}
		}

		if totalCount <= k {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
