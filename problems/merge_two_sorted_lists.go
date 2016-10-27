// Source : https://leetcode.com/problems/merge-two-sorted-lists/
// Author : Ke Li
// Date   : 2016-10-27

// Merge two sorted linked lists and return it as a new list. The new list should be made by
// splicing together the nodes of the first two lists.
//
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	link := &ListNode{}
	head := link

	for l1 != nil && l2 != nil {
		var v int
		if l1.Val < l2.Val {
			v = l1.Val
			l1 = l1.Next
		} else {
			v = l2.Val
			l2 = l2.Next
		}
		link.Val = v
		link.Next = &ListNode{}
		link = link.Next
	}
	if l1 != nil {
		link.Val = l1.Val
		link.Next = l1.Next
	} else if l2 != nil {
		link.Val = l2.Val
		link.Next = l2.Next
	}
	return head
}

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l := mergeTwoLists(l1, l2)
	for ; l != nil; l = l.Next {
		fmt.Println(l.Val)
	}
}
