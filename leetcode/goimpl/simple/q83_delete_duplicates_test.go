package simple

import (
	"goimpl/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q83: 删除排序链表中的重复元素
// 与Q26数组删除重复元素类似
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
// 示例 1：
// 输入：head = [1,1,2]
// 输出：[1,2]
//
// 示例 2：
// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]

// left == right, 将left.Next = right，注意考虑边界，设置left.Next = nil
func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	left, right := head, head.Next
	for right != nil {
		if left.Val != right.Val {
			left.Next = right
			left = left.Next
		} else {
			// 相等，切断
			left.Next = nil
		}
		right = right.Next
	}
	return head
}

func Test_deleteDuplicates(t *testing.T) {
	list := common.NewListNode([]int{1, 1, 1, 2, 3, 3, 4, 4, 5, 6})
	r := deleteDuplicates(list)
	assert.Equal(t, "1 -> 2 -> 3 -> 4 -> 5 -> 6", r.String())

	list = common.NewListNode([]int{1, 1, 1})
	r = deleteDuplicates(list)
	assert.Equal(t, "1", r.String())

	list = common.NewListNode([]int{1, 1, 1, 2})
	r = deleteDuplicates(list)
	assert.Equal(t, "1 -> 2", r.String())

	list = common.NewListNode([]int{})
	r = deleteDuplicates(list)
	assert.Nil(t, r)
}
