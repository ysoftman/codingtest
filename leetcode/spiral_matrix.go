/*
https://leetcode.com/problems/spiral-matrix/
54. Spiral Matrix
Medium

Given an m x n matrix, return all elements of the matrix in spiral order.
Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	left, top, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	r := []int{}
	for left <= right && top <= bottom {
		// left -> right
		if left > right {
			break
		}
		for i := left; i <= right; i++ {
			r = append(r, matrix[top][i])
		}
		top++

		// right -> bottom
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			r = append(r, matrix[i][right])
		}
		right--

		// right -> left
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			r = append(r, matrix[bottom][i])
		}
		bottom--

		// bottom -> top
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			r = append(r, matrix[i][left])
		}
		left++
	}
	return r
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5, 5, 6}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}}))
}
