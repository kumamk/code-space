package array

//https://leetcode.com/problems/container-with-most-water/

func MaxArea(height []int) int {
	n := len(height)

	i := 0
	j := n - 1
	mw := 0
	// decrease w keeping h maximum
	for i < j {
		w := j - i
		h := min(height[i], height[j])
		a := h * w      // area
		mw = max(mw, a) // update max area

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return mw
}
