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

func distance(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	r := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if distance(points[i][0], points[i][1], points[j][0], points[j][1]) == distance(points[i][0], points[i][1], points[k][0], points[k][1]) {
					r++
				}
				if distance(points[j][0], points[j][1], points[i][0], points[i][1]) == distance(points[j][0], points[j][1], points[k][0], points[k][1]) {
					r++
				}
				if distance(points[k][0], points[k][1], points[i][0], points[i][1]) == distance(points[j][0], points[j][1], points[k][0], points[k][1]) {
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
