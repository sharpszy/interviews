package common

import "fmt"

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

func (l *ListNode) Print() {
	if l == nil {
		return
	}
	h := l
	for h != nil {
		fmt.Printf("%v", h.Val)
		if h.Next != nil {
			fmt.Print(" -> ")
		}
		h = h.Next
	}
	fmt.Println()
}
