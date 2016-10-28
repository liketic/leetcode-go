// Source : https://leetcode.com/problems/add-two-numbers/
// Author : Ke Li
// Date   : 2016-10-26

// You are given two linked lists representing two non-negative numbers. The digits are stored in
// reverse order and each of their nodes contain a single digit. Add the two numbers and return it
// as a linked list.
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
//
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := l1

	for l1 != nil && l2 != nil {
		l1.Val += l2.Val
		if l1.Next == nil {
			l1.Next = l2.Next
			break
		}
		l1 = l1.Next
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}
	sum := 0
	for it := head; it != nil; it = it.Next {
		it.Val += sum
		sum = it.Val / 10
		it.Val %= 10
		if it.Next == nil && sum > 0 {
			it.Next = &ListNode{Val: sum}
			break
		}
	}
	return head
}

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	link := addTwoNumbers(l1, l2)
	for link != nil {
		fmt.Println(link.Val)
		link = link.Next
	}
}
