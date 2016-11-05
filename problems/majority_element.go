// Source : https://leetcode.com/problems/majority-element/
// Author : Ke Li
// Date   : 2016-11-5

// Given an array of size n, find the majority element. The majority element is the element that
// appears more than  n/2  times.
//
// You may assume that the array is non-empty and the majority element always exist in the array.
//
// Credits:Special thanks to @ts for adding this problem and creating all test cases.
//
package main

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
}
