/*
https://leetcode.com/problems/spiral-matrix-ii/
59. Spiral Matrix II
Medium

Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

Example 1:
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]

Example 2:
Input: n = 1
Output: [[1]]

Constraints:
1 <= n <= 20

*/

package main

import "fmt"

func generateMatrix(n int) [][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
	}
	cnt := 1
	left, top, right, bottom := 0, 0, n-1, n-1
	for left <= right || top <= bottom {
		for i := left; i <= right; i++ {
			r[top][i] = cnt
			cnt++
		}
		top++
		for i := top; i <= bottom; i++ {
			r[i][right] = cnt
			cnt++
		}
		right--
		for i := right; i >= left; i-- {
			r[bottom][i] = cnt
			cnt++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			r[i][left] = cnt
			cnt++
		}
		left++
	}
	return r
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(5))
	fmt.Println(generateMatrix(6))
}
