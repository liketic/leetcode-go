// Source : https://leetcode.com/problems/sum-of-left-leaves/
// Author : Ke Li
// Date   : 2016-11-3

// Find the sum of all left leaves in a given binary tree.
//
// Example:
//
// 3
// / \
// 9  20
// /  \
// 15   7
//
// There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
//
package main

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}
	r := 0
	if root.Left.Left == nil && root.Left.Right == nil {
		r += root.Left.Val
	} else {
		r += sumOfLeftLeaves(root.Left)
	}
	return r + sumOfLeftLeaves(root.Right)
}
