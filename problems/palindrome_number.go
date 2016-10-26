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
	if x < 0 {
		return false
	}
	digits := []int{}
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(100))
	fmt.Println(isPalindrome(121))
}
