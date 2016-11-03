// Source : https://leetcode.com/problems/intersection-of-two-arrays/
// Author : Ke Li
// Date   : 2016-11-3

// Given two arrays, write a function to compute their intersection.
//
// Example:
// Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
//
// Note:
//
// Each element in the result must be unique.
// The result can be in any order.
//
package main

func intersection(nums1 []int, nums2 []int) []int {
	r, n := []int{}, len(nums1)
	flag := make(map[int]bool)

	for i, count := 0, len(nums2); i < count; i++ {
		for j := 0; j < n; j++ {
			if nums1[j] != nums2[i] {
				continue
			}
			if _, ok := flag[nums2[i]]; !ok {
				r = append(r, nums2[i])
				flag[nums2[i]] = true
			}
		}
	}
	return r
}
