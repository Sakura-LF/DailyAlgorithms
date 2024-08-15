package BinaryTree

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	// 定义递归函数
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 递归的终止条件, 如果当前节点为空，则返回空切片
		if node == nil {
			return
		}
		// 中
		res = append(res, node.Val)
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func preorderTraversal2(root *TreeNode) (res []int) {
	// 递归的终止条件, 如果当前节点为空，则返回空切片
	if root == nil {
		return nil
	}

	// 当前节点的值
	currentNodeValue := []int{root.Val}

	// 递归遍历左子树
	leftResult := preorderTraversal2(root.Left)

	// 递归遍历右子树
	rightResult := preorderTraversal2(root.Right)

	// 合并中间节点、左子树和右子树的结果
	return append(append(currentNodeValue, leftResult...), rightResult...)
}

// preorderTraversal 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	// 定义递归函数
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 递归的终止条件, 如果当前节点为空，则返回空切片
		if node == nil {
			return
		}

		// 左
		traversal(node.Left)
		// 中
		res = append(res, node.Val)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// preorderTraversal2 中序遍历
func inorderTraversal2(root *TreeNode) (res []int) {
	// 递归的终止条件, 如果当前节点为空，则返回空切片
	if root == nil {
		return nil
	}

	// 递归遍历左子树
	leftResult := inorderTraversal2(root.Left)

	// 中间节点的值
	currentNodeValue := []int{root.Val}

	// 递归遍历右子树
	rightResult := inorderTraversal2(root.Right)

	// 合并中间节点、左子树和右子树的结果
	return append(append(leftResult, currentNodeValue...), rightResult...)
}

// postorderTraversal 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	// 定义递归函数
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 递归的终止条件, 如果当前节点为空，则返回空切片
		if node == nil {
			return
		}

		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
		// 中
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func postorderTraversal2(root *TreeNode) (res []int) {
	// 递归的终止条件, 如果当前节点为空，则返回空切片
	if root == nil {
		return nil
	}

	// 递归遍历左子树
	leftResult := inorderTraversal2(root.Left)

	// 递归遍历右子树
	rightResult := inorderTraversal2(root.Right)

	// 中间节点的值
	currentNodeValue := []int{root.Val}

	// 合并中间节点、左子树和右子树的结果
	return append(append(leftResult, rightResult...), currentNodeValue...)
}
