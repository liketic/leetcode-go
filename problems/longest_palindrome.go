// Source : https://leetcode.com/problems/longest-palindrome/
// Author : Ke Li
// Date   : 2016-11-5

// Given a string which consists of lowercase or uppercase letters, find the length of the longest
// palindromes that can be built with those letters.
//
// This is case sensitive, for example "Aa" is not considered a palindrome here.
//
// Note:
// Assume the length of given string will not exceed 1,010.
//
// Example:
//
// Input:
// "abccccdd"
//
// Output:
// 7
//
// Explanation:
// One longest palindrome that can be built is "dccaccd", whose length is 7.
//
package main

import "fmt"

func longestPalindrome(s string) int {
	r := 0
	charNums := make(map[rune]int)
	for _, c := range s {
		charNums[c]++
	}
	for _, v := range charNums {
		r += v - v%2
	}
	if r < len(s) {
		// We can add a single char in the middle of palindrome
		return r + 1
	}
	return r
}

func main() {
	fmt.Println(longestPalindrome("abcded"))
}
