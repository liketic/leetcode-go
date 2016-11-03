// Source : https://leetcode.com/problems/first-unique-character-in-a-string/
// Author : Ke Li
// Date   : 2016-11-3

// Given a string, find the first non-repeating character in it and return it's index. If it
// doesn't exist, return -1.
//
// Examples:
//
// s = "leetcode"
// return 0.
//
// s = "loveleetcode",
// return 2.
//
// Note: You may assume the string contain only lowercase letters.
//
package main

func firstUniqChar(s string) int {
	dict := make(map[rune]int, 26)
	for _, c := range s {
		if _, ok := dict[c]; !ok {
			dict[c] = 1
		} else {
			dict[c]++
		}
	}
	for i, c := range s {
		if dict[c] == 1 {
			return i
		}
	}
	return -1
}
