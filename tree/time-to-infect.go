package tree

// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/description/

func AmountOfTimeByBFS(root *TreeNode, start int) int {
	var que []int
	visited := make(map[int]bool)
	adj := make(map[int][]int) // make adj list for graph traversal

	var makeGraph func(root *TreeNode, parent int)
	makeGraph = func(root *TreeNode, parent int) {
		if root == nil {
			return
		}

		if parent != -1 {
			adj[root.Val] = append(adj[root.Val], parent)
		}

		if root.Left != nil {
			adj[root.Val] = append(adj[root.Val], root.Left.Val)
		}
		if root.Right != nil {
			adj[root.Val] = append(adj[root.Val], root.Right.Val)
		}

		makeGraph(root.Left, root.Val)
		makeGraph(root.Right, root.Val)
	}

	makeGraph(root, -1)
	// BFS traversal
	ans := 0
	que = append(que, start)
	visited[start] = true

	for len(que) > 0 {
		size := len(que)

		for size > 0 {
			u := que[0]
			que = que[1:]

			for _, v := range adj[u] {
				if !visited[v] {
					que = append(que, v)
					visited[v] = true
				}
			}
			size = size - 1
		}
		ans++
	}
	return ans - 1
}

// dfs with height of tree flow
func AmountOfTimeByDFS(root *TreeNode, start int) int {
	result := 0

	var solve func(*TreeNode, int) int
	solve = func(root *TreeNode, start int) int {
		if root == nil {
			return 0
		}

		lh := solve(root.Left, start)
		rh := solve(root.Right, start)

		if root.Val == start {
			result = max(lh, rh)
			return -1 // to say start node is found
		} else if lh >= 0 && rh >= 0 {
			return max(lh, rh) + 1 // normal height logic for tree
		} else {
			d := abs(lh) + abs(rh)
			result = max(result, d)
			return min(lh, rh) - 1 // add -1 to -ve value so far
		}
	}
	solve(root, start)
	return result
}

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}
