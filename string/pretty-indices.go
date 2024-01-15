package string

func BeautifulIndices(s string, a string, b string, k int) []int {
	a_list := KMP(a, s)
	b_list := KMP(b, s)

	var result []int
	i := 0
	j := 0

	for i < len(a_list) && j < len(b_list) {
		if abs(a_list[i]-b_list[j]) <= k {
			result = append(result, a_list[i])
			i++
		} else if b_list[j]-a_list[i] > k {
			i++
		} else {
			j++
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
