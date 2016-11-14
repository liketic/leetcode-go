// Source : https://leetcode.com/problems/happy-number/
// Author : Ke Li
// Date   : 2016-11-14

// Write an algorithm to determine if a number is "happy".
//
// A happy number is a number defined by the following process: Starting with any positive integer,
// replace the number by the sum of the squares of its digits, and repeat the process until the
// number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy numbers.
//
// Example:19 is a happy number
//
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1
//
// Credits:Special thanks to @mithmatt and @ts for adding this problem and creating all test cases.
//
package main

import "github.com/smartystreets/assertions"

func isHappy(n int) bool {
	dict := make(map[int]bool)
	for n > 1 {
		if _, ok := dict[n]; ok {
			return false
		}
		dict[n] = true
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return n == 1
}

func main() {
	assertions.ShouldBeTrue(isHappy(19))
}
