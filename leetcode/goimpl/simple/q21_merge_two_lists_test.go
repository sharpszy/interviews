package simple

import (
	"goimpl/common"
	"testing"
)

// Q:21 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例 1：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：
// 输入：l1 = [], l2 = [0]
// 输出：[0]

// 迭代，时间复杂度O(M + N)，空间复杂的O(1)
func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	dumy := &common.ListNode{}
	p := dumy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dumy.Next
}

// 递归，时间复杂度O(M + N)，空间复杂的O(M + N)
func mergeTwoLists2(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	left, right := list1, list2
	if list1.Val >= list2.Val {
		left = list2
		right = list1
	}
	left.Next = mergeTwoLists2(left.Next, right)
	return left
}

func Test_mergeTwoLists(t *testing.T) {
	l1 := common.NewListNode([]int{1, 2, 3})
	l2 := common.NewListNode([]int{1, 3, 4})
	merge := mergeTwoLists(l1, l2)
	merge.Print()

	l3 := common.NewListNode([]int{1, 2, 3})
	l4 := common.NewListNode([]int{1, 3, 4})
	merge2 := mergeTwoLists2(l3, l4)
	merge2.Print()
}
