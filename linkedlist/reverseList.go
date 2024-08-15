package linkedlist

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		// 先保存 cur.Next这个值
		tmp := cur.Next
		// 使得 cur.Next 指向上一个
		cur.Next = pre
		// 将 pre 向前移动
		pre = cur
		// tmp 就是之前保存的 cur.Next
		cur = tmp
	}
	return cur
}
