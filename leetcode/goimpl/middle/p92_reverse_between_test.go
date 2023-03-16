package middle

import (
	"goimpl/common"
	"testing"
)

// Q92: 反转链表 II
// https://leetcode.cn/problems/reverse-linked-list-ii/
// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//
// 示例 1：
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]

func reverseBetween(head *common.ListNode, left int, right int) *common.ListNode {
	dummy := &common.ListNode{}
	dummy.Next = head
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// curr 与 prev 指向的节点保持不变
	curr := prev.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

func Test_reverseBetween(t *testing.T) {
	list := common.NewListNode([]int{1, 2, 3, 4, 5, 6, 7})
	left, right := 2, 4
	r := reverseBetween(list, left, right)
	r.Print()

	list = common.NewListNode([]int{1, 2, 3, 4, 5, 6, 7})
	left, right = 6, 7
	r = reverseBetween(list, left, right)
	r.Print()
}
