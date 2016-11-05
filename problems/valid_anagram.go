// Source : https://leetcode.com/problems/valid-anagram/
// Author : Ke Li
// Date   : 2016-11-5

// Given two strings s and t, write a function to determine if t is an anagram of s.
//
// For example,
// s = "anagram", t = "nagaram", return true.
// s = "rat", t = "car", return false.
//
// Note:
// You may assume the string contains only lowercase alphabets.
//
// Follow up:
// What if the inputs contain unicode characters? How would you adapt your solution to such case?
//
package main

import (
	"../utils"
)

func isAnagram(s string, t string) bool {
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	for _, c := range t {
		count[c]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	utils.MustTrue(isAnagram("axxx", "xxxa"))
}
