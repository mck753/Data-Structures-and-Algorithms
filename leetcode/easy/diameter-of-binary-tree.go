package easy

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := maxDepth(root.Left)
		r := maxDepth(root.Right)
		ret = max(ret, l+r)

		return 1 + max(l, r)
	}

	maxDepth(root)

	return ret
}
