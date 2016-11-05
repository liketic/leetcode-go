// Source : https://leetcode.com/problems/intersection-of-two-arrays-ii/
// Author : Ke Li
// Date   : 2016-11-5

// Given two arrays, write a function to compute their intersection.
//
// Example:
// Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].
//
// Note:
//
// Each element in the result should appear as many times as it shows in both arrays.
// The result can be in any order.
//
// Follow up:
//
// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot
// load all elements into the memory at once?
//
package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	r := []int{}
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums2[j] == nums1[i] {
			r = append(r, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return r
}
