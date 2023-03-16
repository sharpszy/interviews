package simple

import (
	"goimpl/common"
	"testing"
)

// Q206: 反转链表
// https://leetcode.cn/problems/reverse-linked-list/
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

func reverseList(head *common.ListNode) *common.ListNode {
	var prev *common.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func Test_reverseList(t *testing.T) {
	list := common.NewListNode([]int{1, 2, 3, 4, 5, 6, 7})
	r := reverseList(list)
	r.Print()
}
