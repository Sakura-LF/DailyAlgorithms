package linkedlist

import (
	"fmt"
)

type SingleNode struct {
	Val  int         // 节点的值
	Next *SingleNode // 下一个节点的指针
}

type MyLinkedList struct {
	dummyHead *SingleNode // 虚拟节点
	Size      int         // 链表长度
}

// Constructor 构建链表
func Constructor() MyLinkedList {
	dummyNode := &SingleNode{
		Val:  0,
		Next: nil,
	}
	return MyLinkedList{
		dummyHead: dummyNode,
		Size:      0,
	}
}

// Get 返回指定索引的链表元素
func (this *MyLinkedList) Get(index int) int {
	// 判断索引是否合法
	if this == nil || index >= this.Size || index < 0 {
		return -1
	}
	// curNode 就是头节点
	curNode := this.dummyHead.Next
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	return curNode.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 创建新节点
	newNode := &SingleNode{
		Val: val,
	}
	// 将虚拟头节点的 Next 指向 新节点的 Next
	// 因为 dummyHead.Next 是之前的头节点
	newNode.Next = this.dummyHead.Next
	this.dummyHead.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &SingleNode{
		Val: val,
	}
	curNode := this.dummyHead
	// 遍历到最后一个节点
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	// 循环结束说明 curNode.Next == nil 了
	// curNode 已经是最后一饿节点了
	curNode.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 判断索引是否合法
	// 这里的 index 不用 >=
	// 因为 = size 的时候就是添加到尾部
	if this == nil || index > this.Size || index < 0 {
		return
	}

	newNode := &SingleNode{
		Val: val,
	}
	// 因为要遍历到指定索引的前一个节点
	// 所以这里不能 this.dummyHead.Next, 因为 index 可能为 0
	curNode := this.dummyHead

	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	// 先把后面的(Next) 指向新节点
	newNode.Next = curNode.Next
	// 新节点的 Next 指向新节点
	curNode.Next = newNode

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 索引无效直接返回
	if index < 0 || index >= this.Size {
		return
	}
	curNode := this.dummyHead
	// 遍历到要删除的前一个节点
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	// 将目前节点的 Next 指向 Next.Next
	curNode.Next = curNode.Next.Next
	this.Size--
}

// Print 打印链表
func (this *MyLinkedList) Print() {
	curNode := this.dummyHead.Next
	for curNode.Next != nil {
		fmt.Printf("%d -> ", curNode.Val)
		curNode = curNode.Next
	}
	// 循环结束说明已经到最后一个节点了 curNode.Next == nil
	fmt.Printf("%d -> nil \n", curNode.Val)

}
