package string

// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii
// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i

func MinimumTimeToInitialState(word string, k int) int {
	n := len(word)
	LPS := findLPS(word, n) // find LPS of string

	suffix_len := LPS[n-1]

	// find lps until dvisible by k
	for suffix_len > 0 && (n-suffix_len)%k != 0 {
		suffix_len = LPS[suffix_len-1]
	}

	if (n-suffix_len)%k == 0 {
		return (n - suffix_len) / k
	}
	// in case lps is 0 for all cases
	return (n + k - 1) / k // ceil opn
}
