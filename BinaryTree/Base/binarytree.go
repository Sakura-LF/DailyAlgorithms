package Base

import "container/list"

// TreeNode 二叉树节点
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 构建二叉树节点
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func levelOrder(root *TreeNode) []int {
	// 初始化队列
	queue := list.New()
	queue.PushBack(root)
	nums := make([]int, 0)

	for queue.Len() > 0 {
		// 队列出队
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Value)
		// 如果左子节点不为nil,入队
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}

// 前序遍历
func PreOrder(node *TreeNode) []int {
	var result []int
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Value)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(node)
	return result
}

// 中序遍历
func MiddleOrder(node *TreeNode) []int {
	var result []int
	var middleOrder func(node *TreeNode)
	middleOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		middleOrder(node.Left)
		result = append(result, node.Value)
		middleOrder(node.Right)
	}
	middleOrder(node)
	return result
}

// 中序遍历
func PostOrder(node *TreeNode) []int {
	var result []int
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		result = append(result, node.Value)
	}
	postOrder(node)
	return result
}
