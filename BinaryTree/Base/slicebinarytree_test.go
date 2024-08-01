package Base

import (
	"fmt"
	"testing"
)

func TestNewSliceBinaryTree(t *testing.T) {
	Tree := []any{1, 2, 3, nil}
	tree := NewSliceBinaryTree(Tree)
	fmt.Println(len(tree.Tree))
}
