// Source : https://leetcode.com/problems/ransom-note/
// Author : Ke Li
// Date   : 2016-11-3

// Given an arbitrary ransom note string and another string containing letters from all the
// magazines, write a function that will return true if the ransom
// note can be constructed from the magazines ; otherwise, it will return false.
//
// Each letter in the magazine string can only be used once in your ransom note.
//
// Note:
// You may assume that both strings contain only lowercase letters.
//
// canConstruct("a", "b") -> false
// canConstruct("aa", "ab") -> false
// canConstruct("aa", "aab") -> true
//
package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	idx := make([]int, 26)
	for _, c := range ransomNote {
		idx[c-'a']++
	}
	for _, c := range magazine {
		idx[c-'a']--
	}
	for i := 0; i < 26; i++ {
		if idx[i] > 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
