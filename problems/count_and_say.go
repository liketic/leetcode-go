// Source : https://leetcode.com/problems/count-and-say/
// Author : Ke Li
// Date   : 2016-10-28

// The count-and-say sequence is the sequence of integers beginning as follows:
// 1, 11, 21, 1211, 111221, ...
//
// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.
//
// Given an integer n, generate the nth sequence.
//
// Note: The sequence of integers will be represented as a string.
//
package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		temp := ""
		count := 0
		for i, n := 0, len(s); i < n; i++ {
			count++
			if (i < n-1 && s[i] != s[i+1]) || i == n-1 {
				temp += strconv.Itoa(count) + string(s[i])
				count = 0
			}
		}
		s = temp
	}
	return s
}

func main() {
	fmt.Println(countAndSay(5))
}
