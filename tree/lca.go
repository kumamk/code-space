package tree

//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// base case
	if root == q || root == p {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	// root is LCA
	if left != nil && right != nil {
		return root
	}
	// left is LCA
	if left != nil {
		return left
	}

	return right
}
