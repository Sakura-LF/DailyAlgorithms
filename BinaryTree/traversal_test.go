package BinaryTree

import (
	"fmt"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {

}

func TestInOrderTraversal(t *testing.T) {
	// 构建一个简单的二叉树
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}

	result := InOrderTraversal(root)
	fmt.Println(result)
}
