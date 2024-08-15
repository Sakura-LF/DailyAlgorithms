package Base

// BinarySearchTree 定义二叉搜索树
type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

// InsertInto 插入
func (b *BinarySearchTree) InsertInto(value int) {
	curNode := b.root
	// 若树为空,初始化根节点
	if curNode == nil {
		b.root = NewTreeNode(value)
		return
	}

	// 待插入节点之前的节点
	var pre *TreeNode = nil
	// 循环查找
	for curNode != nil {
		if curNode.Val == value {
			return
		}
		if curNode.Val < value {
			pre = curNode
			curNode = curNode.Right
		} else {
			pre = curNode
			curNode = curNode.Left
		}
	}
	// 插入节点
	node := NewTreeNode(value)
	if pre.Val < value {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

// 删除

// Search 查询
func (b *BinarySearchTree) Search(value int) *TreeNode {
	curNode := b.root
	for curNode != nil {
		if curNode.Val < value {
			curNode = curNode.Right
		}
		if curNode.Val > value {
			curNode = curNode.Left
		}
		if curNode.Val == value {
			return curNode
		}
	}
	return nil
}
