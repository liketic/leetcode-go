// Source : https://leetcode.com/problems/valid-parentheses/
// Author : Ke Li
// Date   : 2016-10-26

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
//
package main

import "fmt"

func match(r rune) rune {
	switch r {
	case ']':
		return '['
	case ')':
		return '('
	case '}':
		return '{'
	}
	return '-'
}

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			stack = append(stack, c)
		} else {
			if n := len(stack); n > 0 && stack[n-1] == match(c) {
				stack = stack[:n-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[]"))
}
