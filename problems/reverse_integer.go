// Source : https://leetcode.com/problems/reverse-integer/
// Author : Ke Li
// Date   : 2016-10-26

// Reverse digits of an integer.
//
// Example1: x =  123, return  321
// Example2: x = -123, return -321
//
// click to show spoilers.
//
package main

import "fmt"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	positive := true
	if x < 0 {
		positive = false
		x = -x
	}
	list := []int64{}
	for x > 0 {
		list = append(list, int64(x%10))
		x /= 10
	}
	var r int64 = 0
	for i := 0; i < len(list); i++ {
		r = r*(int64(10)) + list[i]
	}
	if !positive {
		r *= -1
	}
	if r > ((1<<31)-1) || r < -(1<<31) {
		return 0
	}
	return int(r)
}

func main() {
	fmt.Println(reverse(100))
	fmt.Println(reverse(-100))
}
