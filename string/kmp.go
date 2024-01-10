package string

// KMP algo to find pattern in given string
func KMP(pat, txt string) []int {
	n := len(txt)
	m := len(pat)

	var result []int // to store pattern found start index
	lps := findLPS(pat, m)

	i := 0
	j := 0

	for i < n {
		if j < m && pat[j] == txt[i] {
			i++
			j++
		}

		if j == m {
			result = append(result, i-j+1)
			j = lps[j-1]
			i++
		} else if pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return result
}

func findLPS(pat string, m int) []int {
	l := 0
	lps := make([]int, m)
	lps[0] = 0

	i := 1
	for i < m {
		if pat[i] == pat[l] {
			l++
			lps[i] = l
			i++
		} else {
			if l != 0 {
				l = lps[l-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}
