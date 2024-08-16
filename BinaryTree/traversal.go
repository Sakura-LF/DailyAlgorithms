package BinaryTree

import (
	"container/list"
)

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

	// 追加中间节点的值到 res
	res = append(res, root.Val)

	leftResult := preorderTraversal2(root.Left)

	// 递归遍历右子树
	rightResult := preorderTraversal2(root.Right)

	// 再 追加 lestResult 和 rightResult
	res = append(res, leftResult...)
	res = append(res, rightResult...)

	// 最后的一次遍历结果也可以改成这样
	return res
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

// 二叉树层序遍历, 使用 container 包模拟队列

func levelOrder(root *TreeNode) (res []int) {
	// 1.首先判断 root 是否为空
	if root == nil {
		return nil
	}
	// 初始化一个队列
	queue := list.New()
	// 将根节点放入队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		// 出队
		node := queue.Remove(queue.Front()).(*TreeNode)
		res = append(res, node.Val)

		if node.Left != nil {
			// 入队
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			// 入队
			queue.PushBack(node.Right)
		}
	}
	return res
}

func levelOrder2(root *TreeNode) (res [][]int) {
	// 1.首先判断 root 是否为空
	if root == nil {
		return nil
	}
	// 初始化一个队列
	queue := list.New()
	// 将根节点放入队列
	queue.PushBack(root)

	for queue.Len() > 0 {
		// 获取当前层的节点个数
		length := queue.Len()
		// 因为要返回每层的值元素,所以需要一个临时切片存储每层的值
		var temp []int
		for i := 0; i < length; i++ {
			// 出队
			node := queue.Remove(queue.Front()).(*TreeNode)
			temp = append(temp, node.Val)

			if node.Left != nil {
				// 入队
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				// 入队
				queue.PushBack(node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
