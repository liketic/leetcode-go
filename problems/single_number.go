// Source : https://leetcode.com/problems/single-number/
// Author : Ke Li
// Date   : 2016-11-3

// Given an array of integers, every element appears twice except for one. Find that single one.
//
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using
// extra memory?
//
package main

func singleNumber(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}
