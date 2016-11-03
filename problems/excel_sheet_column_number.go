// Source : https://leetcode.com/problems/excel-sheet-column-number/
// Author : Ke Li
// Date   : 2016-11-3

// Related to question Excel Sheet Column Title
// Given a column title as appear in an Excel sheet, return its corresponding column number.
//
// For example:
// A - 1
// B - 2
// C - 3
// ...
// Z - 26
// AA - 27
// AB - 28
//
// Credits:Special thanks to @ts for adding this problem and creating all test cases.
//
package main

func titleToNumber(s string) int {
	r := 0
	for _, c := range s {
		r = r*26 + (c - 'A' + 1)
	}
	return r
}
