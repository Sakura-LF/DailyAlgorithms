package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	// 复制一份虚拟头节点,为了不影响
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		// 需要保存的两个临时值
		firstNext := cur.Next
		SecondNext := cur.Next.Next.Next

		// 第一步, 虚拟头节点的 next 指向 2
		cur.Next = cur.Next.Next
		// 第二步, 2 指向 保存的第一个值
		cur.Next.Next = firstNext
		// 第三步, 1 指向 保存的第二个值
		cur.Next.Next.Next = SecondNext

		// cur 移动两位,准备下一次交换
		cur = cur.Next.Next
	}
	return dummyHead.Next
}
