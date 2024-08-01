package Base

import (
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	// 定义节点
	node1 := NewTreeNode(1)
	node2 := NewTreeNode(2)
	node3 := NewTreeNode(3)
	node4 := NewTreeNode(4)
	node5 := NewTreeNode(5)
	node6 := NewTreeNode(6)
	node7 := NewTreeNode(7)

	// 构建二叉树
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	// 层序遍历
	//order := levelOrder(node1)
	//fmt.Println("层序遍历二叉树：", order)

	// 前序遍历
	//preOrderResult := PreOrder(node1)
	//fmt.Println("前序遍历:", preOrderResult)
	//
	//// 中序遍历
	//middleOrderResult := MiddleOrder(node1)
	//fmt.Println("中序遍历:", middleOrderResult)
	//
	//// 后序遍历
	//postOrderResult := PostOrder(node1)
	//fmt.Println("后序遍历:", postOrderResult)
}
func TestSliceTree(t *testing.T) {

}
