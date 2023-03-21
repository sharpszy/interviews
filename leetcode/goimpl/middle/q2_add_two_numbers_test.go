package middle

import (
	"goimpl/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q2: 两数相加
// https://leetcode.cn/problems/add-two-numbers/
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例 1：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
//
// 示例 2：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	var (
		p, r *common.ListNode

		carry = 0
		sum   = 0
	)
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum = carry + n1 + n2
		carry = sum / 10
		sum = sum % 10

		if r == nil {
			r = &common.ListNode{Val: sum}
			p = r
		} else {
			p.Next = &common.ListNode{Val: sum}
			p = p.Next
		}
	}

	if carry > 0 {
		p.Next = &common.ListNode{Val: carry}
	}
	return r
}

func Test_addTwoNumbers(t *testing.T) {
	var (
		l1 = common.NewListNode([]int{1, 2, 3})
		l2 = common.NewListNode([]int{4, 5, 8, 3})
	)
	l := addTwoNumbers(l1, l2)
	assert.Equal(t, "5 -> 7 -> 1 -> 4", l.String())

	l1 = common.NewListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = common.NewListNode([]int{9, 9, 9, 9})
	l = addTwoNumbers(l1, l2)
	assert.Equal(t, "8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1", l.String())
}
