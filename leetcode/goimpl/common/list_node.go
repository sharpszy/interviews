package common

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type NodeType interface {
	constraints.Integer
}

// Definition for singly-linked list.
type ListNode[T NodeType] struct {
	Val  T
	Next *ListNode[T]
}

func NewListNode[T NodeType](nums []T) *ListNode[T] {
	if len(nums) == 0 {
		return nil
	}
	var (
		r = &ListNode[T]{}
		p = r
	)
	for _, n := range nums {
		p.Next = &ListNode[T]{Val: n}
		p = p.Next
	}
	return r.Next
}

func (l *ListNode[T]) Print() {
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
