// Source : https://leetcode.com/problems/power-of-two/
// Author : Ke Li
// Date   : 2016-11-14

// Given an integer, write a function to determine if it is a power of two.
//
// Credits:Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
//
package main

func isPowerOfTwo(n int) bool {
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return n == 1
}
