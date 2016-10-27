// Source : https://leetcode.com/problems/palindrome-number/
// Author : Ke Li
// Date   : 2016-10-26

// Determine whether an integer is a palindrome. Do this without extra space.
//
// click to show spoilers.
//
package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 10 {
		return x >= 0
	}
	digitNums := 1
	for t := x; t >= 10; t /= 10 {
		digitNums *= 10
	}

	for digitNums > 0 {
		if x%10 != (x/digitNums)%10 {
			return false
		}
		digitNums /= 100
		x /= 10
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(100))
	fmt.Println(isPalindrome(121))
}
