// Source : https://leetcode.com/problems/string-to-integer-atoi/
// Author : Ke Li
// Date   : 2016-10-27

// Implement atoi to convert a string to an integer.
//
// Hint: Carefully consider all possible input cases. If you want a challenge, please do not see
// below and ask yourself what are the possible input cases.
//
// Notes:
// It is intended for this problem to be specified vaguely (ie, no given input specs). You are
// responsible to gather all the input requirements up front.
//
// Update (2015-02-10):
// The signature of the C++ function had been updated. If you still see your function signature
// accepts a const char * argument, please click the reload button  to reset your code definition.
//
// spoilers alert... click to show requirements for atoi.
//
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	var factor int64 = 1
	if str[0] == '-' {
		factor = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	var num int64
	for _, c := range str {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			break
		}
		num = num*10 + int64(v)
		if num*factor < math.MinInt32 {
			return math.MinInt32
		}
		if num*factor > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return int(num * factor)
}

func main() {
	fmt.Println(myAtoi("     100"))
	fmt.Println(myAtoi("  -123"))
}
