package tree

// https://leetcode.com/problems/trim-a-binary-search-tree/description/

func TrimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = TrimBST(root.Left, low, high)
	root.Right = TrimBST(root.Right, low, high)

	// if valid root val return root
	if root.Val >= low && root.Val <= high {
		return root
	}
	// if root left non empty
	if root.Left != nil {
		return root.Left
	}
	// if root right non empty
	if root.Right != nil {
		return root.Right
	}

	return nil
}
