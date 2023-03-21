package common

import (
	"bytes"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var (
		r = &ListNode{}
		p = r
	)
	for _, n := range nums {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}
	return r.Next
}

func (l *ListNode) String() string {
	if l == nil {
		return "Nil"
	}

	var buf bytes.Buffer
	h := l
	for h != nil {
		buf.WriteString(fmt.Sprintf("%v", h.Val))
		if h.Next != nil {
			buf.WriteString(" -> ")
		}
		h = h.Next
	}
	return buf.String()
}
