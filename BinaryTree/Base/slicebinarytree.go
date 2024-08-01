package Base

// SliceBinaryTree 二叉树数组表示
type SliceBinaryTree struct {
	// 因为表示任意二叉树需要有 nil,所以类型为 any
	Tree []any
}

// NewSliceBinaryTree 构建数组二叉树
func NewSliceBinaryTree(tree []any) *SliceBinaryTree {
	return &SliceBinaryTree{
		Tree: tree,
	}
}
