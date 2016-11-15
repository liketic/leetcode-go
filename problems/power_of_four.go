// Source : https://leetcode.com/problems/power-of-four/
// Author : Ke Li
// Date   : 2016-11-15

// Given an integer (signed 32 bits), write a function to check whether it is a power of 4.
//
// Example:
// Given num = 16, return true.
// Given num = 5, return false.
//
// Follow up: Could you solve it without loops/recursion?
//
// Credits:Special thanks to @yukuairoy  for adding this problem and creating all test cases.
//
package main

func isPowerOfFour(num int) bool {
	for num > 1 {
		if num%4 != 0 {
			return false
		}
		num /= 4
	}
	return num == 1
}
