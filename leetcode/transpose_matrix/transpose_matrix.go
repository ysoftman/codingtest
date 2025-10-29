/*
https://leetcode.com/problems/transpose-matrix/
867. Transpose Matrix
Easy
Given a 2D integer array matrix, return the transpose of matrix.
The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]

Example 2:
Input: matrix = [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 1000
1 <= m * n <= 105
-109 <= matrix[i][j] <= 109
*/
package main

import "fmt"

/*
1,2,3
4,5,6

1,4
2,5
3,6
*/
func transpose(matrix [][]int) [][]int {
	// prepare matrix ex) 2x3 => 3x2
	result := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		result[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

func main() {
	r := transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			fmt.Printf("%v ", r[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-----")
	r = transpose([][]int{{1, 2, 3}, {4, 5, 6}})
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			fmt.Printf("%v ", r[i][j])
		}
		fmt.Println()
	}
}
