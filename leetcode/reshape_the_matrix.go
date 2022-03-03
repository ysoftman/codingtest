/*
https://leetcode.com/problems/reshape-the-matrix/

566. Reshape the Matrix
In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.
The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input: mat = [[1,2],[3,4]], r = 1, c = 4
Output: [[1,2,3,4]]

Example 2:
Input: mat = [[1,2],[3,4]], r = 2, c = 4
Output: [[1,2],[3,4]]
*/
package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	// if legal, return original matrix
	row := len(mat)
	col := len(mat[0])
	if row*col != r*c {
		return mat
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}
	x, y := 0, 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			result[y][x] = mat[i][j]
			x++
			if x >= c {
				y++
				x = 0
			}
		}
	}
	return result
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4)) // legal
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 4, 1))
}
