package Base

import (
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	// 初始化二叉树根节点
	binarySearchTreeRoot := NewBinarySearchTree()
	binarySearchTreeRoot.InsertInto(5)
	binarySearchTreeRoot.InsertInto(2)
	binarySearchTreeRoot.InsertInto(1)
	binarySearchTreeRoot.InsertInto(7)
	//
	//fmt.Println(tree.root.Value)
	//fmt.Println(tree.root.Left.Left.Value)
	//fmt.Println(tree.root.Value)
	//fmt.Println(tree.root.Right.Value)

}
