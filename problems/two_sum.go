// Source : https://leetcode.com/problems/two-sum/
// Author : Ke Li
// Date   : 2016-10-26

// Given an array of integers, return indices of the two numbers such that they add up to a specific
// target.
//
// You may assume that each input would have exactly one solution.
//
// Example:
//
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
//
// UPDATE (2016/2/13):
// The return format had been changed to zero-based indices. Please read the above updated
// description carefully.
//
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}
