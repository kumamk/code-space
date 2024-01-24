package bitmanipulation

// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/description/

func MaxLength(arr []string) int {
	var uniqueChars []int // list to hold string in number form

	for _, s := range arr {
		if hasCommon(s) {
			continue
		}
		// convert string to number form
		val := 0
		for _, ch := range s {
			val = val | (1 << (ch - 'a'))
		}
		uniqueChars = append(uniqueChars, val)
	}

	ans := 0
	var solve func(int, int)
	solve = func(idx, tmp int) {
		ans = max(ans, numOfSetBits(tmp))

		for i := idx; i < len(uniqueChars); i++ {
			if (tmp & uniqueChars[i]) == 0 { // 2 numbers not contains any duplicate char
				solve(i+1, (tmp | uniqueChars[i]))
			}
		}
	}

	solve(0, 0)
	return ans
}

// string contains any duplicate char or not
func hasCommon(s string) bool {
	freq := make(map[int]int, 26)

	for _, ch := range s {
		if freq[int(ch-'a')] > 0 {
			return true
		}
		freq[int(ch-'a')]++
	}
	return false
}

// count set bit in number
func numOfSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
