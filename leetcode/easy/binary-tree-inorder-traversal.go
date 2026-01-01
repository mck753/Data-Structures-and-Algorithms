package easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	r1 := inorderTraversal(root.Left)
	r1 = append(r1, root.Val)
	r2 := inorderTraversal(root.Right)
	r1 = append(r1, r2...)
	return r1
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{root.Val}
	r1 := preorderTraversal(root.Left)
	r2 := preorderTraversal(root.Right)

	ret = append(ret, r1...)
	ret = append(ret, r2...)

	return ret
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	r1 := postorderTraversal(root.Left)
	r2 := postorderTraversal(root.Right)

	ret = append(ret, r1...)
	ret = append(ret, r2...)
	ret = append(ret, root.Val)

	return ret
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ret = append(ret, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return ret
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		tmpArr := make([]int, 0)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmpArr = append(tmpArr, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ret = append(ret, tmpArr)
	}

	return ret
}

func levelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	// 记录当前遍历到的层数（根节点视为第 1 层）
	depth := 1

	for len(q) > 0 {
		// 获取当前队列长度
		sz := len(q)
		for i := 0; i < sz; i++ {
			// 弹出队列头
			cur := q[0]
			q = q[1:]
			// 访问 cur 节点，同时知道它所在的层数
			fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)

			// 把 cur 的左右子节点加入队列
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
}
