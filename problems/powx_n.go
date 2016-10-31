// Source : https://leetcode.com/problems/powx-n/
// Author : Ke Li
// Date   : 2016-10-31

// Implement pow(x, n).
//
package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return myPow(1/x, n*-1)
	}
	f := myPow(x, n/2)
	return f * f * myPow(x, n-n/2*2)
}

func main() {
	fmt.Println(myPow(2, 10))
}
