// Source : https://leetcode.com/problems/number-of-boomerangs/
// Author : Ke Li
// Date   : 2016-11-14

// Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points
// (i, j, k) such that the distance between i and j equals the distance between i and k (the order
// of the tuple matters).
//
// Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of
// points are all in the range [-10000, 10000] (inclusive).
//
// Example:
//
// Input:
// [[0,0],[1,0],[2,0]]
//
// Output:
// 2
//
// Explanation:
// The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
//
package main

import "fmt"

func distance(p [][]int, i, j int) int {
	return (p[i][0]-p[j][0])*(p[i][0]-p[j][0]) + (p[i][1]-p[j][1])*(p[i][1]-p[j][1])
}

func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	r := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if distance(points, i, j) == distance(points, i, k) {
					r++
				}
				if distance(points, i, j) == distance(points, j, k) {
					r++
				}
				if distance(points, i, k) == distance(points, j, k) {
					r++
				}
			}
		}
	}
	return r * 2
}

func main() {
	points := [][]int{}
	points = append(points, []int{1, 0})
	points = append(points, []int{0, 0})
	points = append(points, []int{2, 0})
	fmt.Println(numberOfBoomerangs(points))
}
