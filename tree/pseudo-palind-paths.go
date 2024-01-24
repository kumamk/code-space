package tree

// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/description/

func PseudoPalindromicPaths(root *TreeNode) int {
	ans := 0
	var solve func(*TreeNode, int)
	solve = func(root *TreeNode, count int) {
		if root == nil {
			return
		}

		count = (count ^ (1 << root.Val)) // xor of number to remove duplicate

		if root.Left == nil && root.Right == nil {
			if (count & (count - 1)) == 0 { // to check if only 1 set bit
				ans++
			}
		}
		solve(root.Left, count)
		solve(root.Right, count)
	}
	solve(root, 0)
	return ans
}
