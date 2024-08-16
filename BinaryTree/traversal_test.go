package BinaryTree

import (
	"fmt"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {

}

func TestInOrderTraversal(t *testing.T) {
	// 构建一个简单的二叉树
	_ = &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}

	//result := InOrderTraversal(root)
	//fmt.Println(result)
}

// 二叉树遍历层序测试
func Test_levelOrder(t *testing.T) {
	// 构建节点
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}

	// 构建二叉树
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6

	// 返回一维数组的写法
	result1 := levelOrder(root)
	fmt.Println(result1)

	// 返回二维数组的写法
	result2 := levelOrder2(root)
	fmt.Println(result2)

	s := "Sakura"
	ss := []byte(s)
	ss[0] = '1'
	fmt.Println(string(ss))
}
