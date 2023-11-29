package bitmanipulation

//https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/description/

func MinFlips(a int, b int, c int) int {
	flip := 0

	// check if 1 of 3 is finsihed
	for a != 0 || b != 0 || c != 0 {
		if c&1 == 1 {
			if a&1 == 0 && b&1 == 0 {
				flip++
			}
		} else {
			if a&1 == 1 {
				flip++
			}
			if b&1 == 1 {
				flip++
			}
		}
		// right shift all 3
		a >>= 1
		b >>= 1
		c >>= 1
	}

	return flip
}
